package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	s := &Store{}

	if err := s.Init(); err != nil {
		log.Fatalf("failed to initialize store: %v", err)
	}

	m := NewModel(s)

	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		log.Fatalf("unable to run TUI: %v", err)
	}
}
