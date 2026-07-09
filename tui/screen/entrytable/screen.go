package entrytable

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	tuimodels "github.com/sidekick-coder/atlas/tui/models"
)

type Screen struct {
	App    *app.App
	Width  int
	Height int
	Limit  int
	Query  string
}

func Create(payload tuimodels.ScreenPayload) (tuimodels.Screen, error) {
	s := &Screen{
		App:    payload.App,
		Width:  100,
		Height: 100,
		Limit:  30,
		Query:  "",
	}

	if payload.Options["query"] != nil {
		if query, ok := payload.Options["query"].(string); ok {
			s.Query = query
		}
	}

	return s, nil
}

func (s *Screen) Title() string {
	return "tables"
}

func (s *Screen) Init() tea.Cmd {
	return chain.Init(s.RegisterBindings)
}

func (s *Screen) SetSize(width, height int) {
	s.Width = width
	s.Height = height
}

func (s *Screen) Render() string {
	return ""
}

func (s *Screen) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(msg, s.HandleKeypress)
}
