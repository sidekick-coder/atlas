package table

import (
	"log/slog"

	lipgloss "charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (c *Component) SetSize(w, h int) {
	slog.Info("table component size changed", slog.Int("width", w), slog.Int("height", h))
	c.width = w
	c.height = h
	c.column.SetSize(w, h)
}

func (c *Component) Render() string {
	colstyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Current.Muted)).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(lipgloss.Color(theme.Current.Muted))

	colFocusStyle := colstyle.
	    Foreground(lipgloss.Color(theme.Current.Primary)).
		BorderForeground(lipgloss.Color(theme.Current.Primary)).
		Bold(true)

	rowStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Current.Foreground))

	rowFocusStyle := lipgloss.NewStyle().
		Background(lipgloss.Color(theme.Current.Primary)).
		Bold(true)

	var rows []string
	var columns []string

	for index, column := range c.column.GetColumns() {
		label := c.column.ParseColumnText(column, column.Label)

		if c.column.Selection.IsSelected(index) {
			columns = append(columns, colFocusStyle.Render(label))
			continue
		}

		columns = append(columns, colstyle.Render(label))
	}

	rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Left, columns...))

	for index, item := range c.items {
		var row []string

		for _, column := range c.column.GetColumns() {
			value, ok := item.Values[column.Field]

			if !ok {
				value = ""
			}

			value = c.column.ParseColumnText(column, value)

			row = append(row, value)
		}

		rendered := rowStyle.Render(
			lipgloss.JoinHorizontal(lipgloss.Left, row...),
		)

		if c.itemSelection.IsSelected(index) {
			rendered = rowFocusStyle.Render(
				lipgloss.JoinHorizontal(lipgloss.Left, row...),
			)
		}

		rows = append(rows, rendered)
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
