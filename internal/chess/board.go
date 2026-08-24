// Package chess
package chess

import (
	"fmt"
)

type Board struct {
	// [rank][file]
	Squares [8][8]Piece
}

func StartPosition() Board {
	var b Board
	backRank := []PieceType{
		rook, knight, bishop, queen, king, bishop, knight, rook,
	}
	for i := 0; i < 8; i++ {
		b.Squares[0][i] = makePiece(backRank[i], white)
		b.Squares[1][i] = makePiece(pawn, white)
	}
	for i := 7; i >= 0; i-- {
		b.Squares[7][i] = makePiece(backRank[i], black)
		b.Squares[6][i] = makePiece(pawn, black)
	}
	return b
}

func (b Board) String() string {
	str := ""
	for i := 7; i >= 0; i-- {
		str += fmt.Sprintf("\n")
		for j := 0; j < 8; j++ {
			str += fmt.Sprintf("%v", b.Squares[i][j])
		}
	}

	return str
}
