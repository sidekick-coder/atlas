package root

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/components"
	"github.com/sidekick-coder/atlas/tui/components/input"
	"github.com/sidekick-coder/atlas/tui/models"
	"github.com/sidekick-coder/atlas/tui/screen/empty"
	"github.com/sidekick-coder/atlas/tui/screen/entry"
)

type model struct {
	app *app.App

	emptyScreen models.Screen

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

	es := empty.Create(models.ScreenPayload{
		App: a,
	})

	tabBar := components.NewTabBar()

	toolbar := components.NewToolbar()
	toolbar.SetTitle("󰉋 " + a.WorkspacePath())

	footer := components.NewFooter()

	input := input.New()

	m := model{
		app:          a,
		currentIndex: 0,

		screens:     screens,
		emptyScreen: es,

		tabBar:  tabBar,
		toolbar: toolbar,
		footer:  footer,
		input:   input,
		availableScreens: availableScreens,
	}

	m.SetCurrentScreen(0)
	m.SetBindings()

	return m
}

func (m *model) SetBindings() {
	bindings := m.GetBindings()

	m.footer.SetBindings(bindings...)
}

func (m *model) SetCurrentScreen(index int) {
	if index < 0 || index >= len(m.screens) {
		return
	}

	m.currentIndex = index
	m.tabBar.SetCurrent(index)
	m.footer.SetBindings(m.GetBindings()...)

	m.SetSize(m.width, m.height)

	s := m.screens[index]
	s.Init()
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
	m.emptyScreen.SetSize(width, contentHeight)

	for _, s := range m.screens {
		s.SetSize(width, contentHeight)
	}
}

func (m model) Init() tea.Cmd {
	m.availableScreens["entries"] = entry.Create

	return nil
}
