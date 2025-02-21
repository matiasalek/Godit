package internal

import tea "github.com/charmbracelet/bubbletea"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "i":
			if m.mode == normal {
				m.mode = insert
				m.textArea.Focus()
				m.textArea.Reset() // Clear textarea to avoid showing 'i'

				return m, tea.Batch(tea.EnterAltScreen, tea.ClearScreen)
			}

		case "esc":
			if m.mode == insert {
				m.mode = normal
				m.textArea.Blur()
				return m, nil
			}
		}

		if m.mode == insert {
			var cmd tea.Cmd
			m.textArea, cmd = m.textArea.Update(msg)
			return m, cmd
		}

	case tea.WindowSizeMsg:
		m.textArea.SetWidth(msg.Width)
		m.textArea.SetHeight(msg.Height - 1)
		m.ready = true
	}

	return m, nil
}
