package ante

import (
	"fmt"

	sdk "github.com/sisu-network/cosmos-sdk/types"
	"github.com/sisu-network/sisu/utils"
	"github.com/sisu-network/sisu/x/tss"
)

type TssDecorator struct {
	validator tss.TssValidator
}

func NewTssDecorator(validator tss.TssValidator) TssDecorator {
	return TssDecorator{
		validator: validator,
	}
}

func (d TssDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if ctx.IsReCheckTx() {
		return next(ctx, tx, simulate)
	}

	utils.LogDebug("Checking TSS transaction")
	msgs := tx.GetMsgs()
	if len(msgs) == 0 {
		return ctx, fmt.Errorf("Empty mesage list")
	}

	if err := d.validator.CheckTx(ctx, msgs); err != nil {
		return ctx, err
	}

	return next(ctx, tx, simulate)
}
