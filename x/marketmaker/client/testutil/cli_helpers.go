package testutil

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var commonArgs = []string{
	fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
	fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
}

//// MsgCreateFixedAmountPlanExec creates a transaction for creating a private fixed amount plan.
//func MsgCreateFixedAmountPlanExec(clientCtx client.Context, from string, file string,
//	extraArgs ...string) (testutil.BufferWriter, error) {
//
//	args := append([]string{
//		file,
//		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
//	}, commonArgs...)
//
//	args = append(args, commonArgs...)
//
//	return clitestutil.ExecTestCLICmd(clientCtx, cli.NewCreateFixedAmountPlanCmd(), args)
//}
//
//// MsgStakeExec creates a transaction for staking coin.
//func MsgStakeExec(clientCtx client.Context, from string, stakingCoins string,
//	extraArgs ...string) (testutil.BufferWriter, error) {
//
//	args := append([]string{
//		stakingCoins,
//		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
//	}, commonArgs...)
//
//	args = append(args, commonArgs...)
//
//	return clitestutil.ExecTestCLICmd(clientCtx, cli.NewStakeCmd(), args)
//}
//
//// MsgAdvanceEpochExec creates a transaction to advance epoch by 1.
//func MsgAdvanceEpochExec(clientCtx client.Context, from string,
//	extraArgs ...string) (testutil.BufferWriter, error) {
//
//	args := append([]string{
//		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
//	}, commonArgs...)
//
//	args = append(args, commonArgs...)
//
//	return clitestutil.ExecTestCLICmd(clientCtx, cli.NewAdvanceEpochCmd(), args)
//}
//
//// MsgSendExec creates a transaction to transfer coins.
//func MsgSendExec(clientCtx client.Context, from string, to string, amount string,
//	extraArgs ...string) (testutil.BufferWriter, error) {
//
//	args := append([]string{
//		from,
//		to,
//		amount,
//	}, commonArgs...)
//
//	args = append(args, commonArgs...)
//
//	return clitestutil.ExecTestCLICmd(clientCtx, bankcli.NewSendTxCmd(), args)
//}
