// Package ui
package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/ConnorBrightman/termchess/internal/chess"
)

type model struct {
	board            chess.Board
	cursorR, cursorF int
}

func initialModel() model {
	return model{
		board: chess.StartPosition(),
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
			if m.cursorR < 7 {
				m.cursorR++
			} else {
				m.cursorR = 0
			}
		case "down", "j":
			if m.cursorR > 0 {
				m.cursorR--
			} else {
				m.cursorR = 7
			}
		case "left", "h":
			if m.cursorF > 0 {
				m.cursorF--
			} else {
				m.cursorF = 7
			}
		case "right", "l":
			if m.cursorF < 7 {
				m.cursorF++
			} else {
				m.cursorF = 0
			}
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	// The header
	s := "Welcome to TermChess\n"

	for r := 7; r >= 0; r-- {
		for f := 0; f < 8; f++ {
			if r == m.cursorR && f == m.cursorF {
				s += fmt.Sprintf("[%v]", m.board.Squares[r][f])
			} else {
				s += fmt.Sprintf(" %v ", m.board.Squares[r][f])
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
