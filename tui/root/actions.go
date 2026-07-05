package root

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/messages"
	"github.com/charmbracelet/bubbles/key"
)

func (m *model) actionBindingMessageHandler(mgs tea.Msg) tea.Cmd {
	keyMsg, ok := mgs.(tea.KeyMsg)

	if !ok {
		return nil
	}

	keymaps := m.app.Config().GetKeymapsByGroup("global")

	for _, km := range keymaps {
		b := key.NewBinding(
			key.WithKeys(km.Keys...),
			key.WithHelp(km.Keys[0], km.Description),
		)

		if key.Matches(keyMsg, b) {
			return messages.ActionCmd(km.Action)
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

	err := am.Execute(a.Name, a.Context)

	if (err != nil) {
		return messages.ToastErrorCmd("Error executing action: " + err.Error(), 3 * 1000)
	}

	return  nil
}
