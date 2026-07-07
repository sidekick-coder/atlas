package syncer

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/sync/v2"
	tuimodels "github.com/sidekick-coder/atlas/tui/models"
	"github.com/sidekick-coder/atlas/tui/screen/empty"
)

type Entry struct {
	Path    string
	Success bool
	Error   error
}

type Screen struct {
	App     *app.App
	Program *tea.Program
	Width   int
	Height  int
	Running bool
	Entries []Entry
	Syncer  *sync.Sync
}

func Create(p tuimodels.ScreenPayload) (tuimodels.Screen, error) {
	syncer := p.App.Syncer()

	s := &Screen{
		App:     p.App,
		Syncer:  syncer,
		Program: p.Program,
		Width:   100,
		Height:  100,
		Running: false,
	}

	return s, nil
}

func (s *Screen) Title() string {
	return "Syncer"
}

func (s *Screen) Init() tea.Cmd {
	return nil
}

func (s *Screen) Update(msg tea.Msg) tea.Cmd {
	handlers := []func(tea.Msg) tea.Cmd{}

	handlers = append(
		handlers,
		s.HandleScreenKeymaps,
		s.HandleMessages,
		// s.HandleUserKeyMaps,
	)

	for _, handler := range handlers {
		cmd := handler(msg)

		if cmd != nil {
			return cmd
		}
	}

	return nil
}

func (s *Screen) SetSize(width, height int) {
	s.Width = width
	s.Height = height
}

func (s *Screen) Render() string {
	if len(s.Entries) > 0 {
		return s.RenderEntries()
	}

	return empty.Placeholder(empty.PlaceholderPayload{
		Width:  s.Width,
		Height: s.Height,
		Text:   "Press 'enter' to start syncing all entries.",
	})
}
