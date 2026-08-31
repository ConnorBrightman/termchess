// Package chess
package chess

import (
	"fmt"
)

type Board struct {
	// [Rank] [File]
	squares [8][8]Piece
}

// PieceAt returns the [Piece] at a square
func (b Board) PieceAt(s Square) Piece {
	if !s.Valid() {
		return makePiece(Empty, Black)
	}
	return b.squares[s.Rank][s.File]
}

// StartPosition makes [Board] with the Pieces set up for a new game
func StartPosition() Board {
	var b Board
	backRank := []PieceType{
		Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook,
	}
	for i := 0; i < 8; i++ {
		b.squares[0][i] = makePiece(backRank[i], White)
		b.squares[1][i] = makePiece(Pawn, White)
	}
	for i := 7; i >= 0; i-- {
		b.squares[7][i] = makePiece(backRank[i], Black)
		b.squares[6][i] = makePiece(Pawn, Black)
	}
	return b
}

// String form of the board
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

// MakeMove moves a [Piece] from one square to another
func (b Board) MakeMove(m Move) Board {
	from := b.squares[m.From.Rank][m.From.File]
	b.squares[m.To.Rank][m.To.File] = from
	b.squares[m.From.Rank][m.From.File] = makePiece(Empty, Black)

	return b
}

// Moves returns the valid squares a piece can move to
func (b Board) Moves(from Square) []Square {
	moves := []Square{}
	p := b.PieceAt(from)

	switch p.PieceType() {
	case Pawn:
		moves = b.pawnMoves(from)
	case Knight:
		moves = b.baseMoves(knightOffsets[:], from)
	case King:
		moves = b.baseMoves(kingOffsets[:], from)
	case Rook:
		moves = b.slideMoves(rookDirs[:], from)
	case Bishop:
		moves = b.slideMoves(bishopDirs[:], from)
	case Queen:
		moves = b.slideMoves(queenDirs[:], from)
	}
	return moves
}

func (b Board) baseMoves(offsets []offset, from Square) []Square {
	moves := []Square{}
	p := b.PieceAt(from)
	for _, o := range offsets {
		to := Square{from.Rank + o.dR, from.File + o.dF}
		if to.Valid() {
			dst := b.PieceAt(to)
			if !dst.IsEmpty() && dst.PieceColour() == p.PieceColour() {
				continue
			}
			moves = append(moves, to)
		}
	}
	return moves
}

func (b Board) slideMoves(offsets []offset, from Square) []Square {
	moves := []Square{}
	p := b.PieceAt(from)
	for _, d := range offsets {
		to := from
		for {
			to = Square{to.Rank + d.dR, to.File + d.dF}
			if !to.Valid() {
				break
			}
			dst := b.PieceAt(to)
			if dst.IsEmpty() {
				moves = append(moves, to)
				continue
			}
			if dst.PieceColour() != p.PieceColour() {
				moves = append(moves, to)
			}
			break
		}
	}
	return moves
}

func (b Board) pawnMoves(from Square) []Square {
	moves := []Square{}
	p := b.PieceAt(from)
	pColor := p.PieceColour()
	dir := Rank(1)
	if pColor == Black {
		dir = -1
	}
	isFirstMove := (from.Rank == Rank2 && pColor == White) || (from.Rank == Rank7 && pColor == Black)

	one := Square{from.Rank + dir, from.File}
	dst := b.PieceAt(one)
	if one.Valid() && dst.IsEmpty() {
		moves = append(moves, one)
		if isFirstMove {
			two := Square{from.Rank + 2*dir, from.File}
			dst = b.PieceAt(two)
			if two.Valid() && dst.IsEmpty() {
				moves = append(moves, two)
			}
		}
	}

	// capture moves
	dirs := []File{1, -1}
	for _, d := range dirs {
		to := Square{from.Rank + dir, from.File + d}
		dst := b.PieceAt(to)
		if !to.Valid() || dst.IsEmpty() || dst.PieceColour() == pColor {
			continue
		}
		moves = append(moves, to)
	}

	return moves
}

type offset struct {
	dR Rank
	dF File
}

var (
	knightOffsets = [...]offset{{1, 2}, {2, 1}, {1, -2}, {-2, 1}, {-1, 2}, {2, -1}, {-1, -2}, {-2, -1}}
	kingOffsets   = [...]offset{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	rookDirs      = [...]offset{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	bishopDirs    = [...]offset{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	queenDirs     = [...]offset{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
)
