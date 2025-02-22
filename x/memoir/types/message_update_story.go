package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateStory{}

func NewMsgUpdateStory(creator string, title string, content string, id uint64) *MsgUpdateStory {
	return &MsgUpdateStory{
		Creator: creator,
		Title:   title,
		Content: content,
		Id:      id,
	}
}

func (msg *MsgUpdateStory) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
