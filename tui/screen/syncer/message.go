package syncer

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/messages"
)

type EntryAdd struct {
	AddEntry Entry
}

func (s *Screen) HandleMessages(msg tea.Msg) tea.Cmd {
	if eam, ok := msg.(EntryAdd); ok {
		s.Entries = append(s.Entries, eam.AddEntry)
		return messages.SkipCmd()
	}

	return nil
}
