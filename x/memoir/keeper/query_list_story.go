package keeper

import (
	"context"
	"memoir/x/memoir/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListStory(goCtx context.Context, req *types.QueryListStoryRequest) (*types.QueryListStoryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoryKey))

	var stories []types.Story
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var story types.Story
		k.cdc.MustUnmarshal(iterator.Value(), &story)
		stories = append(stories, story)
	}

	storyPointers := make([]*types.Story, len(stories))
	for i := range stories {
		storyPointers[i] = &stories[i]
	}
	return &types.QueryListStoryResponse{Stories: storyPointers}, nil
}
