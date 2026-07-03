package tui

import (
	tea "charm.land/bubbletea/v2"
)

type model struct {}

func (m model) Init() tea.Cmd {
	// Initialize the model
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	v := tea.NewView("Hello, TUI!\nPress 'q' or 'Ctrl+C' to quit.")

	v.AltScreen = true

	return v
}

func Run() error {
	p := tea.NewProgram(model{})

	_, err := p.Run()

	if err != nil {
		return err
	}

	return nil
}
