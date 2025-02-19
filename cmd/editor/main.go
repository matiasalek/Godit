package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type EditorModel struct {
	content string
	preview string
	width   int
	height  int
}

func (m EditorModel) Init() tea.Cmd {
	return nil
}

func (m EditorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "backspace":
			if len(m.content) > 0 {
				m.content = m.content[:len(m.content)-1]
			}
		default:
			if msg.Type == tea.KeyRunes {
				m.content += string(msg.Runes)
			}
		}
		// For now, just mirror content to preview
		m.preview = m.content

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func (m EditorModel) View() string {
	// Basic split screen
	leftStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderRight(true).
		Width(m.width/2 - 1).
		Height(m.height - 2)

	rightStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderLeft(true).
		Width(m.width/2 - 1).
		Height(m.height - 2)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftStyle.Render(m.content),
		rightStyle.Render(m.preview),
	)
}

func main() {
	p := tea.NewProgram(
		EditorModel{},
		tea.WithAltScreen(),       // Use alternate screen buffer
		tea.WithMouseCellMotion(), // Enable mouse support
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}
