package ui

// uses hexidecimal representation of colours e.g 0x3E434D
type pallette struct {
	name        string
	darkSquare  colour
	lightSquare colour
	whitePiece  colour
	blackPiece  colour
	cursor      colour
	selected    colour
	legalMove   colour
}

var themes = []pallette{
	{
		name:        "neon",
		lightSquare: mustHex("#ffc3fc"),
		darkSquare:  mustHex("#b701b4"),
		whitePiece:  mustHex("#ffc4df"),
		blackPiece:  mustHex("#45007e"),
		cursor:      mustHex("#fff200"),
		selected:    mustHex("#70faff"),
		legalMove:   mustHex("#d8ff00"),
	},
	{
		name:        "red-blue",
		lightSquare: mustHex("#cb4848"),
		darkSquare:  mustHex("#0a01b7"),
		whitePiece:  mustHex("#ffc4e7"),
		blackPiece:  mustHex("#006b7e"),
		cursor:      mustHex("#ffff00"),
		selected:    mustHex("#ff8c42"),
		legalMove:   mustHex("#00ff9c"),
	},
	{
		name:        "navy",
		lightSquare: mustHex("#7DE2E2"),
		darkSquare:  mustHex("#123C69"),
		whitePiece:  mustHex("#EFFFFF"),
		blackPiece:  mustHex("#25004F"),
		cursor:      mustHex("#FFD60A"),
		selected:    mustHex("#FF4DDE"),
		legalMove:   mustHex("#7CFF6B"),
	},
	{
		name:        "orange",
		lightSquare: mustHex("#FFB347"),
		darkSquare:  mustHex("#162A70"),
		whitePiece:  mustHex("#FFF3D6"),
		blackPiece:  mustHex("#35005C"),
		cursor:      mustHex("#FFFF00"),
		selected:    mustHex("#FF4F81"),
		legalMove:   mustHex("#54FF9F"),
	},
	{
		name:        "rgb",
		lightSquare: mustHex("#00E5FF"),
		darkSquare:  mustHex("#FF00AA"),
		whitePiece:  mustHex("#FFFF00"),
		blackPiece:  mustHex("#240046"),
		cursor:      mustHex("#FFFFFF"),
		selected:    mustHex("#00FF66"),
		legalMove:   mustHex("#FF6600"),
	},
}
