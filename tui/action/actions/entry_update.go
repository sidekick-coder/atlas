package actions

import (
	"fmt"
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/components/toast"
)

type EntryUpdateMsg struct {
	Path   string
	Values map[string]any
}

type EntryUpdateEndMsg struct {
	Path   string
	Values map[string]any
}

func EntryUpdateAction(ctx map[string]any) (map[string]any, error) {
	result := make(map[string]any)

	msg := EntryUpdateMsg{}

	if path, ok := ctx["path"].(string); ok {
		msg.Path = path
	}

	if values, ok := ctx["values"].(map[string]any); ok {
		msg.Values = values
	}

	result["tea_message"] = msg

	return result, nil
}

func HandleEntryUpdate(app *app.App, msg tea.Msg) tea.Cmd {
	if m, ok := msg.(EntryUpdateMsg); ok {
		err := app.Syncer().One(m.Path) // Sync the entry before updating

		if err != nil {
			return toast.Error(fmt.Sprintf("Failed to sync entry: %v", err))
		}

		for k, v := range m.Values {
			err := app.SetEntryMeta(m.Path, k, fmt.Sprintf("%v", v))

			if err != nil {
				return toast.Error(fmt.Sprintf("Failed to update entry meta: %v", err))
			}
		}

		return func() tea.Msg {
			return EntryUpdateEndMsg{Path: m.Path, Values: m.Values}
		}
	}

	return nil
}
