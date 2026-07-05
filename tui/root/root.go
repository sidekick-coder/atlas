package root

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/components"
	"github.com/sidekick-coder/atlas/tui/components/input"
	"github.com/sidekick-coder/atlas/tui/models"
)

type model struct {
	app *app.App

	screens      []models.Screen
	availableScreens map[string]models.ScreenFactory
	screenHeight int
	currentIndex int

	width  int
	height int
	input  *input.Input

	tabBar  *components.TabBar
	toolbar *components.Toolbar
	footer  *components.Footer
}

func New(a *app.App) model {
	screens := []models.Screen{}
	availableScreens := make(map[string]models.ScreenFactory)

	tabBar := components.NewTabBar()

	toolbar := components.NewToolbar()
	toolbar.SetTitle("󰉋 " + a.WorkspacePath())

	footer := components.NewFooter()

	input := input.New()

	m := model{
		app:          a,
		currentIndex: 0,

		screens:     screens,
		tabBar:  tabBar,
		toolbar: toolbar,
		footer:  footer,
		input:   input,
		availableScreens: availableScreens,
	}

	m.SetCurrentScreen(0)

	return m
}

func (m *model) SetSize(width int, height int) {
	m.width = width
	m.height = height

	components.GlobalInput.SetSize(width, height)
	components.GlobalToast.SetSize(width, height)

	m.tabBar.SetWidth(width)
	m.toolbar.SetWidth(width)
	m.footer.SetWidth(width)
	m.input.SetScreenSize(width, height)

	toolbarHeight := 1
	tabBarHeight := 1
	footerHeight := 1
	contentHeight := height - toolbarHeight - tabBarHeight - footerHeight
	m.screenHeight = contentHeight

	for _, s := range m.screens {
		s.SetSize(width, contentHeight)
	}
}

func (m model) Init() tea.Cmd {
	m.LoadScreenRegistry()
	m.LoadBindings()
	return nil
}
