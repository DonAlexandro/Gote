package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	appNameStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FAFAFA")).Background(lipgloss.Color("#7D56F4")).Padding(0, 1)
	faint           = lipgloss.NewStyle().Faint(true).Foreground(lipgloss.Color("255"))
	enumaratorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4")).MarginRight(1)
)

func (m model) View() string {
	s := appNameStyle.Render("NOTES APP") + "\n\n"

	if m.state == titleView {
		s += "Note title:\n\n"
		s += m.textinput.View() + "\n\n"
		s += faint.Render("Enter to save, esc to cancel")
	}

	if m.state == bodyView {
		s += "Note:\n\n"
		s += m.textarea.View() + "\n\n"
		s += faint.Render("ctrl+s to save, esc to cancel")
	}

	if m.state == listView {
		for i, n := range m.notes {
			prefix := " "

			if i == m.listIndex {
				prefix = ">"
			}

			shortBody := strings.ReplaceAll(n.Body, "\n", " ")

			if len(shortBody) > 40 {
				shortBody = shortBody[:40] + "..."
			}

			s += enumaratorStyle.Render(prefix) + n.Title + " | " + faint.Render(shortBody) + "\n"
		}

		s += faint.Render("n - new note, q - quit")
	}

	return s
}
