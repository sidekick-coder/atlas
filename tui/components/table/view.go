package table

import (
	"strings"

	lipgloss "charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (c *Component) SetSize(w, h int) {
	c.width = w
	c.height = h
}

func (c *Component) GetColumnIndex(column Column) int {
	for i, col := range c.columns {
		if col.Field == column.Field {
			return i
		}
	}

	return -1
}

func (c *Component) ParseColumnText(column Column,  text string) string {
	colIndex := c.GetColumnIndex(column)

	if colIndex == -1 {
		return text
	}

	colw := c.columnSizes[colIndex]

	var result string

	if len(text) > colw {
		result = text[:colw-3] + "..."
	}

	if len(text) < colw {
		padding := colw - len(text)
		result = text + strings.Repeat("\u00A0", padding)
	}

	return result
}

func (c *Component) Render() string {
	colstyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Current.Primary)).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(lipgloss.Color(theme.Current.Primary))

	colFocusStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Current.Primary)).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(lipgloss.Color(theme.Current.Primary)).
		Bold(true)

	itemFocusStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Current.Primary)).
		Bold(true)

	var rows []string
	var columns []string

	for index, column := range c.columns {
		label := c.ParseColumnText(column, column.Label)

		if c.columnSelection.IsSelected(index) {
			columns = append(columns, colFocusStyle.Render(label))
			continue
		}

		columns = append(columns, colstyle.Render(label))
	}

	rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Left, columns...))

	for index, item := range c.items {
		var row []string 

		for _, column := range c.columns {
			value, ok := item.Values[column.Field]

			if !ok {
				value = ""
			}

			value = c.ParseColumnText(column, value)

			row = append(row, value)
		}

		rendered := lipgloss.JoinHorizontal(lipgloss.Left, row...)
		
		if c.itemSelection.IsSelected(index) {
			rendered = itemFocusStyle.Render(rendered)
		}

		rows = append(rows, rendered)
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
