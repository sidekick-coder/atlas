package entrytable

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/app/screen"
	"github.com/sidekick-coder/atlas/tui/components/container"
	"github.com/sidekick-coder/atlas/tui/components/table"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/entryloader"
	"github.com/sidekick-coder/atlas/tui/messages"
	tuimodels "github.com/sidekick-coder/atlas/tui/models"
)

type Screen struct {
	app     *app.App
	options map[string]any
	width   int
	height  int

	openScreen  string
	openOptions map[string]any

	loader    entryloader.Feature
	table     table.Component
	container container.Component
}

func Create(p tuimodels.ScreenPayload) (tuimodels.Screen, error) {
	openScreen := "entry_single"

	if os, ok := p.Options["open_screen"].(string); ok {
		openScreen = os
	}

	s := &Screen{
		app:     p.App,
		options: p.Options,
		width:   100,
		height:  100,

		openScreen: openScreen,

		loader:    *entryloader.Create(*p.App.EntryRepo()),
		table:     *table.Create(),
		container: *container.Create(),
	}

	return s, nil
}

func (s *Screen) Title() string {
	pt, ok := s.options["title"].(string)

	if ok {
		return pt
	}

	return "tables"
}

func (s *Screen) OpenEntry(cursor int) tea.Cmd {

	e, err := s.loader.GetEntry(cursor)

	if err != nil {
		return messages.ToastErrorCmd(err.Error())
	}

	return screen.Add(s.openScreen, map[string]any{
		"entry": e.ToMap(),
	})
}

func (s *Screen) Init() tea.Cmd {
	s.table.OnSelect(s.OpenEntry)

	q := s.options["query"]

	if q != nil {
		s.loader.SetQuery([]string{q.(string)})
	}

	return chain.Init(
		s.loader.Init,
		s.LoadBindings,
		s.LoadColumns,
		s.table.Init,
	)
}

func (s *Screen) Dispose() tea.Cmd {
	return chain.Dispose(
		s.table.Dispose,
		s.UnloadBindings,
	)
}
