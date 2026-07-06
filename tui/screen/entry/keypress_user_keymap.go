package entry

import (
	"path/filepath"

	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

func (s *Screen) HandleUserKeyMaps(mgs tea.Msg) tea.Cmd {
	keyMsg, ok := mgs.(tea.KeyMsg)

	if !ok {
		return nil
	}

	keymaps := s.GetUserKeymaps()

	if len(s.List.Entries) == 0 {
		return nil
	}

	ctx := map[string]any{}

	entry, selected := s.List.SelectedEntry()

	if selected {
		ctx["Entry"] = entry
		ctx["EntryID"] = entry.ID
		ctx["EntryPath"] = entry.Path
		ctx["EntryAbsolutePath"] = filepath.Join(s.App.WorkspacePath(), entry.Path)
	}

	ctx = s.App.ActionManager().CreateContext(ctx)

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
