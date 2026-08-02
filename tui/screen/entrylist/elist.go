package entrylist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/selection"
)

func (s *Screen) InitList() tea.Cmd {
	lctx := s.list.Context()

	lctx.SetParent(s.ctx)

	s.list.SetProps(s.options)

	s.focus.Add(s.list)
	s.focus.Focus(s.list)

	s.ctx.SetAll(s.options)

	lselection := s.list.GetSelection()

	lselection.Change.On(func(event selection.ChangeEvent) {
		s.LoadView()
	})

	return s.list.Init()
}
