package model

import (
	"fmt"

	"github.com/shion0625/shogiApp/engine/src/utils/iterrors"
)

type Owner int8

const (
	NoOwner Owner = iota // 所有者なし
	Black                // 先手
	White                // 後手
)

func NewOwner(owner Owner) (_ *Owner, err error) {
	defer iterrors.Wrap(&err, "model.NewPiece(%d)", owner)

	if isValidOwner(owner) {
		return nil, fmt.Errorf("The owner must be between 0 and 2. ")
	}
	return &owner, nil
}

func isValidOwner(value Owner) bool {
	return (value < 0 || 2 < value)
}
