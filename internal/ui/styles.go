package ui

import (
	"fmt"
	"strconv"
	"strings"
)

type SQType int

const (
	sqCursor SQType = iota
	sqLegal
	sqSelected
	sqLight
	sqDark
)

type colour uint32
type colourType int

const (
	fg colourType = 38
	bg colourType = 48
)

// returns the colour as an escape sequence
func (c colour) addColour(ctype colourType) string {
	r := c >> 16 & 0xFF
	g := c >> 8 & 0xFF
	b := c & 0xFF
	return fmt.Sprintf("\x1b[%d;2;%d;%d;%dm", ctype, r, g, b)
}

const resetColour = "\x1b[0m"

// return a painted string
func (c colour) paint(s string, ctype colourType) string {
	// \x1b [ paintType ; 2 ; R ; G ; B m
	return fmt.Sprintf("%s%s%s", c.addColour(ctype), s, resetColour)
}

// takes a hex colour code and return a colour in unit32 hex
func mustHex(s string) colour {
	s = strings.TrimPrefix(s, "#")
	clr, err := strconv.ParseUint(s, 16, 32)

	if err != nil {
		panic(fmt.Errorf("%s is invalid", s))
	}
	return colour(clr)
}
