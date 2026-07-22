package model

import (
	"fmt"
	"maps"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/tui/app/footer"
	"github.com/sidekick-coder/atlas/tui/app/screen"
	"github.com/sidekick-coder/atlas/tui/app/tabbar"
	"github.com/sidekick-coder/atlas/tui/app/toaster"
	"github.com/sidekick-coder/atlas/tui/app/toolbar"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/action"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/keymaps"
	"github.com/sidekick-coder/atlas/tui/models"

	"github.com/sidekick-coder/atlas/tui/screen/custom"
	"github.com/sidekick-coder/atlas/tui/screen/empty"
	"github.com/sidekick-coder/atlas/tui/screen/entrysingle"
	"github.com/sidekick-coder/atlas/tui/screen/entrytable"
)

var Program *tea.Program

type model struct {
	app    *app.App
	width  int
	height int
	ready  bool

	screen  *screen.Feature
	tabbar  *tabbar.Component
	footer  *footer.Component
	toolbar *toolbar.Component
	toaster *toaster.Component
}

func Create(a *app.App) model {
	m := model{
		app:    a,
		width:  100,
		height: 100,
		ready:  false,

		toolbar: toolbar.Create(a),
		footer:  footer.Create(),
		toaster: toaster.Create(),

		screen: screen.Create(),
		tabbar: tabbar.Create(),
	}

	action.SetManager(a.Action)
	keymaps.LoadConfigKeymaps(a.Config())

	return m
}

func (m *model) AddScreenEmpty() tea.Cmd {
	entries := []empty.Entry{}

	entries = append(entries, empty.Entry{
		ID:      "entry_list",
		Options: map[string]any{},
	})

	entries = append(entries, empty.Entry{
		ID:      "entry_table",
		Options: map[string]any{},
	})

	us, err := m.app.Config().GetScreens()

	if err != nil {
		return toast.Error(err.Error())
	}

	for _, s := range us {
		entries = append(entries, empty.Entry{
			ID:      s.ID,
			Options: s.Options,
		})
	}

	options := map[string]any{
		"entries": entries,
	}

	_, err = m.screen.Add("empty", options)

	if err != nil {
		return toast.Error(err.Error())
	}

	return nil
}

func (m *model) LoadUserScreen(screen config.Screen) (models.ScreenFactory, error) {
	original, ok := m.screen.GetDefinition(screen.Type)

	if !ok {
		return nil, fmt.Errorf("invalid screen type: %s", screen.Type)
	}

	fac := func(p models.ScreenPayload) (models.Screen, error) {
		maps.Copy(p.Options, screen.Options)

		return original(p)
	}

	return fac, nil
}

func (m model) EmptyScreenFactory(p models.ScreenPayload) (models.Screen, error) {
	entries := []empty.Entry{}

	entries = append(entries, empty.Entry{
		ID:      "entry_list",
		Options: map[string]any{},
	})

	entries = append(entries, empty.Entry{
		ID:      "entry_table",
		Options: map[string]any{},
	})

	us, err := m.app.Config().GetScreens()

	if err != nil {
		return nil, fmt.Errorf("error getting user screens: %w", err)
	}

	for _, s := range us {
		entries = append(entries, empty.Entry{
			ID:      s.ID,
			Options: s.Options,
		})
	}

	p.Options = map[string]any{
		"entries": entries,
	}

	return empty.Create(p)
}

func (m model) InitScreen() tea.Cmd {
	m.screen.SetApp(m.app)

	m.screen.SetDefinition("empty", m.EmptyScreenFactory)
	m.screen.SetDefinition("entry_table", entrytable.Create)
	m.screen.SetDefinition("entry_single", entrysingle.Create)
	m.screen.SetDefinition("custom", custom.Create)

	us, err := m.app.Config().GetScreens()

	if err != nil {
		return toast.Error(err.Error())
	}

	for _, s := range us {
		fac, err := m.LoadUserScreen(s)

		if err != nil {
			return toast.Error(err.Error())
		}

		m.screen.SetDefinition(s.ID, fac)
	}

	return nil
}

func (m model) InitTabbar() tea.Cmd {
	m.tabbar.SetScreen(m.screen)

	return nil
}

func (m model) InitKeymaps() tea.Cmd {
	keymaps.AddGroup("global", []string{"global"})

	return nil
}

func (m model) Init() tea.Cmd {

	return chain.Init(
		m.footer.Init,
		m.LoadBindings,
		m.toaster.Init,
		chain.OnError(m.screen.Init),
		m.InitTabbar,
		m.InitScreen,
		action.Init,
		m.InitKeymaps,
	)
}
