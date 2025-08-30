package keeper

import (
	"context"
	"fmt"

	"voter/x/voter/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// goCtx 是 标准 Go 的 Context，来自 gRPC 调用。
func (k msgServer) CreatePoll(goCtx context.Context, msg *types.MsgCreatePoll) (*types.MsgCreatePollResponse, error) {
	//把goCtx转换为 Cosmos SDK 的 Context (sdk.Context)。
	// 转换后你才能使用 KVStore、ctx.EventManager()、ctx.Logger() 等区块链功能。
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	//获得模块的账户地址 sdk.AccAddress类型
	moduleAcct := k.accountKeeper.GetModuleAddress(types.ModuleName)
	if moduleAcct == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnknownAddress, "module account does not exist")
	}

	//定义费用，即将string类型token转换成coin类型token，必须是：数字+币单位
	feeCoins, err := sdk.ParseCoinsNormalized("27token")
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidCoins, "invalid fee amount")
	}

	//获得交易人地址 将string类型地址转换为AccAddress类型地址
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	//检查交易人地址是否有足够的token来支付费用
	spendableCoins := k.bankKeeper.SpendableCoins(ctx, creator)
	if !spendableCoins.IsAllGTE(feeCoins) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInsufficientFunds, "insufficient funds")
	}

	//支付费用
	//ctx:当前区块链上下文（sdk.Context），包含区块高度、时间戳、KVStore 等信息。所有链上状态变更都必须通过 ctx 执行。
	//creator:交易人地址，即将支付费用的地址。
	//moduleAcct:模块账户地址，即接收支付费用的地址。
	//feeCoins:支付的token数量。
	if err := k.bankKeeper.SendCoins(ctx, creator, moduleAcct, feeCoins); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to pay fee")
	}

	k.DebugPoll()
	fmt.Println("支付200token成功。moduleaccount = ", moduleAcct.String())
	fmt.Println("feeCoins = ", feeCoins.String())
	fmt.Println("spendableCoins = ", spendableCoins.String())
	k.DebugPoll()

	poll := types.Poll{
		Creator: msg.Creator,
		Title:   msg.Title,
		Options: msg.Options,
	}
	id := k.AppendPoll(ctx, poll)

	return &types.MsgCreatePollResponse{Id: int32(id), Title: msg.Title}, nil
}
