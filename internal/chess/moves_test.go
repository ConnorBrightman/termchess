package chess

import (
	"sort"
	"strings"
	"testing"
)

// Test helpers. sq and sqName are stand-ins for ParseSquare and Square.String();
// once those exist on Square itself, these should go away.

// sq parses algebraic notation, e.g. "e4", into a Square.
func sq(n string) Square {
	return Square{Rank: Rank(n[1] - '1'), File: File(n[0] - 'a')}
}

// sqName renders a Square as algebraic notation, e.g. "e4".
func sqName(s Square) string {
	return string(rune('a'+s.File)) + string(rune('1'+s.Rank))
}

// place builds a board containing only the given pieces, keyed by square name.
// Everything else is left empty, which keeps each case to the pieces it is about.
func place(pieces map[string]Piece) Board {
	var b Board
	for n, p := range pieces {
		s := sq(n)
		b.squares[s.Rank][s.File] = p
	}
	return b
}

// movesFrom returns the destinations available to the piece on `from`, sorted
// and space separated, so a case reads as one string.
func movesFrom(b Board, from string) string {
	out := []string{}
	for _, s := range b.Moves(sq(from)) {
		out = append(out, sqName(s))
	}
	sort.Strings(out)
	return strings.Join(out, " ")
}

type genCase struct {
	name  string
	board Board
	from  string
	want  string
}

func runCases(t *testing.T, cases []genCase) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := movesFrom(c.board, c.from)
			if got != c.want {
				t.Errorf("moves from %s:\n got  %q (%d)\n want %q (%d)",
					c.from, got, len(strings.Fields(got)),
					c.want, len(strings.Fields(c.want)))
			}
		})
	}
}

var (
	wPawn   = makePiece(Pawn, White)
	bPawn   = makePiece(Pawn, Black)
	wKnight = makePiece(Knight, White)
	wBishop = makePiece(Bishop, White)
	wRook   = makePiece(Rook, White)
	wQueen  = makePiece(Queen, White)
	wKing   = makePiece(King, White)
	bKing   = makePiece(King, Black)
)

func TestKnightMoves(t *testing.T) {
	start := StartPosition()
	runCases(t, []genCase{
		{"b1 at start", start, "b1", "a3 c3"},
		{"g1 at start", start, "g1", "f3 h3"},
		{"b8 at start (black)", start, "b8", "a6 c6"},

		{"centre of empty board", place(map[string]Piece{"d4": wKnight}), "d4",
			"b3 b5 c2 c6 e2 e6 f3 f5"},

		{"corner a1", place(map[string]Piece{"a1": wKnight}), "a1", "b3 c2"},
		{"corner h1", place(map[string]Piece{"h1": wKnight}), "h1", "f2 g3"},
		{"corner a8", place(map[string]Piece{"a8": wKnight}), "a8", "b6 c7"},
		{"corner h8", place(map[string]Piece{"h8": wKnight}), "h8", "f7 g6"},

		{"own piece blocks landing square",
			place(map[string]Piece{"d4": wKnight, "e6": wPawn}), "d4",
			"b3 b5 c2 c6 e2 f3 f5"},
		{"enemy piece is a capture",
			place(map[string]Piece{"d4": wKnight, "e6": bPawn}), "d4",
			"b3 b5 c2 c6 e2 e6 f3 f5"},
	})
}

func TestKingMoves(t *testing.T) {
	start := StartPosition()
	runCases(t, []genCase{
		{"e1 at start is boxed in", start, "e1", ""},
		{"e8 at start is boxed in (black)", start, "e8", ""},

		{"centre of empty board", place(map[string]Piece{"d4": wKing}), "d4",
			"c3 c4 c5 d3 d5 e3 e4 e5"},

		{"corner a1", place(map[string]Piece{"a1": wKing}), "a1", "a2 b1 b2"},
		{"corner h8 (black)", place(map[string]Piece{"h8": bKing}), "h8", "g7 g8 h7"},
		{"edge a4", place(map[string]Piece{"a4": wKing}), "a4", "a3 a5 b3 b4 b5"},

		{"own piece blocks",
			place(map[string]Piece{"d4": wKing, "d5": wPawn}), "d4",
			"c3 c4 c5 d3 e3 e4 e5"},
		{"enemy piece is a capture",
			place(map[string]Piece{"d4": wKing, "d5": bPawn}), "d4",
			"c3 c4 c5 d3 d5 e3 e4 e5"},
	})
}

