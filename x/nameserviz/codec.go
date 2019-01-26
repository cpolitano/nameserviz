package nameserviz

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on wire codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetName{}, "nameserviz/SetName", nil)
	cdc.RegisterConcrete(MsgBuyName{}, "nameserviz/BuyName", nil)
}
