package chess

import (
	"testing"
)

func TestMakeMove(t *testing.T) {
	b := StartPosition()
	m := Move{Square{Rank2, FileE}, Square{Rank4, FileE}}
	after := b.MakeMove(m)

	want := makePiece(Pawn, White)
	if got := after.PieceAt(Square{Rank4, FileE}); got != want {
		t.Errorf("e4 = %v, want %v", got, want)
	}
	want = makePiece(Empty, Black)
	if got := after.PieceAt(Square{Rank2, FileE}); got != want {
		t.Errorf("e2 = %v, want %v", got, want)
	}
	if b != StartPosition() {
		t.Errorf("board = %v, want %v", b, StartPosition())
	}
}

// isAttackedCase describes one query. Every position holds pieces of a SINGLE
// colour, so "is this square attacked" has an unambiguous answer without
// needing to name a side - which keeps the cases honest against the current
// one-argument signature.
type isAttackedCase struct {
	name   string
	pieces map[string]Piece
	target string
	by     Colour
	want   bool
}

func TestIsAttacked(t *testing.T) {
	wRook := makePiece(Rook, White)
	bRook := makePiece(Rook, Black)
	wQueen := makePiece(Queen, White)
	bQueen := makePiece(Queen, Black)
	wBishop := makePiece(Bishop, White)
	bBishop := makePiece(Bishop, Black)
	wPawn := makePiece(Pawn, White)
	wKing := makePiece(King, White)
	bKing := makePiece(King, Black)

	cases := []isAttackedCase{
		// Empty target squares. These are the ones a king consults before
		// stepping anywhere, so they are the cases that matter most.
		{"white rook covers the a-file", map[string]Piece{"a1": wRook}, "a8", White, true},
		{"black rook covers the a-file", map[string]Piece{"a1": bRook}, "a8", Black, true},
		{"white queen covers the d-file", map[string]Piece{"d8": wQueen}, "d1", White, true},
		{"black queen covers the d-file", map[string]Piece{"d8": bQueen}, "d1", Black, true},
		{"white bishop covers its diagonal", map[string]Piece{"c8": wBishop}, "a6", White, true},
		{"black bishop covers its diagonal", map[string]Piece{"c8": bBishop}, "a6", Black, true},

		// Occupied target squares.
		{"white rook attacks a black king", map[string]Piece{"a1": wRook, "a8": bKing}, "a8", White, true},
		{"black rook attacks a white king", map[string]Piece{"a1": bRook, "a8": wKing}, "a8", Black, true},

		// Squares nobody covers, and sliders stopping at a blocker.
		{"square off the rook's lines", map[string]Piece{"a1": wRook}, "h8", White, false},
		{"rook blocked by its own piece", map[string]Piece{"a1": wRook, "a4": wPawn}, "a8", White, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			b := place(c.pieces)
			got := b.IsAttacked(sq(c.target), c.by)
			if got != c.want {
				t.Errorf("IsAttacked(%s, %s) = %v, want %v", c.target, c.by, got, c.want)
			}
		})
	}
}
