package entrylist

import (
	"maps"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/template"
	"github.com/sidekick-coder/atlas/internal/utils/maputil"
	"github.com/sidekick-coder/atlas/tui/components/borderlabel"
	"github.com/sidekick-coder/atlas/tui/components/entrylist"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/component"
	"github.com/sidekick-coder/atlas/tui/features/focusmanager"
	tuimodels "github.com/sidekick-coder/atlas/tui/models"
)

type Screen struct {
	app           *app.App
	options       map[string]any
	viewComponent map[string]any

	width  int
	height int

	focus    *focusmanager.Feature
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

	s := &Screen{
		app:           p.App,
		options:       p.Options,
		viewComponent: viewComponent,

		width:  100,
		height: 100,

		focus:    focusmanager.Create(),
		registry: component.CreateRegistry(),

		list:      entrylist.Create(),
		container: borderlabel.Create(),
	}

	return s, nil
}

func (s *Screen) Init() tea.Cmd {
	return chain.Init(
		s.LoadBindings,
		s.InitView,
		s.InitList,
		s.InitFocus,
	)
}

func (s *Screen) Dispose() tea.Cmd {
	return chain.Dispose(
		s.UnloadBindings,
		s.list.Dispose,
	)
}

func (s *Screen) Title() string {
	if pt, ok := s.options["title"].(string); ok {
		return pt
	}

	return "entries"
}

func (s *Screen) InitList() tea.Cmd {
	props := map[string]any{}

	if lk, ok := maputil.GetString(s.options, "label_key"); ok {
		props["label_key"] = lk
	}

	s.list.SetProps(props)

	return s.list.Init()
}

func (s *Screen) InitView() tea.Cmd {
	name := "metas"

	if n, ok := maputil.GetString(s.viewComponent, "type"); ok {
		name = n
	}

	view, ok := s.registry.Get(name)

	if !ok {
		return toast.Error("view with name " + name + " not found")
	}

	s.view = view

	s.focus.Add(s.view)

	return nil
}

func (s *Screen) InitFocus() tea.Cmd {
	s.focus.Add(s.list)

	s.focus.Focus(s.list)
	return nil
}

func (s *Screen) SetViewProps(payload map[string]any) {
	props := map[string]any{}

	maps.Copy(props, payload)

	cp := maputil.Except(s.viewComponent, "type")

	maps.Copy(props, cp)

	computed, err := template.EvaluateMap(props, props)

	if err == nil {
		props = computed
	}

	s.view.SetProps(props)
}
