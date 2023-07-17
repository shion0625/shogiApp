package model

import (
	"fmt"

	"github.com/shion0625/shogiApp/engine/src/utils/iterrors"
)

// Owner は駒の所有者を表す型です。
type Owner int8

const (
	NoOwner Owner = iota // 所有者なし
	Black                // 先手
	White                // 後手
)

// NewOwner は指定された所有者を使用して新しい Owner インスタンスを作成します。
// 所有者の値が有効な範囲外の場合はエラーを返します。
// 引数:
//   - owner: 所有者の値
//
// 返り値:
//   - *Owner: 新しい Owner インスタンスへのポインタ
//   - error: エラー情報（所有者の値が範囲外の場合など）
func NewOwner(owner Owner) (_ *Owner, err error) {
	defer iterrors.Wrap(&err, "model.NewPiece(%d)", owner)

	if isValidOwner(owner) {
		return nil, fmt.Errorf("The owner must be between 0 and 2. ")
	}
	return &owner, nil
}

// isValidOwner は与えられた所有者の値が有効かどうかを判定します。
// 引数:
//   - value: 判定する所有者の値
//
// 返り値:
//   - bool: 所有者の値が有効な場合は true、無効な場合は false
func isValidOwner(value Owner) bool {
	return (value < 0 || 2 < value)
}
