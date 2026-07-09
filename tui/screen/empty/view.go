package empty

import (
	lipgloss "charm.land/lipgloss/v2"
)

type PlaceholderPayload struct {
	Width  int
	Height int
	Text   string
}

func Placeholder(p PlaceholderPayload) string {
	border := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Width(p.Width-4).
		Height(p.Height-4).
		Margin(0, 2).
		Align(lipgloss.Center, lipgloss.Center).
		BorderForeground(lipgloss.Color("12"))

	content := "Empty Screen"

	if p.Text != "" {
		content = p.Text
	}

	return border.Render(content)
}

func (s *Screen) SetSize(width, height int) {
	s.Width = width
	s.Height = height

	s.list.SetSize(width, 1)

	s.container.
		SetSize(width-4, height).
		SetBorder("12").
		SetMargin(0, 2, 0, 2)
}

func (s *Screen) Render() string {
	list := s.list.Render()

	s.container.SetContent(list)

	return s.container.Render()
}
