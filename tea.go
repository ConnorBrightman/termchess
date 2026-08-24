package main

// These imports will be used later in the tutorial. If you save the file
// now, Go might complain they are unused, but that's fine.
// You may also need to run `go mod tidy` to download bubbletea and its
// dependencies.
import (
	"fmt"
	"os"

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
		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursorR < 7 {
				m.cursorR++
			} else {
				m.cursorR = 0
			}
		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursorR > 0 {
				m.cursorR--
			} else {
				m.cursorR = 7
			}
		// The "down" and "j" keys move the cursor down
		case "left", "h":
			if m.cursorF > 0 {
				m.cursorF--
			} else {
				m.cursorF = 7
			}
		// The "down" and "j" keys move the cursor down
		case "right", "l":
			if m.cursorF < 7 {
				m.cursorF++
			} else {
				m.cursorF = 0
			}
		}
		// Return the updated model to the Bubble Tea runtime for processing.
		// Note that we're not returning a command.
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

func chessGame() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
