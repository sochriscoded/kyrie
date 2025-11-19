package screens

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainMenu struct {
	width    int
	height   int
	cursor   int
	choices  []string
	selected string
}

func NewMainMenu() MainMenu {
	return MainMenu{choices: []string{"Rosary", "Settings", "Quit"}}
}

func (m MainMenu) Init() tea.Cmd {
	return nil
}

func (m MainMenu) Update(msg tea.Msg) (Screen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "enter":
			m.selected = m.choices[m.cursor]
			if m.selected == "Quit" {
				return m, tea.Quit
			}
			if m.selected == "Settings" {
				return NewSettingsScreen(), nil
			}
			if m.selected == "Rosary" {
				return NewRosary(), nil
			}

		case "q", "ctrl+c":
			return m, tea.Quit

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")) // pink-purple

	cursorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("212")) // light blue

	menuBox = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 3).
		BorderForeground(lipgloss.Color("240"))

	faintText = lipgloss.NewStyle().Faint(true)
)

func (m MainMenu) View() string {
	s := titleStyle.Render("🍵 Bubble Tea App") + "\n"
	s += faintText.Render("Use ↑/↓ to move, Enter to select, q to quit") + "\n\n"

	for i, choice := range m.choices {
		cursor := " "
		line := choice
		if m.cursor == i {
			cursor = cursorStyle.Render(">")
			line = cursorStyle.Render(choice)
		}
		s += fmt.Sprintf("%s %s\n", cursor, line)
	}

	// Put menu inside a box
	menu := menuBox.Render(s)

	// Center it on screen
	if m.width == 0 || m.height == 0 {
		// Before first resize event, just return uncentered
		return menu
	}

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		menu,
		lipgloss.WithWhitespaceChars(" "),
		lipgloss.WithWhitespaceForeground(lipgloss.Color("0")),
	)
}
