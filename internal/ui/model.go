// Package ui
package ui

import (
	"fmt"
	"slices"

	tea "charm.land/bubbletea/v2"
	"github.com/ConnorBrightman/termchess/internal/chess"
)

type model struct {
	board        chess.Board
	cursor       chess.Square
	selected     chess.Square
	hasSelection bool
	turn         chess.Colour
	message      string
	themeNo      int
	glyphsNo     int
	piecesNo     int
	theme        pallette
	glyphs       glyphSet
	pieces       pieceStyle
}

func initialModel() model {
	return model{
		board:        chess.StartPosition(),
		hasSelection: false,
		turn:         chess.White,
		themeNo:      0,
		glyphsNo:     0,
		piecesNo:     0,
		theme:        themes[0],
		glyphs:       glyphTypes[0],
		pieces:       pieceStyles[0],
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		m.message = ""
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "t":
			m.nextTheme()
		case "p":
			m.nextPieces()
		case "g":
			m.nextGlyphs()
		case "up", "k", "w":
			if m.cursor.Rank < 7 {
				m.cursor.Rank++
			} else {
				m.cursor.Rank = 0
			}
		case "down", "j", "s":
			if m.cursor.Rank > 0 {
				m.cursor.Rank--
			} else {
				m.cursor.Rank = 7
			}
		case "left", "h", "a":
			if m.cursor.File > 0 {
				m.cursor.File--
			} else {
				m.cursor.File = 7
			}
		case "right", "l", "d":
			if m.cursor.File < 7 {
				m.cursor.File++
			} else {
				m.cursor.File = 0
			}
		case "space", "enter":
			m.handleSelection()
		case "backspace":
			if m.hasSelection {
				m.hasSelection = false
			}
		}
	}
	return m, nil
}

func (m model) View() tea.View {

	// The header
	s := "Welcome to TermChess\n"

	s += fmt.Sprintf("Theme [%s]\n", m.theme.name)
	s += fmt.Sprintf("Pieces [%s]\n", m.pieces.name)
	s += fmt.Sprintf("Glyphs [%s]\n", m.glyphs.name)

	switch {
	case m.board.IsCheckmate(m.turn):
		s += fmt.Sprintf("Checkmate — %v wins\n", m.turn.Opponent())
	case m.board.IsStalemate(m.turn):
		s += "Stalemate — draw\n"
	case m.board.IsCheck(m.turn):
		s += fmt.Sprintf("%v to move — CHECK\n", m.turn)
	default:
		s += fmt.Sprintf("%v to move\n", m.turn)
	}

	s += renderBoard(m)

	s += fmt.Sprintf("\n%s", m.message)

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return tea.NewView(s)
}

func Run() error {
	p := tea.NewProgram(initialModel())
	_, err := p.Run()

	return err
}

func renderBoard(m model) string {
	s := ""
	legal := map[chess.Square]bool{}
	if m.hasSelection {
		for _, dest := range m.board.Moves(m.selected) {
			legal[dest] = true
		}
	}
	// top column notation
	colNotation := "  " + columnNotation
	s += colNotation
	s += "\n"

	// render board
	for r := chess.Rank(7); r >= 0; r-- {
		// add row notation tp start
		s += fmt.Sprintf("%v ", r)

		for f := chess.File(0); f < 8; f++ {

			sq := chess.Square{Rank: r, File: f}
			dark := (int(sq.Rank)+int(sq.File))%2 == 0
			p := m.board.PieceAt(sq)
			empty := (p.PieceType() == chess.Empty)

			pc := p.PieceColour()
			ps := m.setPieceLooks(p)

			var glyph string
			var col colour

			if empty {
				if dark {
					glyph = m.glyphs.sqDarkEmpty
					col = m.theme.darkSquare
				} else {
					glyph = m.glyphs.sqLightEmpty
					col = m.theme.lightSquare
				}
			} else {
				glyph = ps
				if pc == chess.Black {
					col = m.theme.blackPiece
				} else {
					col = m.theme.whitePiece
				}
			}
			ps = col.paint(glyph, fg)

			// render squares
			// selected piece square
			var sqGlyph SQType
			var sqColour colour
			var paintType colourType = fg
			switch {
			case sq == m.cursor:
				sqGlyph = sqCursor
				sqColour = m.theme.cursor
			case m.hasSelection && sq == m.selected:
				sqGlyph = sqSelected
				sqColour = m.theme.selected
			case legal[sq]:
				sqGlyph = sqLegal
				sqColour = m.theme.legalMove
			case dark:
				sqGlyph = sqDark
				sqColour = m.theme.darkSquare
			default:
				sqGlyph = sqLight
				sqColour = m.theme.lightSquare
			}
			ch := m.glyphs.squareChars[sqGlyph]
			s += sqColour.paint(ch.open, paintType) + ps + sqColour.paint(ch.close, paintType)
		}
		// add row notation to end
		s += fmt.Sprintf(" %v", r)
		s += "\n"
	}
	s += colNotation
	return s
}

var columnNotation string = " a  b  c  d  e  f  g  h "

func (m *model) handleSelection() {
	sq := m.cursor
	p := m.board.PieceAt(sq)
	pt := p.PieceType()
	pc := p.PieceColour()

	// if there isn't a selected square already
	if !m.hasSelection {
		// not your piece
		if pt == chess.Empty || pc != m.turn {
			m.message = "Not your piece"
			return
		}
		// is your piece
		m.selected = sq
		m.hasSelection = true
		return
	}
	// if the square is already selected
	if sq == m.selected {
		m.hasSelection = false
		return
	}
	// if you have selected a valid piece move it to the new square
	mv := chess.Move{From: m.selected, To: sq}
	// if the cusror is on a valid move square
	if slices.Contains(m.board.Moves(m.selected), m.cursor) {
		m.board = m.board.MakeMove(mv)
		m.hasSelection = false
		m.turn = m.turn.Opponent()
	} else {
		m.message = "Not a valid move"
	}
}

func (m *model) setTheme(i int) {
	m.themeNo = i
	m.theme = themes[m.themeNo]
}
func (m *model) nextTheme() {
	i := m.themeNo + 1
	if i < len(themes) {
		m.setTheme(i)
	} else {
		m.setTheme(0)
	}
}

func (m *model) setGlyphs(i int) {
	m.glyphsNo = i
	m.glyphs = glyphTypes[m.glyphsNo]
}
func (m *model) nextGlyphs() {
	i := m.glyphsNo + 1
	if i < len(glyphTypes) {
		m.setGlyphs(i)
	} else {
		m.setGlyphs(0)
	}
}

func (m *model) setPieces(i int) {
	m.piecesNo = i
	m.pieces = pieceStyles[m.piecesNo]
}
func (m *model) nextPieces() {
	i := m.piecesNo + 1
	if i < len(pieceStyles) {
		m.setPieces(i)
	} else {
		m.setPieces(0)
	}
}

func (m model) setPieceLooks(p chess.Piece) string {
	ps := p.String()
	switch ps {
	case "P":
		ps = m.pieces.wPawn
	case "N":
		ps = m.pieces.wKnight
	case "B":
		ps = m.pieces.wBishop
	case "R":
		ps = m.pieces.wRook
	case "Q":
		ps = m.pieces.wQueen
	case "K":
		ps = m.pieces.wKing
	case "p":
		ps = m.pieces.bPawn
	case "n":
		ps = m.pieces.bKnight
	case "b":
		ps = m.pieces.bBishop
	case "r":
		ps = m.pieces.bRook
	case "q":
		ps = m.pieces.bQueen
	case "k":
		ps = m.pieces.bKing
	}
	return ps
}
