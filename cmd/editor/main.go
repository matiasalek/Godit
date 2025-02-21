package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type mode int

const (
	normal mode = iota
	insert
)

type model struct {
	viewport  viewport.Model
	textInput textinput.Model
	mode      mode
	ready     bool
}

func initialModel() model {
	vp := viewport.New(20, 10) // Initial placeholder size
	vp.SetContent("This is the viewport.\nUse 'i' to enter insert mode.\nPress 'Esc' to return to normal mode.")

	ti := textinput.New()
	ti.Placeholder = "Insert text here..."
	ti.Prompt = ":"
	ti.CharLimit = 256
	ti.Focus()

	return model{
		viewport:  vp,
		textInput: ti,
		mode:      normal,
		ready:     false,
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
		case "ctrl+c", "q":
			return m, tea.Quit

		case "i":
			if m.mode == normal {
				m.mode = insert
				m.textInput.Focus()
			}

		case "esc":
			if m.mode == insert {
				m.mode = normal
				m.textInput.Blur()
			}
		}

		if m.mode == insert {
			var cmd tea.Cmd
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		} else {
			m.viewport, _ = m.viewport.Update(msg)
		}

	case tea.WindowSizeMsg:
		inputHeight := 1 // Height of the input line
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - inputHeight
		m.ready = true
	}

	return m, nil
}

func (m model) View() string {
	if !m.ready {
		return "Loading..."
	}
	if m.mode == insert {
		return m.textInput.View() + "\n" + m.viewport.View()
	}
	return m.viewport.View()
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
