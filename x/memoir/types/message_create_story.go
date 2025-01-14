package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateStory{}

func NewMsgCreateStory(creator string, title string, content string) *MsgCreateStory {
	return &MsgCreateStory{
		Creator: creator,
		Title:   title,
		Content: content,
	}
}

func (msg *MsgCreateStory) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
