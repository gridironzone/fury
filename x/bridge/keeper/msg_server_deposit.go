package keeper

import (
	"context"
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/bridge/types"
)

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := hex.DecodeString(msg.Receiver)
	if err != nil {
		return nil, types.ErrReceiverFormatNotRight
	}
	if msg.Amount.LTE(sdk.ZeroInt()) {
		return nil, types.ErrDepositAmountZero
	}

	chainId := uint8(msg.DestChainId)
	if !k.Keeper.HasChainId(ctx, chainId) {
		return nil, types.ErrChainIdNotSupport
	}

	resourceIdSlice, err := hex.DecodeString(msg.ResourceId)
	if err != nil {
		return nil, types.ErrResourceIdFormatNotRight
	}
	var resourceId [32]byte
	copy(resourceId[:], resourceIdSlice)

	denom, found := k.Keeper.GetDenomByResourceId(ctx, resourceId)
	if !found {
		return nil, types.ErrResourceIdNotFound
	}
	userAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// relay fee
	relayFeeReceiver, found := k.Keeper.GetRelayFeeReceiver(ctx)
	if !found {
		return nil, types.ErrRelayFeeReceiverNotSet
	}
	relayFee := k.Keeper.GetRelayFee(ctx, chainId)
	if relayFee.Amount.GT(sdk.ZeroInt()) {
		err := k.bankKeeper.SendCoins(ctx, userAddress, relayFeeReceiver, sdk.NewCoins(relayFee))
		if err != nil {
			return nil, err
		}
	}

	// lock or burn token
	balance := k.bankKeeper.GetBalance(ctx, userAddress, denom)
	if balance.Amount.LT(msg.Amount) {
		return nil, types.ErrBalanceNotEnough
	}

	shouldBurnedOrLockedCoins := sdk.NewCoins(sdk.NewCoin(denom, msg.Amount))
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, userAddress, types.ModuleName, shouldBurnedOrLockedCoins)
	if err != nil {
		return nil, err
	}

	resourceIdType := k.Keeper.GetResourceIdType(ctx, resourceId)
	if resourceIdType == types.ResourceIdTypeForeign {
		err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, shouldBurnedOrLockedCoins)
		if err != nil {
			return nil, err
		}
	}

	//update deposit count
	count := k.Keeper.GetDepositCounts(ctx, chainId)
	k.Keeper.SetDepositCounts(ctx, chainId, count+1)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDeposit,
			sdk.NewAttribute(types.AttributeKeyDestChainId, fmt.Sprintf("%d", chainId)),
			sdk.NewAttribute(types.AttributeKeyResourceId, msg.ResourceId),
			sdk.NewAttribute(types.AttributeKeyDepositNonce, fmt.Sprintf("%d", count)),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
		),
	)

	return &types.MsgDepositResponse{}, nil
}
