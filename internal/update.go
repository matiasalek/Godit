package internal

import tea "github.com/charmbracelet/bubbletea"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		// Quit
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		// Enter insert mode
		if msg.String() == "i" && m.mode == normal {
			m.mode = insert
			m.textArea.Focus()
			return m, tea.Batch(tea.EnterAltScreen, tea.ClearScreen)
		}

		// Back to normal mode if insert mode is on
		if msg.String() == "esc" && m.mode == insert {
			m.mode = normal
			return m, nil
		}

		// Handle cursor movement only in normal mode
		if m.mode == normal {
			switch msg.String() {
			case "k":
				m.textArea.CursorUp()
			case "j":
				m.textArea.CursorDown()
			case "h":
				m.textArea.CursorStart()
			case "l":
				m.textArea.CursorEnd()
			}
		}

		// Allow typing only in insert mode
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
