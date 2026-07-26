package keyvalue

import (
	"strings"

	lipgloss "charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (c *Component) SetSize(width, height int) *Component {
	c.width = width
	c.height = height

	return c
}

func (c *Component) Render() string {
	srs := theme.BaseStyle().
		Width(c.width).
		Background(lipgloss.Color(theme.Current.Selection))

	ks := theme.BaseStyle().
		Foreground(lipgloss.Color(theme.Current.Accent))

	vs := theme.BaseStyle().
		Foreground(lipgloss.Color(theme.Current.Foreground))

	var items []string

	for index, em := range c.items {
		name := em.Key + ":"
		value := em.Value
		value = strings.ReplaceAll(value, "\n", "\\n")

		if len(value) > 50 {
			value = value[:50] + "..."
		}

		pad := c.width - len([]rune(name)) - len([]rune(value)) - 2

		pad = max(pad, 0)

		spaces := vs.Render(strings.Repeat(" ", pad))

		row := ks.Render(name) + spaces + vs.Render(value)

		if c.selection.IsSelected(index) {
			row = srs.Render(name + strings.Repeat(" ", pad) + value)
		}

		items = append(items, row)
	}

	content := lipgloss.JoinVertical(lipgloss.Left, items...)

	return content
}
