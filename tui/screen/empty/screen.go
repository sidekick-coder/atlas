package empty

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/container"
	"github.com/sidekick-coder/atlas/tui/components/list"
	"github.com/sidekick-coder/atlas/tui/messages"
	"github.com/sidekick-coder/atlas/tui/models"
)

type Entry struct {
	ID      string
	Options map[string]any
}

type Screen struct {
	Width   int
	Height  int
	Entries []Entry

	list list.Component
	container container.Component
}

func (s *Screen) HandleSelection(index int) tea.Cmd {
	if index < 0 || index >= len(s.Entries) {
		return nil
	}

	entry := s.Entries[index]

	if entry.Options == nil {
		return nil
	}

	return messages.ReplaceScreenCmd(messages.ReplaceCurrentScreen{
		Name:   entry.ID,
		Options: entry.Options,
	})
}

func Create(p models.ScreenPayload) (models.Screen, error) {
	entries := []Entry{}

	if e, ok := p.Options["entries"].([]Entry); ok {
		entries = e
	}

	s := &Screen{
		Width:   100,
		Height:  100,
		Entries: entries,
		list:    *list.Create(),
		container: *container.Create(),
	}


	return s, nil
}

func (s *Screen) Title() string {
	return "empty"
}

func (s *Screen) Init() tea.Cmd {
	items := []string{}

	for _, entry := range s.Entries {
		items = append(items, entry.ID)
	}

	s.list.SetItems(items)
	s.list.OnSelect(s.HandleSelection)

	return nil
}

func (s *Screen) Dispose() tea.Cmd {
	return nil
}
