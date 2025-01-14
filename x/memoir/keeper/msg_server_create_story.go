package keeper

import (
	"context"

	"memoir/x/memoir/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateStory(goCtx context.Context, msg *types.MsgCreateStory) (*types.MsgCreateStoryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var story = types.Story{
		Title:   msg.Title,
		Content: msg.Content,
		Author:  msg.Creator, // Note: msg has Creator, but we store it as Author
	}

	id := k.AppendStory(
		ctx,
		story,
	)

	return &types.MsgCreateStoryResponse{
		Id: id,
	}, nil
}
