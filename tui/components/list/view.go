package list

import (
	lipgloss "charm.land/lipgloss/v2"
)

func (c *Component) SetSize(w, h int) {
	c.width = w
	c.height = h
}

func (c *Component) Render() string {
	normal := lipgloss.NewStyle().
		Width(c.width).
		Height(c.height)

	focus := lipgloss.NewStyle().
		Width(c.width).
		Height(c.height).
		Background(lipgloss.Color("12")).
		Foreground(lipgloss.Color("0"))

	var items []string

	for index, item := range c.items {
		if index == c.cursor {
			result := focus.Render(item)

			items = append(items, result)
			continue
		}

		items = append(items, normal.Render(item))
	}

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}
