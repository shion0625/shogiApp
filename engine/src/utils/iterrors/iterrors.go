package iterrors

import (
	"fmt"
)

// Wrap はエラーをラップします。エラーが非nilの場合、指定されたフォーマットと引数を使用して新しいエラーを作成し、元のエラーをラップします。
// 引数:
//   - err: エラーポインタ
//   - format: エラーメッセージのフォーマット文字列
//   - args: フォーマット文字列に埋め込む引数
func Wrap(err *error, format string, args ...interface{}) {
	if *err != nil {
		*err = fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), *err)
	}
}
