package logs

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/logger"
	"github.com/sidekick-coder/atlas/tui/components/container"
	"github.com/sidekick-coder/atlas/tui/components/list"
	"github.com/sidekick-coder/atlas/tui/components/viewport"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/selection"
	"github.com/sidekick-coder/atlas/tui/features/theme"
	"github.com/sidekick-coder/atlas/tui/models"
)

type Entry struct {
	ID      string
	Options map[string]any
}

type Screen struct {
	Width  int
	Height int

	selection *selection.Feature

	list      list.Component
	container container.Component
	viewport  viewport.Component
}

func Create(p models.ScreenPayload) (models.Screen, error) {

	s := &Screen{
		Width:     100,
		Height:    100,
		selection: selection.Create(),

		list:      *list.Create(),
		container: *container.Create(),
		viewport: *viewport.Create(),
	}

	return s, nil
}

func (s *Screen) Title() string {
	return "logs"
}

func (s *Screen) Init() tea.Cmd {
	return chain.Init(chain.OnVoid(s.LoadBindings), s.list.Init, s.InitSelection, s.Load, s.list.Activate)
}

func (s *Screen) Dispose() tea.Cmd {
	return chain.Dispose(chain.OnVoid(s.UnloadBindings), s.list.Dispose)
}

func (s *Screen) Level(v string) string {
	if v == logger.Levels.Error {
		return theme.BaseStyle().Foreground(theme.Error()).Render(v)
	}

	if v == logger.Levels.Warn {
		return theme.BaseStyle().Foreground(theme.Warning()).Render(v)
	}

	if v == logger.Levels.Info {
		return theme.BaseStyle().Foreground(theme.Info()).Render(v)
	}

	if v == logger.Levels.Debug {
		return theme.BaseStyle().Foreground(theme.Primary()).Render(v)
	}

	return v
}


func (s *Screen) Load() tea.Cmd {
	items := []string{}

	logs, _ := logger.List()

	muted := theme.BaseStyle().Foreground(theme.Muted())

	for _, log := range logs {
		item := fmt.Sprintf("[%s] %s: %s", log.Time, s.Level(log.Level), log.Msg)

		for key, value := range log.Options {
			item += muted.Render(fmt.Sprintf(" %s=%v", key, value))
		}

		items = append(items, item)
	}

	s.list.SetItems(items)
	s.list.SetSelection(s.selection)
	s.selection.SetTotal(len(items))

	return s.list.Init()
}
