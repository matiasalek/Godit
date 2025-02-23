package internal

func (m model) View() string {
	textView := m.textArea.View()

	status := ""
	switch m.mode {
	case insert:
		status = " -- INSERT -- "
	case normal:
		status = " -- NORMAL -- "
	case command:
		status = ":" + m.commandInput
	}

	if m.statusMessage != "" {
		status += " | " + m.statusMessage
	}

	return textView + "\n" + status
}
