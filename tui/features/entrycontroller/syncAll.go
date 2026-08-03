package entrycontroller

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/logger"
	"github.com/sidekick-coder/atlas/internal/syncer"
	"github.com/sidekick-coder/atlas/tui/app/program"
)

type SyncAllMsg struct {
}

type SyncedAllMsg struct {
}

func SyncAll() tea.Cmd {
	return program.Command(SyncAllMsg{})
}

func (f *Feature) HandleAllSync(msg tea.Msg) tea.Cmd {
	s := f.app.Syncer()

	if _, ok := msg.(SyncAllMsg); ok {

		return func() tea.Msg {
			logger.Debug("SyncAllMsg received, starting sync all operation")

			s.OnComplete(func(result syncer.Result) {
				logger.Debug("Sync all operation completed, sending SyncedAllMsg")
			})

			s.OnError(func(path string, err error) {
				logger.Error("Sync all operation failed: %v", err)
			})

			s.All()

			return SyncedAllMsg{}
		}
	}

	return nil
}

func SyncAllAction(ctx map[string]any) (map[string]any, error) {
	result := make(map[string]any)

	msg := SyncAllMsg{}

	result["tea_message"] = msg

	return result, nil
}
