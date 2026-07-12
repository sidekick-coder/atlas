package custom

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)


func (s *Screen) Update(msg tea.Msg) tea.Cmd {
	handlers := []func(msg tea.Msg) tea.Cmd{}

	handlers = append(
		handlers,
		s.HandleSize,
		chain.OnKey(s.HandleBinding),
	)

	current, ok := s.GetCurrent()

	if ok {
		handlers = append(handlers, current.Definition.Update)
	}

	return chain.Update(msg, handlers...)
}
