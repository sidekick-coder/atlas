package actions

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/components/toast"
)

type EntrySyncMsg struct {
	Path string
}

type EntrySyncEndMsg struct {
	Path string
}

func EntrySyncAction(ctx map[string]any) (map[string]any, error) {
	result := make(map[string]any)

	path, ok := ctx["path"].(string)

	if !ok || path == "" {
		return result, fmt.Errorf("invalid or missing 'path' in context")
	}

	msg := EntrySyncMsg{Path: path}

	result["tea_message"] = msg

	return result, nil
}

func HandleEntrySync(app *app.App, msg tea.Msg) tea.Cmd {
	if m, ok := msg.(EntrySyncMsg); ok {
		err := app.Syncer().One(m.Path)

		if err != nil {
			return toast.Error(fmt.Sprintf("Failed to sync entry: %v", err))
		}

		return func() tea.Msg {
			return EntrySyncEndMsg{Path: m.Path}
		}
	}

	return nil
}
