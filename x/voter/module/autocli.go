package voter

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "voter/api/voter/voter"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "ShowPoll",
					Use:            "show-poll [poll-id]",
					Short:          "Query show-poll",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "pollId"}},
				},

				{
					RpcMethod:      "ShowPollVotes",
					Use:            "show-poll-votes [poll-id]",
					Short:          "Query show-poll-votes",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "pollId"}},
				},

				{
					RpcMethod:      "ShowAllPolls",
					Use:            "show-all-polls",
					Short:          "Query show-all-polls",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreatePoll",
					Use:            "create-poll [title] [options]",
					Short:          "Send a create-poll tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "options"}},
				},
				{
					RpcMethod:      "CastVote",
					Use:            "cast-vote [poll-id] [option]",
					Short:          "Send a cast-vote tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "pollId"}, {ProtoField: "option"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
