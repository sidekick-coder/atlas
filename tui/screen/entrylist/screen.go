package entrylist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/utils/maputil"
	"github.com/sidekick-coder/atlas/tui/components/borderlabel"
	"github.com/sidekick-coder/atlas/tui/components/entrylist"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/component"
	"github.com/sidekick-coder/atlas/tui/features/context"
	"github.com/sidekick-coder/atlas/tui/features/focusmanager"
	tuimodels "github.com/sidekick-coder/atlas/tui/models"
)

type Screen struct {
	app           *app.App
	options       map[string]any
	viewComponent map[string]any

	width  int
	height int

	focus *focusmanager.Feature
	ctx   *context.Feature

	registry *component.Registry

	list      *entrylist.Component
	view      component.Component
	container *borderlabel.Component
}

func Create(p tuimodels.ScreenPayload) (tuimodels.Screen, error) {
	viewComponent := map[string]any{
		"type": "metas",
	}

	if vc, ok := maputil.GetMap(p.Options, "component"); ok {
		viewComponent = vc
	}

	c := context.Create()

	c.SetLabel("screen")

	c.SetAll(p.Options)

	c.Activate()

	s := &Screen{
		app:           p.App,
		options:       p.Options,
		viewComponent: viewComponent,

		width:  100,
		height: 100,

		focus:    focusmanager.Create(),
		ctx:      c,
		registry: component.CreateRegistry(),

		list:      entrylist.Create(),
		container: borderlabel.Create(),
	}

	return s, nil
}

func (s *Screen) Init() tea.Cmd {
	return chain.Init(
		s.LoadBindings,
		s.ctx.Init,
		s.InitList,
		s.InitView,
	)
}

func (s *Screen) Dispose() tea.Cmd {
	return chain.Dispose(
		s.UnloadBindings,
		s.ctx.Dispose,
		s.list.Dispose,
	)
}

