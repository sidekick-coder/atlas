package entrylist

import (
	"maps"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/utils/maputil"
	"github.com/sidekick-coder/atlas/tui/components/toast"
)

func (s *Screen) InitView() tea.Cmd {
	name := "metas"

	if n, ok := maputil.GetString(s.viewComponent, "type"); ok {
		name = n
	}

	view, ok := s.registry.Get(name)

	if !ok {
		return toast.Error("view with name " + name + " not found")
	}

	vctx := view.Context()

	vctx.SetParent(s.ctx)

	s.view = view
	s.view.Init()
	s.focus.Add(s.view)

	return nil
}

func (s *Screen) DisposeView() tea.Cmd {
	if s.view != nil {
		s.focus.Remove(s.view)
		s.view.Dispose()
	}

	return nil
}

func (s *Screen) UpdateView(msg tea.Msg) tea.Cmd {
	if s.view != nil {
		return s.view.Update(msg)
	}

	return nil
}

func (s *Screen) LoadView()  {
	e, ok := s.list.GetCurrent()

	if !ok {
		return
	}

	s.SetViewProps(map[string]any{
		"entry": e.ToMap(),
	})

}

func (s *Screen) SetViewProps(payload map[string]any) {
	props := map[string]any{}

	maps.Copy(props, payload)

	cp := maputil.Except(s.viewComponent, "type")

	maps.Copy(props, cp)

	s.view.SetProps(props)
}
