package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-security/v6/x/launchpad/types"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Launchpad transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCreatePoolCmd(),
		GetBuyTokenCmd(),
		GetSellTokenCmd(),
	)
	return cmd
}

func GetCreatePoolCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-pool <from_address> <token_name> <token_symbol> <token_denom> <pair_denom> <initial_quantity> <description | null> <image_url | null> <website | null> <telegram | null> <twitter | null>",
		Short: "Create new pool",
		Long: strings.TrimSpace(
			fmt.Sprintf("Create new pool.\nExample: memed tx create-pool meme1234..abcd udoge uatom 100000000udoge"),
		),
		Args: cobra.MinimumNArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// coins, err := sdk.ParseCoinsNormalized(args[len(args)-1])
			coins, err := sdk.ParseCoinsNormalized(args[5])
			if err != nil {
				return err
			}

			tokenName := args[1]
			tokenSymbol := args[2]
			tokenDenom := args[3]
			pairDenom := args[4]

			image := args[6]
			description := args[7]
			website := args[8]
			telegram := args[9]
			twitter := args[10]
			if err != nil {
				return err
			}

			msg := types.NewCreatePoolMsg(clientCtx.FromAddress, tokenName, tokenSymbol, tokenDenom, pairDenom, description, image, website, telegram, twitter, &coins[0])
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func GetBuyTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-token <from_address> <pool_id> <amount>",
		Short: "Buy a token",
		Long: strings.TrimSpace(
			fmt.Sprintf("Buy a token.\nExample: memed tx buy-token 1 100000000uatom"),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoinsNormalized(args[len(args)-1])
			if err != nil {
				return err
			}

			poolID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewBuyTokenMsg(clientCtx.FromAddress, poolID, &coins[0])
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func GetSellTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell-token <from_address> <pool_id> <amount>",
		Short: "Sell a token",
		Long: strings.TrimSpace(
			fmt.Sprintf("Sell a token.\nExample: memed tx sell-token 1 100000000udoge"),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoinsNormalized(args[len(args)-1])
			if err != nil {
				return err
			}

			poolID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewSellTokenMsg(clientCtx.FromAddress, poolID, &coins[0])
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
