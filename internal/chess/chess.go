package chess

type Colour uint8

const (
	Black Colour = iota
	White
)

type Rank int

const (
	Rank1 Rank = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
)

type File int

const (
	FileA File = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

func (c Colour) String() string {
	if c == White {
		return "White"
	}
	return "Black"
}

func (c Colour) Opponent() Colour {
	if c == White {
		return Black
	}
	return White
}
