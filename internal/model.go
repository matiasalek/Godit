package internal

import (
	"github.com/charmbracelet/bubbles/textarea"
)

type mode int

const (
	normal mode = iota
	insert
	command
)

type model struct {
	textArea      textarea.Model
	mode          mode
	commandInput  string
	statusMessage string
	ready         bool
}

func InitialModel() model {
	ta := textarea.New()
	ta.Placeholder = ":q to exit. i to insert. esc to normal mode"
	ta.ShowLineNumbers = true

	return model{
		textArea: ta,
		mode:     normal,
		ready:    false,
	}
}
