package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"memoir/x/memoir/types"
)

// AppendStory appends a story in the store with a new id and update the count
func (k Keeper) AppendStory(ctx sdk.Context, story types.Story) uint64 {
	count := k.GetStoryCount(ctx)
	story.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoryKey))
	appendedValue := k.cdc.MustMarshal(&story)
	store.Set(GetStoryIDBytes(story.Id), appendedValue)
	k.SetStoryCount(ctx, count+1)
	return count
}

// GetStoryCount get the total number of story
func (k Keeper) GetStoryCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoryCountKey))
	byteKey := []byte("count")
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// GetStoryIDBytes returns the byte representation of the ID
func GetStoryIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// SetStoryCount set the total number of story
func (k Keeper) SetStoryCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoryCountKey))
	byteKey := []byte("count")
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// GetStory returns a story from its id
func (k Keeper) GetStory(ctx sdk.Context, id uint64) (val types.Story, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoryKey))
	b := store.Get(GetStoryIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStory removes a story from the store
func (k Keeper) RemoveStory(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoryKey))
	store.Delete(GetStoryIDBytes(id))
}

// SetStory set a specific story in the store
func (k Keeper) SetStory(ctx sdk.Context, story types.Story) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoryKey))
	b := k.cdc.MustMarshal(&story)
	store.Set(GetStoryIDBytes(story.Id), b)
}
