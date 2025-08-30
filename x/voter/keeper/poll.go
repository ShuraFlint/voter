package keeper

import (
	"encoding/binary"
	"fmt"
	"voter/x/voter/types"

	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendPoll(goCtx context.Context, poll types.Poll) uint64 {

	count := k.GetPollCount(goCtx)
	poll.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, []byte(types.PollKey))
	appendValue := k.cdc.MustMarshal(&poll)
	store.Set(k.GetIdBytes(poll.Id), appendValue)
	k.SetPollCount(goCtx, count+1)
	k.GetPollCount(goCtx)
	return 1
}

func (k Keeper) GetPollCount(goCtx context.Context) uint64 {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})

	byteKey := []byte(types.PollCountKey)
	bz := store.Get(byteKey)

	if bz == nil {
		k.DebugPoll()
		fmt.Println("Poll count is 0")
		k.DebugPoll()
		return 0
	}
	//把这 8 个字节解释成一个 uint64 数字
	//bz 必须是长度为 8 的字节切片
	num := binary.BigEndian.Uint64(bz)

	k.DebugPoll()
	fmt.Println("Poll count is ", num)
	k.DebugPoll()
	return num
}

func (k Keeper) SetPollCount(goCtx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, []byte{})
	store.Set([]byte(types.PollCountKey), k.GetIdBytes(count))
}

func (k Keeper) GetPollById(goCtx context.Context, id uint64) (val types.Poll, found bool) {
	poll := types.Poll{}
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, []byte(types.PollKey))
	bz := store.Get(k.GetIdBytes(id))
	if bz == nil {
		return poll, false
	}
	k.cdc.MustUnmarshal(bz, &poll)
	return poll, true
}
