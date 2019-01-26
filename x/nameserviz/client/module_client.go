package client

// standard way for modules to export client functionality

import (
	"github.com/cosmos/cosmos-sdk/client"
	nameservizcmd "github.com/cpolitano/nameserviz/x/nameserviz/client/cli" // import my cli code
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	// Group gov queries under a subcommand
	govQueryCmd := &cobra.Command{
		Use:   "nameserviz",
		Short: "Querying commands for the nameserviz module",
	}

	govQueryCmd.AddCommand(client.GetCommands(
		nameservizcmd.GetCmdResolveName(mc.storeKey, mc.cdc),
		nameservizcmd.GetCmdWhois(mc.storeKey, mc.cdc),
	)...)

	return govQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	govTxCmd := &cobra.Command{
		Use:   "nameserviz",
		Short: "Nameserviz transactions subcommands",
	}

	govTxCmd.AddCommand(client.PostCommands(
		nameservizcmd.GetCmdBuyName(mc.cdc),
		nameservizcmd.GetCmdSetName(mc.cdc),
	)...)

	return govTxCmd
}
