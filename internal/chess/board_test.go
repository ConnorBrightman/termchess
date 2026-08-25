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
