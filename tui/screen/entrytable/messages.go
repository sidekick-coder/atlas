package entrytable

import (
	"log/slog"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/action/actions"
	"github.com/sidekick-coder/atlas/tui/components/toast"
)

type ReloadMsg struct{}

func Reload() tea.Msg {
	return ReloadMsg{}
}

func (s *Screen) HandleMessage(msg tea.Msg) tea.Cmd {
	if _, ok := msg.(ReloadMsg); ok {
		s.loader.Load()
		return toast.Success("Reloaded")
	}

	if _, ok := msg.(actions.EntrySyncEndMsg); ok {
		slog.Info("Entry sync completed, reloading table")
		s.loader.Load()
	}

	return nil
}
