package model

import (
	"fmt"

	"github.com/shion0625/shogiApp/engine/src/utils"
)

type PieceType int8

type Piece struct {
	Name PieceType // コマの名前
	// Position *Position // コマの位置
	Owner Owner // コマの位置
}

// コマの列挙
const (
	Empty    PieceType = iota // 空のマス
	Pawn                      // 歩兵
	Lance                     // 香車
	Knight                    // 桂馬
	Bishop                    // 角行
	Rook                      // 飛車
	Gold                      // 金将
	Silver                    // 銀将
	King                      // 王将
	PieceNum                  // 駒の種類の総数
)

func NewPiece(name PieceType, owner Owner) (_ *Piece, err error) {
	defer iterrors.Wrap(&err, "model.NewPiece(%d, %d)", name, owner)

	if isValidName(name) {
		return nil, fmt.Errorf("The name name be between 0 and 8. ")
	}

	new_owner, err := NewOwner(owner)
	if err != nil {
		return nil, err
	}

	return &Piece{
		Name:  name,
		Owner: *new_owner,
	}, nil
}

func isValidName(value PieceType) bool {
	return (value < 0 || 8 < value)
}

// 文字から駒の種類を取得するヘルパー関数
func GetPieceTypeFromChar(char rune) (*Piece, error) {
	switch char {
	case 'r':
		return NewPiece(Rook, White)
	case 'b':
		return NewPiece(Bishop, White)
	case 'g':
		return NewPiece(Gold, White)
	case 's':
		return NewPiece(Silver, White)
	case 'n':
		return NewPiece(Knight, White)
	case 'l':
		return NewPiece(Lance, White)
	case 'p':
		return NewPiece(Pawn, White)
	case 'k':
		return NewPiece(King, White)
	case 'K':
		return NewPiece(King, Black)
	case 'B':
		return NewPiece(Bishop, Black)
	case 'R':
		return NewPiece(Rook, Black)
	case 'G':
		return NewPiece(Gold, Black)
	case 'S':
		return NewPiece(Silver, Black)
	case 'N':
		return NewPiece(Knight, Black)
	case 'L':
		return NewPiece(Lance, Black)
	case 'P':
		return NewPiece(Pawn, Black)
	default:
		empty, _ := NewPiece(Empty, NoOwner)
		return empty, fmt.Errorf("Invalid SFEN piece character: %s", string(char))
	}
}

func PieceTypeToJapanese(piece Piece) string {
	switch piece.Name {
	case Empty:
		return "空"
	case Pawn:
		return "歩"
	case Lance:
		return "香"
	case Knight:
		return "桂"
	case Bishop:
		return "角"
	case Rook:
		return "飛"
	case Gold:
		return "金"
	case Silver:
		return "銀"
	case King:
		return "王"
	default:
		return "不"
	}
}

func PieceTypeToEnglish(piece Piece) string {
	switch piece.Name {
	case Empty:
		return "Empty"
	case Pawn:
		return "Pawn"
	case Lance:
		return "Lance"
	case Knight:
		return "Knight"
	case Bishop:
		return "Bishop"
	case Rook:
		return "Rook"
	case Gold:
		return "Gold"
	case Silver:
		return "Silver"
	case King:
		return "King"
	default:
		return "Unknown"
	}
}
