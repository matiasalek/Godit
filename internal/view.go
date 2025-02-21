package internal

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
