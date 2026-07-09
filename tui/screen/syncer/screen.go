package syncer

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/syncer"
	tuimodels "github.com/sidekick-coder/atlas/tui/models"
	"github.com/sidekick-coder/atlas/tui/screen/empty"
)

type Entry struct {
	Path    string
	Success bool
	Error   error
}

type Screen struct {
	App      *app.App
	Program  *tea.Program
	Width    int
	Height   int
	Running  bool
	ViewList bool
	Entries  []Entry
	Syncer   *syncer.Syncer

	Completed    bool
	TotalEntries int
	Time         time.Duration
}

func Create(p tuimodels.ScreenPayload) (tuimodels.Screen, error) {

	s := &Screen{
		App:       p.App,
		Syncer:    p.App.Syncer(),
		Program:   p.Program,
		Width:     100,
		Height:    100,
		Running:   false,
		Completed: false,
	}

	if i, ok := p.Options["immediate"].(bool); ok && i {
		s.Sync()
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
	if s.ViewList && len(s.Entries) > 0 {
		return s.RenderEntries()
	}

	if (s.Running || s.Completed){
		return s.RenderSummary()
	}

	content := ""

	content += "[e] to start syncing entries\n"
	content += "[E] to start syncing detailed view\n"

	return empty.Placeholder(empty.PlaceholderPayload{
		Width:  s.Width,
		Height: s.Height,
		Text:   content,
	})
}
