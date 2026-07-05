package empty

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/models"
)

type Screen struct {
	Width  int
	Height int
}

func Create(payload models.ScreenPayload) (models.Screen, error) {
	s := &Screen{
		Width:  100,
		Height: 100,
	}

	return s, nil
}

func (s *Screen) Title() string {
	return "empty"
}

func (s *Screen) Init() tea.Cmd {
	// Initialization logic for the Entry Screen
	return nil
}

func (s *Screen) SetSize(width, height int) {
	s.Width = width
	s.Height = height
}

func (s *Screen) Update(msg tea.Msg) tea.Cmd {
	handlers := []func(tea.Msg) tea.Cmd{}

	handlers = append(handlers,
		s.HandleKeyPress,
	)

	for _, handler := range handlers {
		cmd := handler(msg)

		if cmd != nil {
			return cmd
		}
	}

	return nil
}
