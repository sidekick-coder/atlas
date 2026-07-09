package table

import (
	lipgloss "charm.land/lipgloss/v2"
)

func (c *Component) SetSize(w, h int) {
	c.width = w
	c.height = h
}

func (c *Component) Render() string {
	border := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Width(c.width-4).
		Height(c.height-4).
		Margin(0, 2).
		BorderForeground(lipgloss.Color("12"))

	normal := lipgloss.NewStyle().
		Width(c.width).
		Height(1).
		Padding(0, 1)

	focus := lipgloss.NewStyle().
		Background(lipgloss.Color("12")).
		Width(c.width-4).Padding(0, 1).
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

	row := lipgloss.JoinVertical(lipgloss.Left, items...)

	return border.Render(row)
}
