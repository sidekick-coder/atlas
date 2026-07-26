package entrylist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/entrylist"
	"github.com/sidekick-coder/atlas/tui/features/chain"
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

func (s *Screen) UpdateView(msg tea.Msg) tea.Cmd {
	if s.view != nil {
		return s.view.Update(msg)
	}

	return nil
}

func (s *Screen) HandleMessage(msg tea.Msg) tea.Cmd {
	if c, ok := msg.(entrylist.ChangedMsg); ok {
		props := map[string]any{}

		if c.Exists {
			props["entry"] = c.Entry.ToMap()
		}

		s.SetViewProps(props)
	}

	return nil
}
