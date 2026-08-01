package form

import (
	"charm.land/lipgloss/v2"
)

func (c *Component) Resize(width, height int) {
	c.width = width
	c.height = height

	for _, i := range c.fields {
		i.Resize(width, height)
	}
}

func (c *Component) Render() string {
	var lines []string

	for _, i := range c.fields {
		content := i.Render()

		lines = append(lines, content)
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}
