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
	mvs := []Square{}
	me := b.PieceAt(from).PieceColour()
	for _, to := range b.pseudoMoves(from) {
		after := b.MakeMove(Move{from, to})
		if !after.IsAttacked(after.kingSquare(me), me.Opponent()) {
			mvs = append(mvs, to)
		}
	}
	return mvs
}

func (b Board) kingSquare(colour Colour) Square {
	for r := Rank(0); r < 8; r++ {
		for f := File(0); f < 8; f++ {
			sq := Square{r, f}
			if b.PieceAt(sq).PieceType() == King && b.PieceAt(sq).PieceColour() == colour {
				return sq
			}
		}
	}
	// needs changing to be safer at some point
	return Square{9, 9}
}

// returns the valid squares a piece can move to
func (b Board) pseudoMoves(from Square) []Square {
	if b.PieceAt(from).PieceType() == Pawn {
		return b.pawnMoves(from)
	}
	return b.generate(from, captureMove)
}

// returns the squares a Piece controls
func (b Board) attacks(from Square) []Square {
	if b.PieceAt(from).PieceType() == Pawn {
		return b.pawnAttacks(from)
	}
	return b.generate(from, attackMove)
}

func (b Board) pawnAttacks(from Square) []Square {
	attacks := []Square{}
	p := b.PieceAt(from)
	pColor := p.PieceColour()
	dir := Rank(1)
	if pColor == Black {
		dir = -1
	}
	sides := []int{1, -1}
	for _, side := range sides {
		to := Square{from.Rank + dir, from.File + File(side)}
		if to.Valid() {
			attacks = append(attacks, to)
		}
	}

	return attacks
}

func (b Board) baseMoves(offsets []offset, from Square, mvType MoveType) []Square {
	moves := []Square{}
	p := b.PieceAt(from)
	for _, o := range offsets {
		to := Square{from.Rank + o.dR, from.File + o.dF}
		if to.Valid() {
			dst := b.PieceAt(to)
			if !mvType.IncludeOwn && !dst.IsEmpty() && dst.PieceColour() == p.PieceColour() {
				continue
			}
			moves = append(moves, to)
		}
	}
	return moves
}

func (b Board) slideMoves(offsets []offset, from Square, mvType MoveType) []Square {
	moves := []Square{}
	p := b.PieceAt(from)
	incOwn := mvType.IncludeOwn
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
			if incOwn || dst.PieceColour() != p.PieceColour() {
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

func (b Board) IsAttacked(sq Square, by Colour) bool {
	for r := Rank(0); r < 8; r++ {
		for f := File(0); f < 8; f++ {
			from := Square{r, f}
			p := b.PieceAt(from)
			pt := p.PieceType()
			pc := p.PieceColour()
			// if the attacking square is not empty and is of a different colour
			// check to see if sq is attacked
			if pt != Empty && pc == by {
				for _, mv := range b.attacks(from) {
					if sq == mv {
						return true
					}
				}
			}
		}
	}
	return false
}

func (b Board) generate(from Square, mvType MoveType) []Square {
	r, ok := rules[b.PieceAt(from).PieceType()]
	if !ok {
		return nil // Empty, or Pawn — handled by the caller
	}
	if r.slides {
		return b.slideMoves(r.offsets, from, mvType)
	}
	return b.baseMoves(r.offsets, from, mvType)
}
