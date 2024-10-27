package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-security/v6/x/intertx/types"
	"github.com/spf13/cobra"
	"strconv"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Interchain account control(ICA) transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(RegisterButlerCmd())
	cmd.AddCommand(TransferAssetCmd())
	return cmd
}

func RegisterButlerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-butler [from-address-or-key] [target-chain-id] [chain-tag] [connection-id]",
		Short: "Register new butler",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[0]); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			chainTag := args[2]
			connectionID := args[3]

			msg := &types.MsgCreateButler{
				TargetChainId: chainID,
				FromAddress:   clientCtx.FromAddress.String(),
				ChainTag:      chainTag,
				ConnectionId:  connectionID,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// TODO : for testing
func TransferAssetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-asset [from-address-or-key] [id] [target-chain-id] [memo] [token] [pair]",
		Short: "Transfer asset to IBC chain",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[0]); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			chainID, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			memo := args[3]
			token, err := sdk.ParseCoinNormalized(args[4])
			if err != nil {
				return err
			}

			pair, err := sdk.ParseCoinNormalized(args[5])
			if err != nil {
				return err
			}

			msg := &types.MsgTransfer{
				TargetChainId: chainID,
				Sender:        clientCtx.FromAddress.String(),
				Id:            id,
				Token:         &token,
				Pair:          &pair,
				Memo:          memo,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
