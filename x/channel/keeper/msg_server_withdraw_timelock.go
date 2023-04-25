package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dungtt-astra/channel/x/channel/types"
)

func (k msgServer) WithdrawTimelock(goCtx context.Context, msg *types.MsgWithdrawTimelock) (*types.MsgWithdrawTimelockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.Keeper.GetCommitment(ctx, msg.Index)
	if !found {
		return nil, errors.New("time lock is not existing")
	}

	if val.ToTimelock != msg.To {
		return nil, errors.New("not matching receiver address!")
	}

	if val.Blockheight > uint64(ctx.BlockHeight()) {
		return nil, errors.New("wait until valid blokcheight")
	}

	to, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, sdk.Coins{*val.Coinhtlc})
	if err != nil {
		return nil, err
	}
	k.Keeper.RemoveCommitment(ctx, msg.Index)

	return &types.MsgWithdrawTimelockResponse{}, nil
}
