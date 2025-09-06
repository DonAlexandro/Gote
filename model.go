package main

import (
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	listView uint = iota
	titleView
	bodyView
)

type model struct {
	state       uint
	store       *Store
	notes       []Note
	currentNote Note
	listIndex   int
	textarea    textarea.Model
	textinput   textinput.Model
	// ...
}

func NewModel(s *Store) model {
	notes, err := s.GetAllNotes()

	if err != nil {
		log.Fatalf("failed to get notes: %v", err)
	}

	return model{
		state:     listView,
		store:     s,
		notes:     notes,
		textarea:  textarea.New(),
		textinput: textinput.New(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	m.textinput, cmd = m.textinput.Update(msg)
	cmds = append(cmds, cmd)
	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()

		switch m.state {
		case listView:
			switch key {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "n":
				m.textinput.SetValue("")
				m.textinput.Focus()
				m.currentNote = Note{}
				m.state = titleView
			case "up", "k":
				if m.listIndex > 0 {
					m.listIndex--
				}
			case "down", "j":
				if m.listIndex < len(m.notes)-1 {
					m.listIndex++
				}
			case "enter":
				if len(m.notes) > 0 {
					m.currentNote = m.notes[m.listIndex]

					m.textarea.SetValue(m.currentNote.Body)
					m.textarea.Focus()
					m.textarea.CursorEnd()

					m.state = bodyView
				}
			}

		case titleView:
			switch key {
			case "enter":
				title := m.textinput.Value()

				if title != "" {
					m.currentNote.Title = title

					m.textarea.SetValue(m.currentNote.Body)
					m.textarea.Focus()
					m.textarea.CursorEnd()

					m.state = bodyView
				}
			case "esc":
				m.state = listView
			}

		case bodyView:
			switch key {
			case "ctrl+s":
				body := m.textarea.Value()
				m.currentNote.Body = body

				var err error

				if err = m.store.AddNote(m.currentNote); err != nil {
					// TODO: handle error
					return m, tea.Quit
				}

				m.notes, err = m.store.GetAllNotes()

				if err != nil {
					// TODO: handle error
					return m, tea.Quit
				}

				m.currentNote = Note{}
				m.state = listView
			case "esc":
				m.state = listView
			}
		}
	}

	return m, tea.Batch(cmds...)
}
