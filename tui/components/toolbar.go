package components

import (
	lipgloss "charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

type Toolbar struct {
	Title string
	Width int
}

func NewToolbar() *Toolbar {
	return &Toolbar{}
}

func (t *Toolbar) SetTitle(title string) {
	t.Title = title
}

func (t *Toolbar) SetWidth(width int) {
	t.Width = width
}

func (t *Toolbar) Render() string {
	border := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Width(t.Width - 4).
		Margin(0, 2).
		Padding(0, 2).
		BorderForeground(lipgloss.Color(theme.Current.Primary))

	row := lipgloss.NewStyle().
		Render(t.Title)

	return border.Render(row)
}

func (t *Toolbar) GetHeight() int {
	return lipgloss.Height(t.Render())
}

func (t *Toolbar) GetWidth() int {
	return lipgloss.Width(t.Render())
}


