package internal

import (
	"github.com/charmbracelet/bubbles/textarea"
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

func InitialModel() model {
	ta := textarea.New()
	ta.Placeholder = "CTRL + C to quit. i to insert. esc to normal mode"
	ta.ShowLineNumbers = true

	return model{
		textArea: ta,
		mode:     normal,
		ready:    false,
	}
}
