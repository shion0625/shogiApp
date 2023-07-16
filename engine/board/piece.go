package board

type Piece int

type Pieces struct {
	Empty  Piece // 空のマス
	Pawn   Piece // 歩兵
	Lance  Piece // 香車
	Knight Piece // 桂馬
	Bishop Piece // 角行
	Rook   Piece // 飛車
	Gold   Piece // 金将
	Silver Piece // 銀将
	King   Piece // 王将
}

const (
	Empty  Piece = iota // 空のマス
	Pawn                // 歩兵
	Lance               // 香車
	Knight              // 桂馬
	Bishop              // 角行
	Rook                // 飛車
	Gold                // 金将
	Silver              // 銀将
	King                // 王将
)

func newPieces() *Pieces {
	return &Pieces{
		Empty:  Empty,
		Pawn:   Pawn,
		Lance:  Lance,
		Knight: Knight,
		Bishop: Bishop,
		Rook:   Rook,
		Gold:   Gold,
		Silver: Silver,
		King:   King,
	}
}
