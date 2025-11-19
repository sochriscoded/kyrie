package settings

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sochriscoded/kyrie/internal/screens/mainmenu"
	"github.com/sochriscoded/kyrie/internal/ui"
)

type SettingsScreen struct{}

func NewSettingsScreen() SettingsScreen {
	return SettingsScreen{}
}

func (m SettingsScreen) Init() tea.Cmd {
	return nil
}

func (m SettingsScreen) Update(msg tea.Msg) (ui.Screen, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "esc":
			// Go back to Main Menu
			return mainmenu.NewMainMenu(), nil

		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m SettingsScreen) View() string {
	return `
SETTINGS
--------

(pretend settings go here)

Press ESC → Back to Main Menu
Press q → Quit`
}
