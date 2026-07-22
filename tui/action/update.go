package action

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/action/actions"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

func Update(msg tea.Msg) tea.Cmd {
	return chain.Update(
		msg,
		func(msg tea.Msg) tea.Cmd {
			return actions.HandleEntrySync(manager.app, msg)
		},
		func(msg tea.Msg) tea.Cmd {
			return  actions.HandleInput(Execute, msg)
		},
	)
}
