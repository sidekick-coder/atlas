package dialog

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/features/layer"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (c *Component) HandleView() tea.Cmd {
	c.layer.SetRender(c.render)

	c.LoadDefaultStyle()

	layer.Add(c.layer)
	return nil
}

func (c *Component) LoadDefaultStyle() {
	c.style = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Width(c.width).
		Height(c.height).
		BorderForeground(lipgloss.Color(theme.Current.Muted))
}

func (c *Component) GetWidth() int {
	return c.width
}

func (c *Component) render() string {
	x := (layer.ScreenWidth - c.width) / 2
	y := (layer.ScreenHeight - c.height) / 2

	c.layer.SetPosition(x, y)

	if !c.open {
		return ""
	}

	if c.onRender != nil {
		content := c.onRender()
		rendered := c.border.SetContent(content).SetLabel(c.title).SetSize(c.width, c.height).Render()
		return lipgloss.PlaceHorizontal(c.width, lipgloss.Center, rendered)
	}

	return c.style.
		Align(lipgloss.Center, lipgloss.Center).
		Render("No render function provided")
}
