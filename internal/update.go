package internal

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"

	"os"
	"strings"
)

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

		if msg.String() == ":" && m.mode == normal {
			m.mode = command
			m.commandInput = ""
			return m, nil
		}

		// Command mode input handling
		if m.mode == command {
			switch msg.String() {
			case "enter":
				trimmedInput := strings.TrimSpace(m.commandInput)

				// Handle quit with ":q"
				if trimmedInput == "q" {
					return m, tea.Quit
				}

				if strings.HasPrefix(m.commandInput, "w ") {
					// Save file logic
					filename := strings.TrimSpace(strings.TrimPrefix(m.commandInput, "w "))
					err := os.WriteFile(filename, []byte(m.textArea.Value()), 0644)
					if err != nil {
						m.statusMessage = "Error saving file: " + err.Error()
					} else {
						m.statusMessage = "File saved: " + filename
					}
					m.mode = normal
					m.commandInput = ""
					// Clear the status message after 2 seconds
					return m, tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
						return clearStatusMsg{}
					})
				}
				m.mode = normal
				m.commandInput = ""
				return m, nil

			case "esc":
				m.mode = normal
				m.commandInput = ""
				return m, nil

			case "backspace":
				if len(m.commandInput) > 0 {
					m.commandInput = m.commandInput[:len(m.commandInput)-1]
				}
				return m, nil

			default:
				m.commandInput += msg.String()
				return m, nil
			}
		}

		// Handle cursor movement only in normal mode
		if m.mode == normal {
			switch msg.String() {
			case "k":
				m.textArea.CursorUp()
			case "j":
				m.textArea.CursorDown()
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

	case clearStatusMsg:
		m.statusMessage = ""
		return m, nil

	}
	return m, nil
}

type clearStatusMsg struct{}
