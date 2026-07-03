package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/screen"
)

type model struct {
	currentView   string
	browserScreen *screen.BrowserScreen
}

func NewModel(a *app.App) model {
	browserScreen := screen.NewBrowserScreen(a)

	return model{
		currentView:   "browser",
		browserScreen: browserScreen,
	}
}

func (m model) Init() tea.Cmd {
	return m.browserScreen.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	if m.currentView == "browser" {
		cmd := m.browserScreen.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() tea.View {
	if m.currentView == "browser" {
		return m.browserScreen.View()
	}

	v := tea.NewView("View not found")

	v.AltScreen = true

	return v
}

func Run() error {
	a, err := app.Create()

	if err != nil {
		return err
	}

	m := NewModel(a)

	p := tea.NewProgram(m)

	_, err = p.Run()

	if err != nil {
		return err
	}

	return nil
}
