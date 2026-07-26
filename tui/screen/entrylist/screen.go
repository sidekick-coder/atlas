package entrylist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/components/borderlabel"
	"github.com/sidekick-coder/atlas/tui/components/entrylist"
	"github.com/sidekick-coder/atlas/tui/components/entrymeta"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/focusmanager"
	tuimodels "github.com/sidekick-coder/atlas/tui/models"
)

type Screen struct {
	app     *app.App
	options map[string]any

	width  int
	height int

	focus *focusmanager.Feature

	list      *entrylist.Component
	view      *entrymeta.Component
	container *borderlabel.Component
}

func Create(p tuimodels.ScreenPayload) (tuimodels.Screen, error) {
	s := &Screen{
		app:     p.App,
		options: p.Options,
		width:   100,
		height:  100,

		focus: focusmanager.Create(),

		list:      entrylist.Create(),
		container: borderlabel.Create(),
		view:      entrymeta.Create(""),
	}

	return s, nil
}

func (s *Screen) Init() tea.Cmd {
	return chain.Init(
		s.LoadBindings,
		s.list.Init,
		s.view.Init,
		s.InitFocus,
	)
}

func (s *Screen) Dispose() tea.Cmd {
	return chain.Dispose(
		s.UnloadBindings,
		s.list.Dispose,
		s.view.Dispose,
	)
}

func (s *Screen) Title() string {
	if pt, ok := s.options["title"].(string); ok {
		return pt
	}

	return "entries"
}

func (s *Screen) InitFocus() tea.Cmd {
	s.focus.Add(s.list)
	s.focus.Add(s.view)

	s.focus.Focus(s.list)
	return nil
}
