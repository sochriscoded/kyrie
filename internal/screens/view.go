package screens

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Screen is a small interface each screen implements.
// This lets the root model easily switch between screens.

type RootModel struct {
	active Screen
}

func NewRootModel() RootModel {
	return RootModel{
		active: NewMainMenu(),
	}
}

func (m RootModel) Init() tea.Cmd {
	return m.active.Init()
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	screen, cmd := m.active.Update(msg)
	m.active = screen
	return m, cmd
}

func (m RootModel) View() string {
	return m.active.View()
}
