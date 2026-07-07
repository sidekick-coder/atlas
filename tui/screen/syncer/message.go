package syncer

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/messages"
)

type EntryAdd struct {
	AddEntry Entry
}

type EntryClear struct{}

type Completed struct {
	Time         time.Duration
	TotalEntries int
}

func (s *Screen) HandleMessages(msg tea.Msg) tea.Cmd {
	if eam, ok := msg.(EntryAdd); ok {
		s.Entries = append(s.Entries, eam.AddEntry)
		return messages.SkipCmd()
	}

	if _, ok := msg.(EntryClear); ok {
		s.Entries = []Entry{}
		s.Completed = false
		return messages.SkipCmd()
	}

	if c, ok := msg.(Completed); ok {
		s.Running = false
		s.Completed = true
		s.Time = c.Time
		s.TotalEntries = c.TotalEntries
		return messages.SkipCmd()
	}

	return nil
}
