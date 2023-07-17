package board

import (
	"fmt"
	"strings"

	"github.com/shion0625/shogiApp/engine/src/domain/model"
	"github.com/shion0625/shogiApp/engine/src/piece"
)

// Bitboard はビットボードを表す型です。
type Bitboard uint64

// Bitboards は駒ごとのビットボードを保持する構造体です。
type Bitboards [2]Bitboard

// PieceBB は駒ごとの配置情報を保持する構造体です。
type PieceBB struct {
	Black Bitboards // 先手の駒の配置
	White Bitboards // 後手の駒の配置
}

// Board は将棋盤を表す構造体です。
type Board struct {
	Pieces    [model.PieceNum]PieceBB // 駒ごとの配置情報
	Occupied  Bitboards               // 全駒の配置情報
	BlackTurn bool                    // 先手の手番かどうか
}

// SetStartingPositionFromSFEN は初期の駒の配置をSFEN形式の文字列から解析し設定する関数です。
// SFEN 形式は将棋の局面表現形式であり、以下の形式を持ちます: {駒の配置} {手番} {持ち駒} {手数}
// 例: "lnsgkgsnl/1r5b1/p1ppppppp/9/9/9/P1PPPPPPP/1B5R1/LNSGKGSNL b - 1"
// 引数:
//   - sfen: SFEN 形式の文字列
//
// 返り値:
//   - *Board: 初期の駒の配置が設定されたボード
//   - error: エラー情報（解析エラーや不正なフォーマットなど）
func SetStartingPositionFromSFEN(sfen string) (*Board, error) {
	board := &Board{
		Occupied: Bitboards{0, 0}, // 初期化された Bitboards
	}

	splitSfen := strings.Split(sfen, " ")
	if len(splitSfen) != 4 {
		return nil, fmt.Errorf("Invalid SFEN format")
	}

	// SFENを解析して駒の配置を設定
	parts := strings.Split(splitSfen[0], "/")
	if len(parts) != 9 {
		return nil, fmt.Errorf("Invalid SFEN format")
	}

	// 駒ごとに配置情報を設定
	for rankIdx, rank := range parts {
		fileIdx := 0
		if fileIdx >= 9 {
			return nil, fmt.Errorf("Invalid SFEN format")
		}

		for _, char := range rank {
			if char >= '1' && char <= '9' {
				// 数字の場合、その回数だけ空白を処理する
				numSpaces := int(char - '0')
				fileIdx += numSpaces
				continue
			}
			//pieceの情報を取得
			piece, err := piece.GetPieceTypeFromChar(char)
			if err != nil {
				return nil, err
			}

			//先手と後手でデータをそれぞれ設定
			switch piece.Owner {
			case model.Black:
				board.Pieces[piece.Name].Black.SetBit(fileIdx, rankIdx)
			case model.White:
				board.Pieces[piece.Name].White.SetBit(fileIdx, rankIdx)
			}
			//全体のデータを設定
			board.Occupied.SetBit(fileIdx, rankIdx)
			fileIdx++
		}
	}
	return board, nil
}

// SetBit は指定した位置にビットを立てるメソッドです。
// fileIdx はファイル（列）のインデックス、rankIdx はランク（行）のインデックスを表します。
func (bb *Bitboards) SetBit(fileIdx, rankIdx int) {
	index := (fileIdx + rankIdx*9) / 64  // ビットボードのインデックスを計算
	offset := (fileIdx + rankIdx*9) % 64 // ビットのオフセットを計算
	mask := Bitboard(1) << uint(offset)  // オフセットに対応するビットマスクを作成

	if bb[index]&mask != 0 {
		bb[index] ^= mask // ビットが既に立っている場合は反転させます
	} else {
		bb[index] |= mask // ビットが立っていない場合は立てます
	}
}

// String はビットボードの文字列表現を返すメソッドです。
func (bb Bitboard) String() string {
	return fmt.Sprintf("%064b", bb)
}
