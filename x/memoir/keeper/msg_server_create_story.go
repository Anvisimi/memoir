package keeper

import (
	"context"

	"memoir/x/memoir/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateStory(goCtx context.Context, msg *types.MsgCreateStory) (*types.MsgCreateStoryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// New ID generation logic that breaks consensus
	lastID := k.GetLastStoryID(ctx)
	newID := lastID + 2 // Breaking change: IDs now increment by 2

	var story = types.Story{
		Id:       newID,
		Title:    msg.Title,
		Content:  msg.Content,
		Author:   msg.Creator,
		Category: "default", // Required field
		Rating:   100,       // Required field
	}

	k.SetStory(ctx, story)
	k.SetLastStoryID(ctx, newID)

	return &types.MsgCreateStoryResponse{
		Id: newID,
	}, nil
}
