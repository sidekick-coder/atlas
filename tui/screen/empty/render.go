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
		Width(p.Width - 4).
		Height(p.Height - 4).
		Margin(0, 2).
		Align(lipgloss.Center, lipgloss.Center).
		BorderForeground(lipgloss.Color("12"))

	content := "Empty Screen"

	if p.Text != "" {
		content = p.Text
	}

	return border.Render(content)
}

func (s *Screen) Render() string {
	return Placeholder(PlaceholderPayload{
		Width:  s.Width,
		Height: s.Height,
		Text:  "Press 'e' to view the entry list. Press 'q' to quit.",
	})
}
