package root

import (
	"errors"
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/components"
	"github.com/sidekick-coder/atlas/tui/components/input"
	"github.com/sidekick-coder/atlas/tui/screen"
	"github.com/sidekick-coder/atlas/tui/screen/entry"
	"github.com/sidekick-coder/atlas/tui/screen/entrysingle"
)

type model struct {
	app          *app.App
	screens      []screen.Screen
	currentIndex int
	width        int
	height       int
	input        *input.Input

	tabBar  *components.TabBar
	toolbar *components.Toolbar
	footer  *components.Footer
}

func New(a *app.App) model {
	screens := []screen.Screen{}

	entryScreen := entry.Create(a)
	screens = append(screens, entryScreen)

	tabBar := components.NewTabBar()
	tabBar.Add("[0]: " + entryScreen.Title())

	toolbar := components.NewToolbar()
	toolbar.SetTitle("󰉋 " + a.WorkspacePath())

	footer := components.NewFooter()

	input := input.New()

	m := model{
		app:          a,
		currentIndex: 0,
		screens:      screens,
		tabBar:       tabBar,
		toolbar:      toolbar,
		footer:       footer,
		input:        input,
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

func (m *model) AddScreen(name string, options map[string]any) error {
	if name == "entry" {
		s := entry.Create(m.app)
		m.screens = append(m.screens, s)
		index := len(m.screens) - 1
		m.tabBar.Add(fmt.Sprintf("[%d]: %s", index, s.Title()))

		m.SetCurrentScreen(index)
		return nil
	}

	if name == "entry-single" {
		path, ok := options["path"].(string)

		if !ok {
			return errors.New("Path not provided")
		}

		s := entrysingle.Create(m.app, path)
		m.screens = append(m.screens, s)
		index := len(m.screens) - 1
		m.tabBar.Add(fmt.Sprintf("[%d]: %s", index, s.Title()))

		m.SetCurrentScreen(index)
		return nil
	}

	return errors.New(fmt.Sprintf("Unknown screen: %s", name))
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

	for _, s := range m.screens {
		s.SetSize(width, contentHeight)
	}
}

func (m model) Init() tea.Cmd {
	s := m.screens[m.currentIndex]

	cmd := s.Init()

	if cmd != nil {
		return cmd
	}

	return nil
}

