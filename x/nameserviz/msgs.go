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

// implement msg interface
// Type should return the name of the module
func (msg MsgSetName) Route() string { return "nameserviz" }

// Name should return the action
func (msg MsgSetName) Type() string { return "set_name"}

// ValdateBasic Implements Msg.
// stateless check that msg data is valid
func (msg MsgSetName) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.NameID) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
// pure function to define msg encoding
func (msg MsgSetName) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
// whose signature is required for Tx to be valid
// owner must sign transaction to reset the name
func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
