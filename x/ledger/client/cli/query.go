package cli

import (
	"encoding/hex"
	"fmt"

	"strconv"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/stafiprotocol/stafihub/x/ledger/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group ledger queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdListExchangeRate())
	cmd.AddCommand(CmdShowExchangeRate())
	cmd.AddCommand(CmdShowEraExchangeRate())
	cmd.AddCommand(CmdEraExchangeRateByDenom())

	cmd.AddCommand(CmdPoolsByDenom())
	cmd.AddCommand(CmdBondedPoolsByDenom())
	cmd.AddCommand(CmdGetPoolDetail())
	cmd.AddCommand(CmdGetChainEra())
	cmd.AddCommand(CmdGetCurrentEraSnapshot())
	cmd.AddCommand(CmdGetReceiver())
	cmd.AddCommand(CmdGetCommission())
	cmd.AddCommand(CmdGetChainBondingDuration())
	cmd.AddCommand(CmdGetUnbondFee())
	cmd.AddCommand(CmdGetUnbondCommission())
	cmd.AddCommand(CmdGetLeastBond())
	cmd.AddCommand(CmdGetEraUnbondLimit())

cmd.AddCommand(CmdGetBondPipeLine())

cmd.AddCommand(CmdGetEraSnapshot())

cmd.AddCommand(CmdGetSnapshot())

// this line is used by starport scaffolding # 1

	return cmd
}

var _ = strconv.Itoa(0)

func CmdPoolsByDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pools-by-denom [denom]",
		Short: "Query pools_by_denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryPoolsByDenomRequest{
				Denom: reqDenom,
			}

			res, err := queryClient.PoolsByDenom(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdBondedPoolsByDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bonded-pools-by-denom [denom]",
		Short: "Query bonded_pools_by_denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryBondedPoolsByDenomRequest{
				Denom: reqDenom,
			}

			res, err := queryClient.BondedPoolsByDenom(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetPoolDetail() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-pool-detail [denom] [pool]",
		Short: "Query get_pool_detail",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]
			reqPool := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetPoolDetailRequest{
				Denom: reqDenom,
				Pool:  reqPool,
			}

			res, err := queryClient.GetPoolDetail(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetChainEra() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-chain-era [denom]",
		Short: "Query getChainEra",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetChainEraRequest{
				Denom: reqDenom,
			}

			res, err := queryClient.GetChainEra(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetCurrentEraSnapshot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-current-era-snap-shot [denom]",
		Short: "Query getCurrentEraSnapshot",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetCurrentEraSnapshotRequest{
				Denom: reqDenom,
			}

			res, err := queryClient.GetCurrentEraSnapshot(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetReceiver() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-receiver",
		Short: "Query getReceiver",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetReceiverRequest{}

			res, err := queryClient.GetReceiver(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetCommission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-commission",
		Short: "Query getCommission",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetCommissionRequest{}

			res, err := queryClient.GetCommission(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetChainBondingDuration() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-chain-bonding-duration [denom]",
		Short: "Query getChainBondingDuration",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetChainBondingDurationRequest{
				Denom: reqDenom,
			}

			res, err := queryClient.GetChainBondingDuration(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetUnbondFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-unbond-fee",
		Short: "Query getUnbondFee",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetUnbondFeeRequest{
			}

			res, err := queryClient.GetUnbondFee(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetUnbondCommission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-unbond-commission",
		Short: "Query getUnbondCommission",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryGetUnbondCommissionRequest{
			}

			res, err := queryClient.GetUnbondCommission(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetLeastBond() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-least-bond [denom]",
		Short: "Query getLeastBond",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryGetLeastBondRequest{
				Denom: reqDenom,
			}

			res, err := queryClient.GetLeastBond(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetEraUnbondLimit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-era-unbond-limit [denom]",
		Short: "Query getEraUnbondLimit",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetEraUnbondLimitRequest{

				Denom: reqDenom,
			}



			res, err := queryClient.GetEraUnbondLimit(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetBondPipeLine() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-bond-pipe-line [denom] [pool]",
		Short: "Query GetBondPipeLine",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]
			reqPool := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryGetBondPipeLineRequest{
				Denom: reqDenom,
				Pool: reqPool,
			}

			res, err := queryClient.GetBondPipeLine(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetEraSnapshot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-era-snap-shot [denom] [era]",
		Short: "Query GetEraSnapshot",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]
			reqEra, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryGetEraSnapshotRequest{
				Denom: reqDenom,
				Era: uint32(reqEra),
			}

			res, err := queryClient.GetEraSnapshot(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetSnapshot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-snap-shot [shot-id]",
		Short: "Query GetSnapShot",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqShotId, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryGetSnapshotRequest{
				ShotId: reqShotId,
			}

			res, err := queryClient.GetSnapshot(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
