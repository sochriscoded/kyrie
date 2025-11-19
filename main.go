package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sochriscoded/kyrie/internal/view"
)

func main() {
	p := tea.NewProgram(view.NewRootModel())
	if err := p.Run(); err != nil {
		panic(err)
	}
}
