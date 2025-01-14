package memoir

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"memoir/testutil/sample"
	memoirsimulation "memoir/x/memoir/simulation"
	"memoir/x/memoir/types"
)

// avoid unused import issue
var (
	_ = memoirsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateStory = "op_weight_msg_create_story"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateStory int = 100

	opWeightMsgUpdateStory = "op_weight_msg_update_story"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateStory int = 100

	opWeightMsgDeleteStory = "op_weight_msg_delete_story"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteStory int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	memoirGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&memoirGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateStory int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateStory, &weightMsgCreateStory, nil,
		func(_ *rand.Rand) {
			weightMsgCreateStory = defaultWeightMsgCreateStory
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateStory,
		memoirsimulation.SimulateMsgCreateStory(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateStory int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateStory, &weightMsgUpdateStory, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateStory = defaultWeightMsgUpdateStory
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateStory,
		memoirsimulation.SimulateMsgUpdateStory(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteStory int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteStory, &weightMsgDeleteStory, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteStory = defaultWeightMsgDeleteStory
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteStory,
		memoirsimulation.SimulateMsgDeleteStory(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateStory,
			defaultWeightMsgCreateStory,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				memoirsimulation.SimulateMsgCreateStory(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateStory,
			defaultWeightMsgUpdateStory,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				memoirsimulation.SimulateMsgUpdateStory(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteStory,
			defaultWeightMsgDeleteStory,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				memoirsimulation.SimulateMsgDeleteStory(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