func TestRookMoves(t *testing.T) {
	start := StartPosition()
	runCases(t, []genCase{
		{"centre of empty board", place(map[string]Piece{"d4": wRook}), "d4",
			"a4 b4 c4 d1 d2 d3 d5 d6 d7 d8 e4 f4 g4 h4"},
		{"corner of empty board", place(map[string]Piece{"a1": wRook}), "a1",
			"a2 a3 a4 a5 a6 a7 a8 b1 c1 d1 e1 f1 g1 h1"},

		// The bug this guards: a rook must not slide through its own pawn.
		{"a1 at start is boxed in", start, "a1", ""},
		{"h1 at start is boxed in", start, "h1", ""},

		{"stops before own piece",
			place(map[string]Piece{"d4": wRook, "d6": wPawn}), "d4",
			"a4 b4 c4 d1 d2 d3 d5 e4 f4 g4 h4"},
		{"captures enemy then stops",
			place(map[string]Piece{"d4": wRook, "d6": bPawn}), "d4",
			"a4 b4 c4 d1 d2 d3 d5 d6 e4 f4 g4 h4"},
		{"cannot pass through the piece it captures",
			place(map[string]Piece{"a6": wRook, "a7": bPawn}), "a6",
			"a1 a2 a3 a4 a5 a7 b6 c6 d6 e6 f6 g6 h6"},

		{"blocked on all four sides", place(map[string]Piece{
			"d4": wRook, "c4": wPawn, "e4": wPawn, "d3": wPawn, "d5": wPawn,
		}), "d4", ""},
	})
}

func TestBishopMoves(t *testing.T) {
	start := StartPosition()
	runCases(t, []genCase{
		{"centre of empty board", place(map[string]Piece{"d4": wBishop}), "d4",
			"a1 a7 b2 b6 c3 c5 e3 e5 f2 f6 g1 g7 h8"},
		{"corner of empty board", place(map[string]Piece{"a1": wBishop}), "a1",
			"b2 c3 d4 e5 f6 g7 h8"},

		{"c1 at start is boxed in", start, "c1", ""},
		{"f8 at start is boxed in (black)", start, "f8", ""},

		{"stops before own piece",
			place(map[string]Piece{"d4": wBishop, "f6": wPawn}), "d4",
			"a1 a7 b2 b6 c3 c5 e3 e5 f2 g1"},
		{"captures enemy then stops",
			place(map[string]Piece{"d4": wBishop, "f6": bPawn}), "d4",
			"a1 a7 b2 b6 c3 c5 e3 e5 f2 f6 g1"},
	})
}

func TestQueenMoves(t *testing.T) {
	start := StartPosition()
	runCases(t, []genCase{
		// A queen is exactly a rook plus a bishop from the same square.
		{"centre of empty board", place(map[string]Piece{"d4": wQueen}), "d4",
			"a1 a4 a7 b2 b4 b6 c3 c4 c5 d1 d2 d3 d5 d6 d7 d8 " +
				"e3 e4 e5 f2 f4 f6 g1 g4 g7 h4 h8"},
		{"corner of empty board", place(map[string]Piece{"a1": wQueen}), "a1",
			"a2 a3 a4 a5 a6 a7 a8 b1 b2 c1 c3 d1 d4 e1 e5 f1 f6 g1 g7 h1 h8"},

		{"d1 at start is boxed in", start, "d1", ""},
		{"d8 at start is boxed in (black)", start, "d8", ""},

		{"blocked on one ray only",
			place(map[string]Piece{"d4": wQueen, "d6": wPawn}), "d4",
			"a1 a4 a7 b2 b4 b6 c3 c4 c5 d1 d2 d3 d5 " +
				"e3 e4 e5 f2 f4 f6 g1 g4 g7 h4 h8"},
	})
}

// Pieces with no generator yet return nothing. This documents the gap and will
// fail loudly once pawns are implemented, as a reminder to write their cases.
func TestPawnMovesNotImplemented(t *testing.T) {
	start := StartPosition()
	if got := movesFrom(start, "e2"); got != "" {
		t.Errorf("pawn generation appears implemented (e2 -> %q) — write real cases", got)
	}
}
