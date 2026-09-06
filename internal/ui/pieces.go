package ui

type pieceStyle struct {
	name    string
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

var pieceStyles = []pieceStyle{
	{
		name:    "notation",
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
	{
		name:    "notation-balanced",
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
	{
		name:    "icons",
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
	{
		name:    "icons-outlined",
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
	{
		name:    "icons-filled",
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
