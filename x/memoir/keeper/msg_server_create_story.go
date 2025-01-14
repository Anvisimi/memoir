package keeper

import (
	"context"

	"memoir/x/memoir/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateStory(goCtx context.Context, msg *types.MsgCreateStory) (*types.MsgCreateStoryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var story = types.Story{
		Id:      0,
		Title:   msg.Title,
		Content: msg.Content,
		Author:  msg.Creator,
	}

	id := k.AppendStory(
		ctx,
		story,
	)

	return &types.MsgCreateStoryResponse{
		Id: id,
	}, nil
}
