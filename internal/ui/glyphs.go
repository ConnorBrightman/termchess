package ui

type glyphSet struct {
	name         string
	sqDarkEmpty  string
	sqLightEmpty string
	squareChars  map[SQType]struct{ open, close string }
}

var glyphTypes = []glyphSet{
	{
		name:         "minimal",
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
	{
		name:         "contrast",
		sqDarkEmpty:  ".",
		sqLightEmpty: "#",
		squareChars: map[SQType]struct{ open, close string }{
			sqCursor:   {"<", ">"},
			sqSelected: {"(", ")"},
			sqLegal:    {"{", "}"},
			sqLight:    {"[", "]"},
			sqDark:     {"[", "]"},
		},
	},
	{
		name:         "full",
		sqDarkEmpty:  "#",
		sqLightEmpty: "#",
		squareChars: map[SQType]struct{ open, close string }{
			sqCursor:   {"<", ">"},
			sqSelected: {"(", ")"},
			sqLegal:    {"{", "}"},
			sqLight:    {"[", "]"},
			sqDark:     {"[", "]"},
		},
	},
}
