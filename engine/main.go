package main

import (
	"fmt"

	"github.com/shion0625/shogiApp/engine/src/board"
	// "github.com/shion0625/shogiApp/engine/src/domain/model"
)

func main() {
	sfen := "lnsgkgsnl/1r5b1/ppppppppp/9/9/9/PPPPPPPPP/1B5R1/LNSGKGSNL b - 1"
	board, err := board.SetStartingPositionFromSFEN(sfen)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// ビットボードを使った処理などを行うことができます
	fmt.Println(board.Occupied)
}
