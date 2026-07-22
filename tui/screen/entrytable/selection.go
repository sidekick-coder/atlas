package entrytable

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/table"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/action"
	"github.com/sidekick-coder/atlas/tui/features/keymaps"
)

func (s *Screen) InitSelection() tea.Cmd {
	return nil
}

func (s *Screen) HandleSelection(msg tea.Msg) tea.Cmd {
	if c, ok := msg.(table.ChangeMsg); ok {
		action.RemoveContext("entry")

		entry, err := s.loader.GetEntry(c.Index)

		if err != nil {
			return toast.Error(fmt.Sprintf("Failed to get entry: %v", err))
		}

		e := s.CreateEntryContext(entry)

		ctx := map[string]any{
			"entry": e,
		}

		action.AddContext("entry", ctx)

		groups := keymaps.MapToGroups(e)

		keymaps.AddGroup("entry-table", groups)
	}

	return nil
}

func (s *Screen) DisposeSelection() tea.Cmd {
	action.RemoveContext("entry")
	keymaps.RemoveGroup("entry-table")

	return nil
}
