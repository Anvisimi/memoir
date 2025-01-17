package keeper

import (
	"context"

	"memoir/x/memoir/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowStory(goCtx context.Context, req *types.QueryShowStoryRequest) (*types.QueryShowStoryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	story, found := k.GetStory(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "story not found")
	}

	return &types.QueryShowStoryResponse{Story: &story}, nil
}
