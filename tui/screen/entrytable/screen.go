package entrytable

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/components/container"
	"github.com/sidekick-coder/atlas/tui/components/table"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	tuimodels "github.com/sidekick-coder/atlas/tui/models"
)

type Screen struct {
	app    *app.App
	width  int
	height int
	limit  int
	query  string
	table  table.Component
	container container.Component
}

func Create(payload tuimodels.ScreenPayload) (tuimodels.Screen, error) {
	s := &Screen{
		app:    payload.App,
		width:  100,
		height: 100,
		limit:  30,
		query:  "",
		table:  *table.Create(),
		container: *container.Create(),
	}

	if payload.Options["query"] != nil {
		if query, ok := payload.Options["query"].(string); ok {
			s.query = query
		}
	}

	return s, nil
}

func (s *Screen) Title() string {
	return "tables"
}

func (s *Screen) Init() tea.Cmd {
	return chain.Init(
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
