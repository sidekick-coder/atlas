package entrytable

import (
	tea "charm.land/bubbletea/v2"
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

	return nil
}
