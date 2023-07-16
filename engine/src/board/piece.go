package board

import ()

type PieceName int8



type Piece struct {
	Name     PieceName // コマの名前
	Position Position  // コマの位置
	Turn     Turn      // コマの位置
	// 他のコマの属性を追加
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

func NewPiece(name PieceName, turn Turn, file int8, rank int8) *Piece {
	return &Piece{
		Name: name,
		Turn: turn,
		Position: Position{
			File: file,
			Rank: rank,
		},
	}
}
