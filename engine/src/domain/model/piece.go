package model

import (
	"fmt"

	"github.com/shion0625/shogiApp/engine/src/utils/iterrors"
)

// PieceType は駒の種類を表す型です。
type PieceType int8

// Piece は将棋の駒を表す構造体です。
type Piece struct {
	Name  PieceType // コマの名前
	Owner Owner     // コマの位置
}

// 駒の種類の定義
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

// NewPiece は指定された駒の種類と所有者を使用して新しい Piece インスタンスを作成します。
// 駒の種類や所有者の値が有効な範囲外の場合はエラーを返します。
// 引数:
//   - name: 駒の種類
//   - owner: 所有者
//
// 返り値:
//   - *Piece: 新しい Piece インスタンスへのポインタ
//   - error: エラー情報（エラーがない場合は nil）
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

// isValidName は与えられた駒の種類の値が有効かどうかを判定します。
// 引数:
//   - value: 判定する駒の種類の値
//
// 返り値:
//   - bool: 駒の種類の値が有効な場合は true、無効な場合は false
func isValidName(value PieceType) bool {
	return (value < 0 || 8 < value)
}
