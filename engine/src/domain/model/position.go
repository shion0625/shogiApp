package model

import (
	"fmt"

	"github.com/shion0625/shogiApp/engine/src/utils"
)

type Position struct {
	File int8 // 筋（列）
	Rank int8 // 段（行）
}

func NewPosition(file int8, rank int8) (_ *Position, err error) {
	defer iterrors.Wrap(&err, "model.NewPosition(%d, %d)", file, rank)

	if isValidPosition(file) {
		return nil, fmt.Errorf("The file must be between 0 and 8. ")
	}

	if isValidPosition(rank) {
		return nil, fmt.Errorf("The rank must be between 0 and 8. ")
	}

	return &Position{
		File: file,
		Rank: rank,
	}, nil
}

func isValidPosition(value int8) bool {
	return (value < 0 || 8 < value)
}
