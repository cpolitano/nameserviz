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
  ownersStoreKey sdk.StoreKey // key to owner of a given name
  pricesStoreKey sdk.StoreKey // key to price paid for name

  cdc *codec.Codec // codec for binary encoding/decoding
}

// sets the value string that a name resolves to
func (k Keeper) SetName(ctx sdk.Context, name string, value string) {
	store := ctx.KVStore(k.namesStoreKey)
	store.Set([]byte(name), []byte(value)) // store only takes []byte
}

// returns the string that the name resolves to
func (k Keeper) ResolveName(ctx sdk.Context, name string) string {
	store := ctx.KVStore(k.namesStoreKey)
	bz := store.Get([]byte(name)) // cast string to []byte
	return string(bz) // cast []byte to string
}

// returns bool whether or not the name already has an owner
func (k Keeper) HasOwner(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.ownersStoreKey)
	bz := store.Get([]byte(name))
	return bz != nil
}

// get the current owner of a name
// sdk.AccAddress is type alias for []byte
func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
	store := ctx.KVStore(k.ownersStoreKey)
	bz := store.Get([]byte(name))
	return bz
}

// sets the current owner of a name
func (k Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAddress) {
	store := ctx.KVStore(k.ownersStoreKey)
	store.Set([]byte(name), owner)
}

// gets the current price of a name.  If price doesn't exist yet, set to 1steak.
// sdk.Coins lacks byte encoding, needs marshalling/unmarshalling with Amino to insert into store
func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
	if !k.HasOwner(ctx, name) {
		return sdk.Coins{sdk.NewInt64Coin("mycoin", 1)}
	}
	store := ctx.KVStore(k.pricesStoreKey)
	bz := store.Get([]byte(name))
	var price sdk.Coins
	k.cdc.MustUnmarshalBinaryBare(bz, &price)
	return price
}

// sets the current price of a name
func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
	store := ctx.KVStore(k.pricesStoreKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(price))
}

// NewKeeper creates new instances of the nameserviz Keeper
func NewKeeper(coinKeeper bank.Keeper, namesStoreKey sdk.StoreKey, ownersStoreKey sdk.StoreKey, priceStoreKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper:     coinKeeper,
		namesStoreKey:  namesStoreKey,
		ownersStoreKey: ownersStoreKey,
		pricesStoreKey: priceStoreKey,
		cdc:            cdc,
	}
}
