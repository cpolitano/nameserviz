package nameserviz

import (
	"github.com/cosmos/cosmos-sdk/codec" // tools to work with Cosmos encoding
	"github.com/cosmos/cosmos-sdk/x/bank" // controls accounts and coin transfers

	sdk "github.com/cosmos/cosmos-sdk/types" // commonly used types in SDK
)

// connect to data store, get/set methods
type Keeper struct {
  coinKeeper bank.coinKeeper

  namesStoreKey sdk.StoreKey // key to access namesStore
  ownersStoreKey sdk.StoreKey // key to owners
  pricesStoreKey sdk.StoreKey // key to prices

  cdc *codec.Codec // wire codec for binary encoding/decoding
}
