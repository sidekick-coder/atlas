package dialog

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/features/layer"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (c *Component) InitView() tea.Cmd {
	c.layer.SetRender(c.Render)

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

func (c *Component) GetSize() (int, int) {
	return c.width, c.height
}

func (c *Component) SetSize(width, height int) *Component {
	c.width = width
	c.height = height
	return c
}

func (c *Component) SetWidth(width int) *Component {
	c.width = width
	return c
}

func (c *Component) SetZIndex(z int) *Component {
	c.layer.SetZIndex(z)
	return c
}

func (c *Component) GetContentSize() (int, int) {
	return c.width - 4, c.height - 2
}

func (c *Component) Render() string {
	x := (layer.ScreenWidth - c.width) / 2
	y := (layer.ScreenHeight - c.height) / 2

	c.layer.SetPosition(x, y)

	if !c.open {
		return ""
	}

	content := c.contentRender() 

	return theme.BaseStyle().
		Width(c.width).
		Render(content)
}
