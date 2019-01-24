package app

import (
	"encoding/json"

	"github.com/tendermint/tendermint/libs/log" // Tendermint logger

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth" // auth module
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/stake"
	"github.com/cpolitano/nameserviz/x/nameserviz"

	bam "github.com/cosmos/cosmos-sdk/baseapp" // Cosmos baseapp, implements ABCI
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db" // working with Tendermint db
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	appName = "nameserviz"
)

type nameServizApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	keyMain          *sdk.KVStoreKey
	keyAccount       *sdk.KVStoreKey
	keyNSnames       *sdk.KVStoreKey
	keyNSowners      *sdk.KVStoreKey
	keyNSprices      *sdk.KVStoreKey
	keyFeeCollection *sdk.KVStoreKey

	accountKeeper       auth.AccountKeeper
	bankKeeper          bank.Keeper
	feeCollectionKeeper auth.FeeCollectionKeeper
	nsKeeper            nameserviz.Keeper
}

func NewNameServizApp(logger log.logger, db dbm.DB) *nameServizApp {

	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	var app = &nameServizApp{
		BaseApp: bApp,
		cdc:     cdc,

		keyMain:          sdk.NewKVStoreKey("main"),
		keyAccount:       sdk.NewKVStoreKey("acc"),
		keyNSnames:       sdk.NewKVStoreKey("ns_names"),
		keyNSowners:      sdk.NewKVStoreKey("ns_owners"),
		keyNSprices:      sdk.NewKVStoreKey("ns_prices"),
		keyFeeCollection: sdk.NewKVStoreKey("fee_collection"),
	}

	return app
}
