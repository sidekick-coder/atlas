package entrylist

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/app/screen"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (s *Screen) Title() string {
	if pt, ok := s.options["title"].(string); ok {
		return pt
	}

	return "entries"
}

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

func (s *Screen) RenderView() string {
	leftWidth := s.width * 2 / 5          // 40%
	rightWidth := s.width - leftWidth - 8 // 60%
	rightColor := theme.Current.Border

	content := ""

	if s.view != nil {
		if s.focus.IsFocused(s.view) {
			rightColor = theme.Current.Primary
		}

		s.view.Resize(rightWidth, s.height)
		content = s.view.Render()
	}

	return s.container.
		SetLabel("view").
		SetSize(rightWidth, s.height).
		SetContent(content).
		SetColor(rightColor).
		Render()
}

func (s *Screen) RenderList() string {
	leftWidth := s.width * 2 / 5 // 40%

	s.list.SetSize(leftWidth, s.height)

	letColor := theme.Current.Border

	if s.focus.IsFocused(s.list) {
		letColor = theme.Current.Primary
	}

	return s.container.
		SetLabel(s.Title()).
		SetColor(letColor).
		SetSize(leftWidth, s.height).
		SetContent(s.list.Render()).
		Render()
}

func (s *Screen) Render() string {
	return lipgloss.JoinHorizontal(lipgloss.Left, s.RenderList(), s.RenderView())
}
