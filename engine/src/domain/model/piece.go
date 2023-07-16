package model

import (
	"fmt"

	"github.com/shion0625/shogiApp/engine/src/utils"
)

type PieceName int8

type Piece struct {
	Name     PieceName // コマの名前
	Position Position  // コマの位置
	Turn     Turn      // コマの位置
}

// コマの列挙
const (
	Empty  PieceName = iota // 空のマス
	Pawn                    // 歩兵
	Lance                   // 香車
	Knight                  // 桂馬
	Bishop                  // 角行
	Rook                    // 飛車
	Gold                    // 金将
	Silver                  // 銀将
	King                    // 王将
)

func NewPiece(name PieceName, turn Turn, file int8, rank int8) (_ *Piece, err error) {
	defer iterrors.Wrap(&err, "model.NewPiece(%d, %d, %d, %d)", name, turn, file, rank)

	if isValidName(name) {
		return nil, fmt.Errorf("The name name be between 0 and 8. ")
	}

	new_turn, err := NewTurn(turn)
	if err != nil {
		return nil, err
	}

	new_position, err := NewPosition(file, rank)
	if err != nil {
		return nil, err
	}

	return &Piece{
		Name:     name,
		Turn:     *new_turn,
		Position: *new_position,
	}, nil
}

func isValidName(value PieceName) bool {
	return (value < 0 || 8 < value)
}
