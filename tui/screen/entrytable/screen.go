package entrytable

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/components/container"
	"github.com/sidekick-coder/atlas/tui/components/table"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/entryloader"
	tuimodels "github.com/sidekick-coder/atlas/tui/models"
)

type Screen struct {
	app    *app.App
	options map[string]any

	width  int
	height int

	loader entryloader.Feature
	table  table.Component
	container container.Component
}

func Create(p tuimodels.ScreenPayload) (tuimodels.Screen, error) {
	s := &Screen{
		app:    p.App,
		options: p.Options,
		width:  100,
		height: 100,

		loader: *entryloader.Create(*p.App.EntryRepo()),
		table:  *table.Create(),
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

func (s *Screen) Init() tea.Cmd {
	return chain.Init(
		s.loader.Init,
		s.RegisterBindings,
		s.LoadColumns,
		s.table.Init,
	)
}

func (s *Screen) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(
		msg,
		s.table.Update,
		s.HandleKeypress,
	)
}
