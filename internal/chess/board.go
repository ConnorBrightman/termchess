// Package chess
package chess

import (
	"fmt"
)

type Square struct {
	Rank Rank
	File File
}

type Board struct {
	// [Rank] [File]
	squares [8][8]Piece
}

func (b Board) PieceAt(s Square) Piece {
	return b.squares[s.Rank][s.File]
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
		str += "\n"
		for j := 0; j < 8; j++ {
			str += fmt.Sprintf("%v", b.squares[i][j])
		}
	}

	return str
}

func (b Board) MakeMove(m Move) Board {
	from := b.squares[m.From.Rank][m.From.File]
	b.squares[m.To.Rank][m.To.File] = from
	b.squares[m.From.Rank][m.From.File] = makePiece(empty, black)

	return b
}
