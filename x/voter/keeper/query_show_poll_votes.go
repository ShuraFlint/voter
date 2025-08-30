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

func (k Keeper) ShowPollVotes(goCtx context.Context, req *types.QueryShowPollVotesRequest) (*types.QueryShowPollVotesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.VoteKey))

	//声明一个vote指针数组，存放符合条件的投票
	var votes []*types.Vote

	//根据store的规则（vote/value/）遍历其中的所有内容，
	//根据 req.Pagination（也就是 PageRequest）来决定返回多少条、从哪条开始、是否统计总数。
	//没便利一条数据，就会调用回调函数func(key []byte, value []byte) error
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var vote types.Vote
		if err := k.cdc.Unmarshal(value, &vote); err != nil {
			return err
		}
		if vote.PollID == req.PollId {
			votes = append(votes, &vote)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryShowPollVotesResponse{
		Votes:      votes,
		Pagination: pageRes,
	}, nil
}
