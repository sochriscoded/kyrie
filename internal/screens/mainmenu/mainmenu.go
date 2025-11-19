package mainmenu

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sochriscoded/kyrie/internal/screens/settings"
	"github.com/sochriscoded/kyrie/internal/ui"
)

type MainMenu struct{}

func NewMainMenu() MainMenu {
	return MainMenu{}
}

func (m MainMenu) Init() tea.Cmd {
	return nil
}

func (m MainMenu) Update(msg tea.Msg) (ui.Screen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "enter":
			// Switch to Settings screen
			return settings.NewSettingsScreen(), nil

		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MainMenu) View() string {
	return `
MAIN MENU
---------

Press ENTER → Settings
Press q → Quit`
}
