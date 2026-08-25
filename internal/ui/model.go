// Package ui
package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/ConnorBrightman/termchess/internal/chess"
)

type model struct {
	board        chess.Board
	cursor       chess.Square
	selected     chess.Square
	hasSelection bool
}

func initialModel() model {
	return model{
		board:        chess.StartPosition(),
		hasSelection: false,
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
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
			sq := m.cursor
			pt := m.board.PieceAt(sq).PieceType()
			if !m.hasSelection {
				if pt == chess.Empty {
					break
				} else {
					m.selected = sq
					m.hasSelection = true
				}
			} else if sq == m.selected {
				m.hasSelection = false
			} else {
				mv := chess.Move{From: m.selected, To: sq}
				m.board = m.board.MakeMove(mv)
				m.hasSelection = false
			}
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

	for r := chess.Rank(7); r >= 0; r-- {
		for f := chess.File(0); f < 8; f++ {
			sq := chess.Square{Rank: r, File: f}
			if sq == m.selected && m.hasSelection {
				s += fmt.Sprintf("(%v)", m.board.PieceAt(sq))
			} else if sq == m.cursor {
				s += fmt.Sprintf("[%v]", m.board.PieceAt(sq))
			} else {
				s += fmt.Sprintf(" %v ", m.board.PieceAt(sq))
			}
		}
		s += "\n"
	}

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
