package chess

type Square struct {
	Rank Rank
	File File
}

func (s Square) Valid() bool {
	return s.Rank >= 0 && s.Rank <= 7 && s.File >= 0 && s.File <= 7
}

func (r Rank) String() string {
	return string(rune('0' + r + 1))
}

func (f File) String() string {
	return string(rune('a' + f))
}
