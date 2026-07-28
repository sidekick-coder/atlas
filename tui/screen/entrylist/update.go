package entrylist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

func (s *Screen) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(
		msg,
		chain.OnCondition(s.list.Update, s.focus.IsFocused(s.list)),
		chain.OnCondition(s.UpdateView, s.focus.IsFocused(s.view)),
		chain.OnKey(s.HadleBinding),
		s.HandleSize,
	)
}

