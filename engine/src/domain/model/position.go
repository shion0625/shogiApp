package model

import (
	"fmt"

	"github.com/shion0625/shogiApp/engine/src/utils/iterrors"
)

// Position は将棋盤上の位置を表す構造体です。
type Position struct {
	File int8 // 筋（列）
	Rank int8 // 段（行）
}

// NewPosition は指定された筋と段を使用して新しい Position インスタンスを作成します。
// 筋や段の値が有効な範囲外の場合はエラーを返します。
// 引数:
//   - file: チェックする筋の値
//   - rank: チェックする段の値
//
// 返り値:
//   - *Position: 新しい Position インスタンスへのポインタ
//   - error: エラー情報（有効な範囲外の場合はエラーメッセージ）
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

// isValidPosition は与えられた位置の値が有効かどうかを判定します。
// 引数:
//   - value: チェックする位置の値
//
// 返り値:
//   - bool: 有効な位置の場合は true、それ以外の場合は false
func isValidPosition(value int8) bool {
	return (value < 0 || 8 < value)
}
