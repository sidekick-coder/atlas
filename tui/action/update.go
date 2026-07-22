package action

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/action/actions"
)

func Update(msg tea.Msg) tea.Cmd {
	return actions.HandleEntrySyncMsg(manager.app, msg)
}
