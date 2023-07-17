package piece

import (
	"fmt"
	"testing"

	"github.com/shion0625/shogiApp/engine/src/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestGetPieceTypeFromChar(t *testing.T) {
	testCases := []struct {
		text     string
		char     rune
		expected *model.Piece
		err      error
	}{
		{"[正常系]Testing Rook(White)", 'r', func() *model.Piece {
			p, _ := model.NewPiece(model.Rook, model.White)
			return p
		}(), nil},
		{"[正常系]Testing Bishop(White)", 'b', func() *model.Piece {
			p, _ := model.NewPiece(model.Bishop, model.White)
			return p
		}(), nil},
		{"[正常系]Testing Gold(White)", 'g', func() *model.Piece {
			p, _ := model.NewPiece(model.Gold, model.White)
			return p
		}(), nil},
		{"[正常系]Testing Silver(White)", 's', func() *model.Piece {
			p, _ := model.NewPiece(model.Silver, model.White)
			return p
		}(), nil},
		{"[正常系]Testing Knight(White)", 'n', func() *model.Piece {
			p, _ := model.NewPiece(model.Knight, model.White)
			return p
		}(), nil},
		{"[正常系]Testing Lance(White)", 'l', func() *model.Piece {
			p, _ := model.NewPiece(model.Lance, model.White)
			return p
		}(), nil},
		{"[正常系]Testing Pawn(White)", 'p', func() *model.Piece {
			p, _ := model.NewPiece(model.Pawn, model.White)
			return p
		}(), nil},
		{"[正常系]Testing King(White)", 'k', func() *model.Piece {
			p, _ := model.NewPiece(model.King, model.White)
			return p
		}(), nil},
		{"[正常系]Testing King(Black)", 'K', func() *model.Piece {
			p, _ := model.NewPiece(model.King, model.Black)
			return p
		}(), nil},
		{"[正常系]Testing Bishop(Black)", 'B', func() *model.Piece {
			p, _ := model.NewPiece(model.Bishop, model.Black)
			return p
		}(), nil},
		{"[正常系]Testing rook(Black)", 'R', func() *model.Piece {
			p, _ := model.NewPiece(model.Rook, model.Black)
			return p
		}(), nil},
		{"[正常系]Testing Gold(Black)", 'G', func() *model.Piece {
			p, _ := model.NewPiece(model.Gold, model.Black)
			return p
		}(), nil},
		{"[正常系]Testing Silver(Black)", 'S', func() *model.Piece {
			p, _ := model.NewPiece(model.Silver, model.Black)
			return p
		}(), nil},
		{"[正常系]Testing Knight(Black)", 'N', func() *model.Piece {
			p, _ := model.NewPiece(model.Knight, model.Black)
			return p
		}(), nil},
		{"[正常系]Testing Lance(Black)", 'L', func() *model.Piece {
			p, _ := model.NewPiece(model.Lance, model.Black)
			return p
		}(), nil},
		{"[正常系]Testing Pawn(Black)", 'P', func() *model.Piece {
			p, _ := model.NewPiece(model.Pawn, model.Black)
			return p
		}(), nil},
		{"[異常系] Insert a value that has not been set.", 'Z', nil, fmt.Errorf("Invalid SFEN piece character: %s", string('Z'))},
	}

	for _, testCase := range testCases {
		piece, err := GetPieceTypeFromChar(testCase.char)
		assert.Equal(t, testCase.expected, piece)
		assert.Equal(t, testCase.err, err)
		t.Log(testCase.text) // テストケースのテキストをログに表示する
	}
}

func TestPieceTypeToJapanese(t *testing.T) {
	testCases := []struct {
		piece    model.Piece
		expected string
	}{
		{model.Piece{Name: model.Empty}, "空"},
		{model.Piece{Name: model.Pawn}, "歩"},
		{model.Piece{Name: model.Lance}, "香"},
		{model.Piece{Name: model.Knight}, "桂"},
		{model.Piece{Name: model.Bishop}, "角"},
		{model.Piece{Name: model.Rook}, "飛"},
		{model.Piece{Name: model.Gold}, "金"},
		{model.Piece{Name: model.Silver}, "銀"},
		{model.Piece{Name: model.King}, "王"},
		{model.Piece{Name: 15}, "不"},
	}

	for _, testCase := range testCases {
		result := PieceTypeToJapanese(testCase.piece)
		assert.Equal(t, testCase.expected, result)
	}
}

func TestPieceTypeToEnglish(t *testing.T) {
	testCases := []struct {
		piece    model.Piece
		expected string
	}{
		{model.Piece{Name: model.Empty}, "Empty"},
		{model.Piece{Name: model.Pawn}, "Pawn"},
		{model.Piece{Name: model.Lance}, "Lance"},
		{model.Piece{Name: model.Knight}, "Knight"},
		{model.Piece{Name: model.Bishop}, "Bishop"},
		{model.Piece{Name: model.Rook}, "Rook"},
		{model.Piece{Name: model.Gold}, "Gold"},
		{model.Piece{Name: model.Silver}, "Silver"},
		{model.Piece{Name: model.King}, "King"},
		{model.Piece{Name: 20}, "Unknown"},
	}

	for _, testCase := range testCases {
		result := PieceTypeToEnglish(testCase.piece)
		assert.Equal(t, testCase.expected, result)
	}
}
