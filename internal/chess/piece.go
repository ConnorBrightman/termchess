package chess

import "strings"

type PieceType uint8

const (
	empty PieceType = iota
	pawn
	knight
	bishop
	rook
	queen
	king
)

type Piece struct {
	pType   PieceType
	pColour Colour
}

func makePiece(pt PieceType, pc Colour) Piece {
	p := Piece{
		pType: pt, pColour: pc,
	}
	return p
}

func (p Piece) String() string {
	str := ""
	if p.pType == empty {
		return "."
	}
	switch p.pType {
	case pawn:
		str = "P"
	case knight:
		str = "N"
	case bishop:
		str = "B"
	case rook:
		str = "R"
	case queen:
		str = "Q"
	case king:
		str = "K"
	}
	if p.pColour == black {
		str = strings.ToLower(str)
	}

	return str
}
