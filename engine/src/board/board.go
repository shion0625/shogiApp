package board

import (
	"fmt"
	"strings"

	"github.com/shion0625/shogiApp/engine/src/domain/model"
)

type Bitboard uint64

// ビットボードを表す型
type Bitboards [2]Bitboard

// 駒ごとのビットボードを保持する構造体
type PieceBB struct {
	Black Bitboards // 先手の駒の配置
	White Bitboards // 後手の駒の配置
}

// 将棋盤を表す構造体
type Board struct {
	Pieces    [model.PieceNum]PieceBB // 駒ごとの配置情報
	Occupied  Bitboards               // 全駒の配置情報
	BlackTurn bool                    // 先手の手番かどうか
}

// 初期の駒の配置をパースして設定する関数
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
			piece, err := model.GetPieceTypeFromChar(char)
			if err != nil {
				return nil, err
			}

			switch piece.Owner {
			case model.Black:
				board.Pieces[piece.Name].Black.SetBit(fileIdx, rankIdx)
			case model.White:
				board.Pieces[piece.Name].White.SetBit(fileIdx, rankIdx)
			}
			board.Occupied.SetBit(fileIdx, rankIdx)
			fileIdx++
		}
	}
	return board, nil
}

// 指定した位置にビットを立てるメソッド
func (bb *Bitboards) SetBit(fileIdx, rankIdx int) {
	index := (fileIdx + rankIdx*9) / 64
	offset := (fileIdx + rankIdx*9) % 64
	mask := Bitboard(1) << uint(offset)
	if bb[index]&mask != 0 {
		bb[index] ^= mask // ビットが既に立っている場合は反転
	} else {
		bb[index] |= mask
	}
}

func (bb Bitboard) String() string {
	return fmt.Sprintf("%064b", bb)
}
