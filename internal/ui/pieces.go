package ui

type pieceStyle struct {
	wPawn   string
	wKnight string
	wBishop string
	wRook   string
	wQueen  string
	wKing   string
	bPawn   string
	bKnight string
	bBishop string
	bRook   string
	bQueen  string
	bKing   string
}

var pieceStyles = map[string]pieceStyle{
	"notation": {
		wPawn:   "P",
		wKnight: "N",
		wBishop: "B",
		wRook:   "R",
		wQueen:  "Q",
		wKing:   "K",
		bPawn:   "p",
		bKnight: "n",
		bBishop: "b",
		bRook:   "r",
		bQueen:  "q",
		bKing:   "k",
	},
	"notation-matched": {
		wPawn:   "P",
		wKnight: "N",
		wBishop: "B",
		wRook:   "R",
		wQueen:  "Q",
		wKing:   "K",
		bPawn:   "P",
		bKnight: "N",
		bBishop: "B",
		bRook:   "R",
		bQueen:  "Q",
		bKing:   "K",
	},
	"icons": {
		wPawn:   "♙",
		wKnight: "♘",
		wBishop: "♗",
		wRook:   "♖",
		wQueen:  "♕",
		wKing:   "♔",
		bPawn:   "♟",
		bKnight: "♞",
		bBishop: "♝",
		bRook:   "♜",
		bQueen:  "♛",
		bKing:   "♚",
	},
	"icons-outlined": {
		wPawn:   "♙",
		wKnight: "♘",
		wBishop: "♗",
		wRook:   "♖",
		wQueen:  "♕",
		wKing:   "♔",
		bPawn:   "♙",
		bKnight: "♘",
		bBishop: "♗",
		bRook:   "♖",
		bQueen:  "♕",
		bKing:   "♔",
	},
	"icons-filled": {
		wPawn:   "♟",
		wKnight: "♞",
		wBishop: "♝",
		wRook:   "♜",
		wQueen:  "♛",
		wKing:   "♚",
		bPawn:   "♟",
		bKnight: "♞",
		bBishop: "♝",
		bRook:   "♜",
		bQueen:  "♛",
		bKing:   "♚",
	},
}
