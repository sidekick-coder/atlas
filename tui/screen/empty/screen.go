package empty

import (
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/models"
)

type Screen struct {
	Width  int 
	Height int
}

func Create(payload models.ScreenPayload) *Screen {
	s := &Screen{
		Width:  100,
		Height: 100,
	}

	return s
}


func (s *Screen) Title() string {
	return "Empty Screen"
}

func (s *Screen) Init() tea.Cmd {
	// Initialization logic for the Entry Screen
	return nil
}

func (s *Screen) SetSize(width, height int) {
	s.Width = width
	s.Height = height
}

func (s *Screen) Load() error {
	return nil
}

func (s *Screen) Render() string {
	border := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Width(s.Width - 4).
		Height(s.Height - 4).
		Margin(0, 2).
		Align(lipgloss.Center, lipgloss.Center).
		BorderForeground(lipgloss.Color("12"))

	content := "Empty Screen"

	return border.Render(content)
}
