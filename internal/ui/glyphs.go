package ui

type glyphSet struct {
	sqDarkEmpty  string
	sqLightEmpty string
	squareChars  map[SQType]struct{ open, close string }
}

var glyphTypes = map[string]glyphSet{
	"notation": {
		sqDarkEmpty:  " ",
		sqLightEmpty: " ",
		squareChars: map[SQType]struct{ open, close string }{
			sqCursor:   {"<", ">"},
			sqSelected: {"(", ")"},
			sqLegal:    {"{", "}"},
			sqLight:    {"[", "]"},
			sqDark:     {"[", "]"},
		},
	},
}
