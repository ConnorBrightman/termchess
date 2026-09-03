package chess

type Move struct {
	From Square
	To   Square
}

type MoveType struct {
	IncludeOwn bool
}

var (
	captureMove = MoveType{IncludeOwn: false}
	attackMove  = MoveType{IncludeOwn: true}
)
