package table

import (
	"strings"

	lipgloss "charm.land/lipgloss/v2"
)

func (c *Component) SetSize(w, h int) {
	c.width = w
	c.height = h
}

func (c *Component) Render() string {
	colw := c.width / len(c.columns)

	var colstyle = lipgloss.NewStyle().
		Width(colw).
		Background(lipgloss.Color("12")).
		Foreground(lipgloss.Color("0"))

	normal := lipgloss.NewStyle()

	focus := lipgloss.NewStyle().
		Background(lipgloss.Color("12")).
		Width(c.width).
		Foreground(lipgloss.Color("0"))

	var rows []string
	var columns []string

	for _, column := range c.columns {
		label := column.Label

		if len(label) > colw {
			label = label[:colw-3] + "..."
		}

		if len(label) < colw {
			padding := colw - len(label)
			label = label + strings.Repeat("\u00A0", padding)
		}

		columns = append(columns, colstyle.Render(label))
	}

	rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Left, columns...))

	for index, item := range c.items {
		if index == c.cursor {
			result := focus.Render(item)

			rows = append(rows, result)
			continue
		}

		rows = append(rows, normal.Render(item))
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
