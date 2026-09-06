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
}

func initialModel() model {
	return model{
		board:        chess.StartPosition(),
		hasSelection: false,
		turn:         chess.White,
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
		case "up", "k":
			if m.cursor.Rank < 7 {
				m.cursor.Rank++
			} else {
				m.cursor.Rank = 0
			}
		case "down", "j":
			if m.cursor.Rank > 0 {
				m.cursor.Rank--
			} else {
				m.cursor.Rank = 7
			}
		case "left", "h":
			if m.cursor.File > 0 {
				m.cursor.File--
			} else {
				m.cursor.File = 7
			}
		case "right", "l":
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
	colNotation := "  " + columnNotation
	s += colNotation
	s += "\n"
	// render board
	for r := chess.Rank(7); r >= 0; r-- {
		s += fmt.Sprintf("%v ", r)

		for f := chess.File(0); f < 8; f++ {
			sq := chess.Square{Rank: r, File: f}
			dark := (int(sq.Rank)+int(sq.File))%2 == 0
			p := m.board.PieceAt(sq)
			empty := (p.PieceType() == chess.Empty)

			ps := p.String()

			if empty {
				if dark {
					ps = sqDarkEmpty
				} else {
					ps = sqLightEmpty
				}
			}
			// render squares
			// selected piece square
			t := sqLight
			switch {
			case m.hasSelection && sq == m.selected:
				t = sqSelected
			case sq == m.cursor:
				t = sqCursor
			case legal[sq]:
				t = sqLegal
			case dark:
				t = sqDark
			}
			ch := squareChars[t]

			s += ch.open + ps + ch.close
		}
		s += fmt.Sprintf(" %v", r)
		s += "\n"
	}
	s += colNotation
	return s
}

var columnNotation string = " a  b  c  d  e  f  g  h "

type SQType int

const (
	sqCursor SQType = iota
	sqLegal
	sqSelected
	sqLight
	sqDark
)

var (
	sqDarkEmpty  string = "."
	sqLightEmpty string = "#"
	squareChars         = map[SQType]struct{ open, close string }{
		sqCursor:   {"<", ">"},
		sqSelected: {"(", ")"},
		sqLegal:    {"{", "}"},
		sqLight:    {"[", "]"},
		sqDark:     {" ", " "},
	}
)

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
