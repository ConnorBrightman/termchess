package chess

import "strings"

type PieceType uint8

const (
	Empty PieceType = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

type Piece struct {
	pType   PieceType
	pColour Colour
}

type pieceRules struct {
	offsets []offset
	slides  bool
}

var rules = map[PieceType]pieceRules{
	Knight: {knightOffsets[:], false},
	King:   {kingOffsets[:], false},
	Rook:   {rookDirs[:], true},
	Bishop: {bishopDirs[:], true},
	Queen:  {queenDirs[:], true},
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

func makePiece(pt PieceType, pc Colour) Piece {
	p := Piece{
		pType: pt, pColour: pc,
	}
	return p
}

func (p Piece) PieceType() PieceType {
	return p.pType
}

func (p Piece) PieceColour() Colour {
	return p.pColour
}

func (p Piece) IsEmpty() bool {
	return p.pType == Empty
}

func (p Piece) String() string {
	str := ""
	if p.pType == Empty {
		return "."
	}
	switch p.pType {
	case Pawn:
		str = "P"
	case Knight:
		str = "N"
	case Bishop:
		str = "B"
	case Rook:
		str = "R"
	case Queen:
		str = "Q"
	case King:
		str = "K"
	}
	if p.pColour == Black {
		str = strings.ToLower(str)
	}

	return str
}
