package empty

import (
	"log/slog"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/app/screen"
	"github.com/sidekick-coder/atlas/tui/components/container"
	"github.com/sidekick-coder/atlas/tui/components/list"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/selection"
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

	list      list.Component
	selection *selection.Feature
	container container.Component
}

func (s *Screen) Select(index int) tea.Cmd {
	slog.Debug("empty selecting entry", "index", index)
	if index < 0 || index >= len(s.Entries) {
		return nil
	}

	entry := s.Entries[index]

	if entry.Options == nil {
		return nil
	}

	return screen.ReplaceCurrent(entry.ID, entry.Options)
}

func Create(p models.ScreenPayload) (models.Screen, error) {
	entries := []Entry{}

	if e, ok := p.Options["entries"].([]Entry); ok {
		entries = e
	}

	s := &Screen{
		Width:     100,
		Height:    100,
		Entries:   entries,

		selection: selection.Create(),

		list:      *list.Create(),
		container: *container.Create(),
	}

	return s, nil
}

func (s *Screen) Title() string {
	return "empty"
}

func (s *Screen) Init() tea.Cmd {
	return chain.Init(chain.OnVoid(s.LoadBindings), s.list.Init, s.Load, s.list.Activate)
}

func (s *Screen) Dispose() tea.Cmd {
	return chain.Dispose(chain.OnVoid(s.UnloadBindings), s.list.Dispose)
}

func (s *Screen) Load() tea.Cmd {
	items := []string{}

	for _, entry := range s.Entries {
		items = append(items, entry.ID)
	}

	s.list.SetItems(items)
	s.list.SetSelection(s.selection)
	s.selection.SetTotal(len(items))

	return s.list.Init()
}
