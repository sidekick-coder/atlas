package model

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/action"
	"github.com/sidekick-coder/atlas/tui/action/actions"
	"github.com/sidekick-coder/atlas/tui/components/formdialog"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/entrycontroller"
)

func (m *model) LoadDefaultActions() tea.Cmd {
	action.AddDefinition("entry-update", actions.EntryUpdateAction)

	action.AddDefinition("entry-sync", entrycontroller.SyncAction)
	action.AddDefinition("entry-set", entrycontroller.SetAction)

	action.AddDefinition("input", actions.InputAction)
	action.AddDefinition("formdialog", formdialog.Action)

	return nil
}

func (m *model) UpdateActions(msg tea.Msg) tea.Cmd {
	return chain.Update(
		msg,
		func(msg tea.Msg) tea.Cmd {
			return actions.HandleEntrySync(m.app, msg)
		},
		func(msg tea.Msg) tea.Cmd {
			return actions.HandleEntryUpdate(m.app, msg)
		},
		func(msg tea.Msg) tea.Cmd {
			return actions.HandleInput(action.Execute, msg)
		},
		func(msg tea.Msg) tea.Cmd {
			return formdialog.HandleAction(action.Execute, msg)
		},
	)
}
