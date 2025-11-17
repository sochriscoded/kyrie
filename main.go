package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	width    int
	height   int
	cursor   int
	choices  []string
	selected string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = m.choices[m.cursor]
			if m.selected == "Quit" {
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

// --- Styling ---

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205"))

	cursorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("212")) // light blue

	menuBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1, 3).
			BorderForeground(lipgloss.Color("240"))

	faintText = lipgloss.NewStyle().Faint(true)
)

func (m model) View() string {
	// Build menu body
	s := titleStyle.Render("Kyrie") + "\n"
	s += faintText.Render("The Liturgy of the Hours (Daily Office), Psalms, and more in your Terminal.") + "\n\n"
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

func main() {
	m := model{
		choices: []string{"Daily Office", "Rosary", "Psalms", "Settings", "Quit"},
	}

	finalModel, err := tea.NewProgram(m, tea.WithAltScreen()).Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	if fm, ok := finalModel.(model); ok && fm.selected != "" {
		fmt.Printf("\nYou selected: %s\n", fm.selected)
	}
}
