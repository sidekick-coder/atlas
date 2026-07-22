package model

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/action"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

func (m *model) HandleUserBindings(mgs tea.KeyMsg) tea.Cmd {
	for _, b := range UserBindings {
		if key.Matches(b) {
			actionId := b.GetMeta("action")

			if actionId == nil {
				return toast.Error("No action defined for key binding: " + b.GetDescription())
			}

			return action.Execute(actionId.(string))
		}
	}

	return nil
}

func (m *model) HandleActions(msg tea.Msg) tea.Cmd {
	a, ok := msg.(messages.Action)

	if !ok {
		return nil
	}

	am := m.app.ActionManager()

	ctx := am.CreateContext()

	err := am.Execute(a.Name, ctx)

	if err != nil {
		return messages.ToastErrorCmd("Error executing action: " + err.Error())
	}

	return nil
}
