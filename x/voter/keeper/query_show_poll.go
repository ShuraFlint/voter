package keeper

import (
	"context"
	"strconv"
	"strings"

	"voter/x/voter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowPoll(goCtx context.Context, req *types.QueryShowPollRequest) (*types.QueryShowPollResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	poll, found := k.GetPollById(goCtx, req.PollId)

	if !found {
		return nil, status.Error(codes.NotFound, "poll not found ")
	}

	return &types.QueryShowPollResponse{
		Creator: poll.Creator,
		Id:      strconv.FormatUint(poll.Id, 10),
		Title:   poll.Title,
		Options: strings.Join(poll.Options, ","),
	}, nil
}
