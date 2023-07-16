package model

import (
	"fmt"

	"github.com/shion0625/shogiApp/engine/src/utils"
)

type Turn int

const (
	SENTE     Turn = iota //先手
	GOTE                  //後手
	TEBAN_NUM             //手数
)

func NewTurn(turn Turn) (_ *Turn, err error) {
	defer iterrors.Wrap(&err, "model.NewPiece(%d)", turn)

	if isValidTurn(turn) {
		return nil, fmt.Errorf("The turn must be between 0 and 2. ")
	}
	return &turn, nil
}

func isValidTurn(value Turn) bool {
	return (value < 0 || 2 < value)
}
