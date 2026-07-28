package entrysingle

import (
	"path/filepath"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/app/screen"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (s *Screen) Title() string {
	maxLength := 20

	baseName := filepath.Base(s.Path)

	if len(baseName) > maxLength {
		return baseName[:maxLength] + "..."
	}

	return baseName
}

func (s *Screen) SetSize(width, height int) {
	s.Width = width-2
	s.Height = height

	s.meta.Resize(s.Width, s.Height)
}

func (s *Screen) HandleSize(msg tea.Msg) tea.Cmd {
	if ss, ok := msg.(screen.SizeMsg); ok {
		s.SetSize(ss.Width, ss.Height)
	}

	return nil
}

func (s *Screen) Render() string {
	container := theme.BaseStyle().
		Width(s.Width).
		Height(s.Height).
		Margin(0, 1).
		BorderForeground(lipgloss.Color(theme.Current.Primary)).
		Border(lipgloss.NormalBorder())

	return container.Render(s.meta.Render())
}
