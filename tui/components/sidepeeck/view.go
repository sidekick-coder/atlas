package sidepeeck

import (
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (c *Component) LoadDefaultStyle() {
	c.style = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		Width(c.width).
		Height(c.height).
		BorderBackground(lipgloss.Color(theme.Current.Background)).
		BorderForeground(lipgloss.Color(theme.Current.Primary))
}

func (c *Component) GetWidth() int {
	return c.width
}

func (c *Component) Render() string {
	if !c.open {
		return ""
	}

	if c.onRender != nil {
		return c.style.Render(c.onRender())
	}

	return c.style.
		Align(lipgloss.Center,lipgloss.Center).
		Render("No render function provided")
}
