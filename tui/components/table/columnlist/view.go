package columnlist

import (
	"strconv"

	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/components/borderlabel"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (c *Component) Render() string {
	var columns []string

	w, h := c.sidepeeck.GetSize()

	w -= 6 // padding

	border := borderlabel.Create().SetSize(w, h)

	for index, column := range c.column.GetColumns() {
		id := column.Field
		field := column.Field
		width := ""

		if column.Width > 0 {
			width = strconv.Itoa(column.Width)
		}

		if column.Width <= 0 {
			width = "auto"
		}

		parts := []string{
			"ID: " + id,
			"Field: " + field,
			"Width: " + width,
		}

		content := lipgloss.JoinVertical(lipgloss.Left, parts...)

		border.SetColor(theme.Current.Muted)

		if c.column.Selection.IsSelected(index) {
			border.SetColor(theme.Current.Primary)
		}

		c := border.SetLabel(column.Label).SetContent(content).Render()

		columns = append(columns, c)
	}

	return lipgloss.JoinVertical(lipgloss.Left, columns...)
}
