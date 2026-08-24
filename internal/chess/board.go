// Package chess
package chess

import (
	"fmt"
)

type Board struct {
	// [rank][file]
	squares [8][8]Piece
}

func StartPosition() Board {
	var b Board
	backRank := []PieceType{
		rook, knight, bishop, queen, king, bishop, knight, rook,
	}
	for i := 0; i < 8; i++ {
		b.squares[0][i] = makePiece(backRank[i], white)
		b.squares[1][i] = makePiece(pawn, white)
	}
	for i := 7; i >= 0; i-- {
		b.squares[7][i] = makePiece(backRank[i], black)
		b.squares[6][i] = makePiece(pawn, black)
	}
	return b
}

func (b Board) String() string {
	str := ""
	for i := 7; i >= 0; i-- {
		str += fmt.Sprintf("\n")
		for j := 0; j < 8; j++ {
			str += fmt.Sprintf("%v", b.squares[i][j])
		}
	}

	return str
}
