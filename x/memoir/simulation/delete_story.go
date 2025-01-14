package simulation

import (
	"math/rand"

	"memoir/x/memoir/keeper"
	"memoir/x/memoir/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgDeleteStory(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgDeleteStory{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the DeleteStory simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "DeleteStory simulation not implemented"), nil, nil
	}
}
