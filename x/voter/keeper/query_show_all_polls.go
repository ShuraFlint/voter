package keeper

import (
	"context"

	"voter/x/voter/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowAllPolls(goCtx context.Context, req *types.QueryShowAllPollsRequest) (*types.QueryShowAllPollsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.PollKey))

	var polls []*types.Poll
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var poll types.Poll
		if err := k.cdc.Unmarshal(value, &poll); err != nil {
			return err
		}
		polls = append(polls, &poll)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryShowAllPollsResponse{
		Polls:      polls,
		Pagination: pageRes,
	}, nil
}