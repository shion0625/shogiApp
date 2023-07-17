package model

import (
	"fmt"

	"github.com/shion0625/shogiApp/engine/src/utils/iterrors"
)

type PieceType int8

type Piece struct {
	Name PieceType // コマの名前
	// Position *Position // コマの位置
	Owner Owner // コマの位置
}

// コマの列挙
const (
	Empty    PieceType = iota // 空のマス
	Pawn                      // 歩兵
	Lance                     // 香車
	Knight                    // 桂馬
	Bishop                    // 角行
	Rook                      // 飛車
	Gold                      // 金将
	Silver                    // 銀将
	King                      // 王将
	PieceNum                  // 駒の種類の総数
)

func NewPiece(name PieceType, owner Owner) (_ *Piece, err error) {
	defer iterrors.Wrap(&err, "model.NewPiece(%d, %d)", name, owner)

	if isValidName(name) {
		return nil, fmt.Errorf("The name name be between 0 and 8. ")
	}

	new_owner, err := NewOwner(owner)
	if err != nil {
		return nil, err
	}

	return &Piece{
		Name:  name,
		Owner: *new_owner,
	}, nil
}

func isValidName(value PieceType) bool {
	return (value < 0 || 8 < value)
}
