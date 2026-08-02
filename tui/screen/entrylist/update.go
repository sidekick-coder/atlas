package entrylist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/action/actions"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/entrycontroller"
)

func (s *Screen) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(
		msg,
		chain.OnCondition(s.list.Update, s.focus.IsFocused(s.list)),
		chain.OnCondition(s.UpdateView, s.focus.IsFocused(s.view)),
		chain.OnKey(s.HadleBinding),
		s.HandleSize,
		s.HandleMessage,
	)
}

func (s *Screen) HandleMessage(msg tea.Msg) tea.Cmd {
	if _, ok := msg.(entrycontroller.UpdatedMsg); ok {
		s.LoadView()
	}

	if _, ok := msg.(entrycontroller.SyncedMsg); ok {
		s.LoadView()
	}

	if _, ok := msg.(actions.EntrySyncEndMsg); ok {
		s.LoadView()
	}

	return nil
}
