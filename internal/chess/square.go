package chess

type Square struct {
	Rank Rank
	File File
}

func (s Square) Valid() bool {
	return s.Rank >= 0 && s.Rank <= 7 && s.File >= 0 && s.File <= 7
}
