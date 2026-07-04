package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/components"
	"github.com/sidekick-coder/atlas/tui/screen"
)

type model struct {
	currentView   string
	browserScreen *screen.BrowserScreen
	width         int
	height        int
}

func NewModel(a *app.App) model {
	return model{
		currentView:   "browser",
		browserScreen: screen.NewBrowserScreen(a),
	}
}

func (m model) Init() tea.Cmd {
	return m.browserScreen.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if ws, ok := msg.(tea.WindowSizeMsg); ok {
		m.width = ws.Width
		m.height = ws.Height
		components.GlobalInput.SetSize(ws.Width, ws.Height)
	}

	// Route all input to GlobalInput while it's active; block everything else.
	if components.GlobalInput.Active() {
		cmd := components.GlobalInput.Update(msg)
		return m, cmd
	}

	if m.currentView == "browser" {
		cmd := m.browserScreen.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() tea.View {
	var content string
	if m.currentView == "browser" {
		content = m.browserScreen.Render()
	}

	if components.GlobalInput.Active() {
		content = components.PlaceOverlay(components.GlobalInput.Box(), content, m.width, m.height)
	}

	v := tea.NewView(content)
	v.AltScreen = true
	return v
}

func Run() error {
	a, err := app.Create()
	if err != nil {
		return err
	}

	p := tea.NewProgram(NewModel(a))
	_, err = p.Run()
	return err
}
