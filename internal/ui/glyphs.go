package ui

type glyphSet struct {
	name         string
	sqDarkEmpty  string
	sqLightEmpty string
	squareChars  map[SQType]struct{ open, close string }
}

var glyphTypes = []glyphSet{
	{
		name:         "basic",
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
