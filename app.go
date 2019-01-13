package app

import (
	"github.com/cosmos/cosmos-sdk/x/auth"       // auth module
	"github.com/tendermint/tendermint/libs/log" // Tendermint logger

	bam "github.com/cosmos/cosmos-sdk/baseapp"     // Cosmos baseapp, implements ABCI
	dbm "github.com/tendermint/tendermint/libs/db" // working with Tendermint db
)

const (
	appName = "nameserviz"
)

type nameservizApp struct {
	*bam.BaseApp
}

func NewnameservizApp(logger log.logger, db dbm.DB) *nameservizApp {

	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	var app = &nameservizApp{
		BaseApp: bApp,
		cdc:     cdc,
	}

	return app
}
