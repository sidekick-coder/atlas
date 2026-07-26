package entrylist

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/app/screen"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (s *Screen) HandleSize(msg tea.Msg) tea.Cmd {
	if ss, ok := msg.(screen.SizeMsg); ok {
		s.SetSize(ss.Width, ss.Height)
	}

	return nil
}

func (s *Screen) SetSize(width, height int) {
	s.width = width
	s.height = height
	s.container.SetColor(theme.Current.Border)

	limit := 20

	limit = max(limit, s.height)
}

func (s *Screen) Render() string {
	leftWidth := s.width * 2 / 5          // 40%
	rightWidth := s.width - leftWidth - 8 // 60%

	s.list.SetSize(leftWidth, s.height)
	s.view.SetSize(rightWidth, s.height)

	letColor := theme.Current.Border
	rightColor := theme.Current.Border

	if s.focus.IsFocused(s.list) {
		letColor = theme.Current.Primary
	}

	if s.focus.IsFocused(s.view) {
		rightColor = theme.Current.Primary
	}

	left := s.container.
		SetLabel("Entries").
		SetColor(letColor).
		SetSize(leftWidth, s.height).
		SetContent(s.list.Render()).
		Render()

	right := s.container.
		SetLabel("Details").
		SetSize(rightWidth, s.height).
		SetContent(s.view.Render()).
		SetColor(rightColor).
		Render()

	return lipgloss.JoinHorizontal(lipgloss.Left, left, right)
}
