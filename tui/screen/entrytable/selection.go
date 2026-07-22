package entrytable

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/table"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/action"
)

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
	}

	return nil
}
