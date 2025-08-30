package keeper

import (
	"context"
	"fmt"
	"strconv"

	"voter/x/voter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CastVote(goCtx context.Context, msg *types.MsgCastVote) (*types.MsgCastVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	pollId, err := strconv.ParseUint(msg.PollId, 10, 64)
	if err != nil {
		return nil, err
	}

	vote := types.Vote{
		Creator: msg.Creator,
		PollID:  pollId,
		Option:  msg.Option,
	}

	k.AppendVote(ctx, vote)
	k.DebugPoll()
	fmt.Println("ceator = ", msg.Creator, " pollId = ", msg.PollId, " option = ", msg.Option)
	k.DebugPoll()

	return &types.MsgCastVoteResponse{}, nil
}
