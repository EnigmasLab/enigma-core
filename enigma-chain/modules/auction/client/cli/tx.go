package cli

import (
	"github.com/spf13/cobra"

	"github.com/EnigmasLab/enigma-core/cli"

	"github.com/EnigmasLab/enigma-core/enigma-chain/modules/auction/types"
)

const (
	FlagRoundNumber = "round"
	FlagBidAmount   = "bid"
)

// NewTxCmd returns a root CLI command handler for certain modules/auction transaction commands.
func NewTxCmd() *cobra.Command {
	cmd := cli.ModuleRootCommand(types.ModuleName, false)

	cmd.AddCommand(
		NewBidCmd(),
	)
	return cmd
}

func NewBidCmd() *cobra.Command {
	cmd := cli.TxCmd("bid",
		"bid on current exchange basket",
		&types.MsgBid{},
		cli.FlagsMapping{"BidAmount": cli.Flag{Flag: FlagBidAmount}, "Round": cli.Flag{Flag: FlagRoundNumber}},
		cli.ArgsMapping{},
	)
	cmd.Example = `enigmad tx auction bid --bid="100000000000000000000fury" --round=4 --from=genesis --keyring-backend=file --yes`
	cmd.Flags().String(FlagBidAmount, "100000000000000000000fury", "Auction bid amount")
	cmd.Flags().Uint64(FlagRoundNumber, 4, "Auction round number")
	return cmd
}