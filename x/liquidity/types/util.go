package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// BulkSendCoinsOperation holds a list of SendCoins operations for bulk execution.
type BulkSendCoinsOperation struct {
	Inputs  []banktypes.Input
	Outputs []banktypes.Output
}

// NewBulkSendCoinsOperation returns an empty BulkSendCoinsOperation.
func NewBulkSendCoinsOperation() *BulkSendCoinsOperation {
	return &BulkSendCoinsOperation{
		Inputs:  []banktypes.Input{},
		Outputs: []banktypes.Output{},
	}
}

// review: SendCoins has the same name as BankKeeper.SendCoins, so it can confuse people, how about refactor to QueueSendCoins or QueueInputOutput.
// SendCoins queues a BankKeeper.SendCoins operation for later execution.
func (op *BulkSendCoinsOperation) SendCoins(fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) {
	if amt.IsValid() && !amt.IsZero() {
		op.Inputs = append(op.Inputs, banktypes.NewInput(fromAddr, amt))
		op.Outputs = append(op.Outputs, banktypes.NewOutput(toAddr, amt))
	}
}

// Run runs BankKeeper.InputOutputCoins once for queued operations.
func (op *BulkSendCoinsOperation) Run(ctx sdk.Context, bankKeeper BankKeeper) error {
	if len(op.Inputs) > 0 && len(op.Outputs) > 0 {
		return bankKeeper.InputOutputCoins(ctx, op.Inputs, op.Outputs)
	}
	return nil
}
