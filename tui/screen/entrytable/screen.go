package entrytable

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
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

	width  int
	height int

	loader    entryloader.Feature
	table     table.Component
	container container.Component
}

func Create(p tuimodels.ScreenPayload) (tuimodels.Screen, error) {
	s := &Screen{
		app:     p.App,
		options: p.Options,
		width:   100,
		height:  100,

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

	return messages.AddScreenCmd(messages.AddScreen{
		Name: "entry_single",
		Options: map[string]any{
			"path":  e.Path,
			"entry": e,
		},
	})
}

func (s *Screen) Init() tea.Cmd {
	limit := 10

	limit = max(limit, s.height-6)

	s.loader.SetLimit(limit)

	s.table.OnSelect(s.OpenEntry)

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

func (s *Screen) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(
		msg,
		s.table.Update,
		chain.OnKey(s.HadleBinding),
	)
}
