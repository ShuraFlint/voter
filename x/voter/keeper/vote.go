package keeper

import (
	"context"
	"encoding/binary"
	"fmt"
	"voter/x/voter/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
)

func (k Keeper) AppendVote(goCtx context.Context, vote types.Vote) error {
	count := k.GetVoteCount(goCtx)
	vote.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, []byte(types.VoteKey))
	store.Set(k.GetIdBytes(vote.Id), k.cdc.MustMarshal(&vote))
	k.SetVoteCount(goCtx, count+1)
	k.GetVoteCount(goCtx)
	return nil
}

func (k Keeper) GetVoteCount(goCtx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, []byte{})
	bytekey := []byte(types.VoteCountKey)
	bz := store.Get(bytekey)
	if bz == nil {
		k.DebugPoll()
		fmt.Println("vote count is 0")
		k.DebugPoll()
		return 0
	}
	num := binary.BigEndian.Uint64(bz)

	k.DebugPoll()
	fmt.Println("vote count is ", num)
	k.DebugPoll()
	return num
}

func (k Keeper) SetVoteCount(goCtx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, []byte{})
	bytekey := []byte(types.VoteCountKey)
	store.Set(bytekey, k.GetIdBytes(count))
}
