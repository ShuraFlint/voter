package voter

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"voter/testutil/sample"
	votersimulation "voter/x/voter/simulation"
	"voter/x/voter/types"
)

// avoid unused import issue
var (
	_ = votersimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreatePoll = "op_weight_msg_create_poll"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePoll int = 100

	opWeightMsgCastVote = "op_weight_msg_cast_vote"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCastVote int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	voterGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&voterGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePoll int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePoll, &weightMsgCreatePoll, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePoll = defaultWeightMsgCreatePoll
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePoll,
		votersimulation.SimulateMsgCreatePoll(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCastVote int
	simState.AppParams.GetOrGenerate(opWeightMsgCastVote, &weightMsgCastVote, nil,
		func(_ *rand.Rand) {
			weightMsgCastVote = defaultWeightMsgCastVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCastVote,
		votersimulation.SimulateMsgCastVote(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreatePoll,
			defaultWeightMsgCreatePoll,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				votersimulation.SimulateMsgCreatePoll(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCastVote,
			defaultWeightMsgCastVote,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				votersimulation.SimulateMsgCastVote(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
