package keeper

import (
	"context"
	"fmt"

	"memoir/x/memoir/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateStory(goCtx context.Context, msg *types.MsgUpdateStory) (*types.MsgUpdateStoryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetStory(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	if msg.Creator != val.Author { // Compare Creator from msg with Author from stored story
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var story = types.Story{
		Id:      msg.Id,
		Title:   msg.Title,
		Content: msg.Content,
		Author:  msg.Creator, // Note: msg has Creator, but we store it as Author
	}

	k.SetStory(ctx, story)

	return &types.MsgUpdateStoryResponse{}, nil
}
