package board

type Board struct {
	All [9][9]Piece
}

func NewBoard() *Board {
	// piece := NewPiece()
	board := &Board{
		All: [9][9]Piece{
			// {piece.Lance, piece.Knight, piece.Silver, piece.Gold, piece.King, piece.Gold, piece.Silver, piece.Knight, piece.Lance},
			// {piece.Empty, piece.Rook, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Bishop, piece.Empty},
			// {piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn},
			// {piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty},
			// {piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty},
			// {piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty},
			// {piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn, piece.Pawn},
			// {piece.Empty, piece.Bishop, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Empty, piece.Rook, piece.Empty},
			// {piece.Lance, piece.Knight, piece.Silver, piece.Gold, piece.King, piece.Gold, piece.Silver, piece.Knight, piece.Lance},
		},
	}
	return board
}
