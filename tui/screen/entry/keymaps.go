package entry

import (
	"path/filepath"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/bubbles/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

func (s *Screen) GetKeymapBindings() []key.Binding {
	keymaps := s.App.Config().GetKeymapsByGroup("entry-list")

	bindings := []key.Binding{}

	for _, km := range keymaps {
		b := key.NewBinding(
			key.WithKeys(km.Keys...),
			key.WithHelp(km.Keys[0], km.Description),
		)

		bindings = append(bindings, b)
	}

	return bindings
}


func (s *Screen) HandleKeyMaps(mgs tea.Msg) tea.Cmd {
	keyMsg, ok := mgs.(tea.KeyMsg)

	if !ok {
		return nil
	}

	keymaps := s.App.Config().GetKeymapsByGroup("entry-list")


	entry := s.List.SelectedEntry()

	ctx := s.App.ActionManager().CreateContext(map[string]any{
		"Entry": entry,
		"EntryID": entry.ID,
		"EntryPath": entry.Path,
		"EntryAbsolutePath": filepath.Join(s.App.WorkspacePath(), entry.Path),
	})

	for _, km := range keymaps {
		b := key.NewBinding(
			key.WithKeys(km.Keys...),
			key.WithHelp(km.Keys[0], km.Description),
		)

		if key.Matches(keyMsg, b) {
			return messages.ActionWithContextCmd(km.Action, ctx)
		}
	}

	return nil
}

