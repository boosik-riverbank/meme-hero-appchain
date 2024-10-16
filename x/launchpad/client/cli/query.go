package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/interchain-security/v6/x/launchpad/types"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Launchpad query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdPools(),
		GetCmdPool(),
		GetCmdPoolAmount(),
	)

	return cmd
}

func GetCmdPools() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pools",
		Short: "Query all pools",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Pools(cmd.Context(), &types.QueryPools{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pool <poolID>",
		Short: "Query pool info",
		Long: strings.TrimSpace(
			fmt.Sprintf("Query pool info.\nExample: memed query launchpad pool 1"),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			poolID, err := strconv.Atoi(args[0])
			res, err := queryClient.Pool(cmd.Context(), &types.QueryPool{
				Id: uint64(poolID),
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdPoolAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pool-amount <poolID>",
		Short: "Query pool amount",
		Long: strings.TrimSpace(
			fmt.Sprintf("Query pool tokens amount.\nExample: %s query launchpad pool-amount 1", version.AppName),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			poolID, err := strconv.Atoi(args[0])
			res, err := queryClient.PoolAmount(cmd.Context(), &types.QueryPoolAmount{
				Id: uint64(poolID),
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
