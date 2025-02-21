package internal

import tea "github.com/charmbracelet/bubbletea"

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		tea.ClearScreen,
	)
}
