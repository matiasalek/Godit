package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type mode int

const (
	normal mode = iota
	insert
)

type model struct {
	textArea textarea.Model
	mode     mode
	ready    bool
}

func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = "CTRL + C to quit. i to insert. esc to normal mode"
	ta.ShowLineNumbers = true

	return model{
		textArea: ta,
		mode:     normal,
		ready:    false,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		tea.ClearScreen,
	)
}

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

func (m model) View() string {
	textView := m.textArea.View()

	status := ""
	if m.mode == insert {
		status = " -- INSERT -- "
	} else {
		status = " -- NORMAL -- "
	}
	return textView + "\n" + status
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
