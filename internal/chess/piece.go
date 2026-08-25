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
