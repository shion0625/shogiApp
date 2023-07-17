package piece

import (
	"fmt"

	"github.com/shion0625/shogiApp/engine/src/domain/model"
)

// 文字から駒の種類を取得するヘルパー関数
func GetPieceTypeFromChar(char rune) (*model.Piece, error) {
	switch char {
	case 'r':
		return model.NewPiece(model.Rook, model.White)
	case 'b':
		return model.NewPiece(model.Bishop, model.White)
	case 'g':
		return model.NewPiece(model.Gold, model.White)
	case 's':
		return model.NewPiece(model.Silver, model.White)
	case 'n':
		return model.NewPiece(model.Knight, model.White)
	case 'l':
		return model.NewPiece(model.Lance, model.White)
	case 'p':
		return model.NewPiece(model.Pawn, model.White)
	case 'k':
		return model.NewPiece(model.King, model.White)
	case 'K':
		return model.NewPiece(model.King, model.Black)
	case 'B':
		return model.NewPiece(model.Bishop, model.Black)
	case 'R':
		return model.NewPiece(model.Rook, model.Black)
	case 'G':
		return model.NewPiece(model.Gold, model.Black)
	case 'S':
		return model.NewPiece(model.Silver, model.Black)
	case 'N':
		return model.NewPiece(model.Knight, model.Black)
	case 'L':
		return model.NewPiece(model.Lance, model.Black)
	case 'P':
		return model.NewPiece(model.Pawn, model.Black)
	default:
		return nil, fmt.Errorf("Invalid SFEN piece character: %s", string(char))
	}
}

func PieceTypeToJapanese(piece model.Piece) string {
	switch piece.Name {
	case model.Empty:
		return "空"
	case model.Pawn:
		return "歩"
	case model.Lance:
		return "香"
	case model.Knight:
		return "桂"
	case model.Bishop:
		return "角"
	case model.Rook:
		return "飛"
	case model.Gold:
		return "金"
	case model.Silver:
		return "銀"
	case model.King:
		return "王"
	default:
		return "不"
	}
}

func PieceTypeToEnglish(piece model.Piece) string {
	switch piece.Name {
	case model.Empty:
		return "Empty"
	case model.Pawn:
		return "Pawn"
	case model.Lance:
		return "Lance"
	case model.Knight:
		return "Knight"
	case model.Bishop:
		return "Bishop"
	case model.Rook:
		return "Rook"
	case model.Gold:
		return "Gold"
	case model.Silver:
		return "Silver"
	case model.King:
		return "King"
	default:
		return "Unknown"
	}
}
