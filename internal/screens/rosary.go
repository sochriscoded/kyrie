package screens

import tea "github.com/charmbracelet/bubbletea"

type Rosary struct {
}

func NewRosary() Rosary {
	return Rosary{}
}

func (m Rosary) Init() tea.Cmd {
	return nil
}

func (m Rosary) Update(msg tea.Msg) (Screen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return NewMainMenu(), nil
		}
	}

	return m, nil
}

func (m Rosary) View() string {
	return "Pray the Rosary!"
}

//TODO: Figure Out How to Make Animations and Text Appear
//TODO: Figure Out How to pull JSON From File
//TODO: Figure Out Logic of Rosary
