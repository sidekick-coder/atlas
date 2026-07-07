package root

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/components"
	"github.com/sidekick-coder/atlas/tui/components/input"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/models"
)

var Program *tea.Program

type model struct {
	app *app.App

	screens      []models.Screen
	availableScreens map[string]models.ScreenFactory
	screenHeight int
	currentIndex int

	width  int
	height int

	tabBar  *components.TabBar
	toolbar *components.Toolbar
	footer  *components.Footer

	input  *input.Input
	toaster *toast.Component
}

func New(a *app.App) model {
	screens := []models.Screen{}
	availableScreens := make(map[string]models.ScreenFactory)

	tabBar := components.NewTabBar()

	toolbar := components.NewToolbar()
	toolbar.SetTitle("󰉋 " + a.WorkspacePath())

	footer := components.NewFooter()

	input := input.New()
	toaster := toast.New()

	m := model{
		app:          a,
		currentIndex: 0,

		screens:     screens,
		tabBar:  tabBar,
		toolbar: toolbar,
		footer:  footer,

		input:   input,
		toaster: toaster,

		availableScreens: availableScreens,
	}

	return m
}

func (m model) Init() tea.Cmd {
	m.LoadScreenRegistry()
	m.LoadBindings()
	return nil
}
