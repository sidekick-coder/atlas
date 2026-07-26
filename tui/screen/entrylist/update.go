package entrylist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/action"
	"github.com/sidekick-coder/atlas/tui/components/entrylist"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/entrycontroller"
	"github.com/sidekick-coder/atlas/tui/features/keymaps"
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
		action.RemoveContext("entry")

		props := map[string]any{}

		if c.Exists {

			e := entrycontroller.CreateContext(c.Entry)

			props["entry"] = e

			ctx := map[string]any{
				"entry": e,
			}

			action.AddContext("entry", ctx)

			groups := keymaps.MapToGroups(e)

			keymaps.AddGroup("entry-table", groups)
		}

		s.SetViewProps(props)
	}

	return nil
}
