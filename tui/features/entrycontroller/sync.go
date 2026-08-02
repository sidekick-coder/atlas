package entrycontroller

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/components/toast"
)

type SyncMsg struct {
	Path string
}

type SyncedMsg struct {
	Path string
}

func Sync(path string) tea.Cmd {
	return program.Command(SyncMsg{Path: path})
}

func syned(path string) tea.Cmd {
	return program.Command(SyncedMsg{Path: path})
}

func (f *Feature) HandleSync(msg tea.Msg) tea.Cmd {
	if m, ok := msg.(SyncMsg); ok {
		err := f.app.Syncer().One(m.Path)

		if err != nil {
			return toast.Error(fmt.Sprintf("Failed to sync entry: %v", err))
		}

		return tea.Batch(
			syned(m.Path),
			updated(m.Path),
		)
	}

	return nil
}


func SyncAction(ctx map[string]any) (map[string]any, error) {
	result := make(map[string]any)

	msg := SyncMsg{}

	if path, ok := ctx["path"].(string); ok {
		msg.Path = path
	}

	result["tea_message"] = msg

	return result, nil
}

