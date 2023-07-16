package board

type Board struct {
	All [9][9]Pieces
}

func NewBoard() *Board {
	piece := NewPieces()
	board := &Board{
		All: [9][9]Pieces{
			{Pieces{piece.Lance, piece.Knight, piece.Silver, piece.Gold, piece.King, piece.Gold, piece.Silver, piece.Knight, piece.Lance}},
			{Pieces{piece.Empty, piece.Rook, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Bishop, piece.Empty}},
			{Pieces{piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn}},
			{Pieces{piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty}},
			{Pieces{piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty}},
			{Pieces{piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty}},
			{Pieces{piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn}},
			{Pieces{piece.Empty, piece.Bishop, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Rook, piece.Empty}},
			{Pieces{piece.Lance, piece.Knight, piece.Silver, piece.Gold, piece.King, piece.Gold, piece.Silver, piece.Knight, piece.Lance}},
		},
	}
	return board
}
