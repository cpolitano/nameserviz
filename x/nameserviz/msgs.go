package nameserviz

import (
  "encoding/json"

sdk "github.com/cosmos/cosmos-sdk/types"
)

// defines req for SetName msg
type MsgSetName struct {
  NameID string
  Value string
  Owner sdk.AccAddress
}

// NewMsgSetName is a constructor function for MsgSetName
// returns struct type MsgSetName
func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
		NameID: name,
		Value:  value,
		Owner:  owner,
	}
}
