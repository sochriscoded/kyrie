package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sochriscoded/kyrie/internal/screens"
)

func main() {
	p := tea.NewProgram(screens.NewRootModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
