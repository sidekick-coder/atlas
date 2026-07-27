package keyvalue

import (
	"strings"

	lipgloss "charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (c *Component) SetSize(width, height int) *Component {
	c.width = width
	c.height = height

	c.viewport.SetSize(width, height)

	return c
}

func (c *Component) Render() string {
	srs := theme.BaseStyle().
		Width(c.width).
		Background(lipgloss.Color(theme.Current.Selection)).
		Foreground(lipgloss.Color(theme.Current.Foreground))

	ks := theme.BaseStyle().
		Foreground(lipgloss.Color(theme.Current.Accent))

	vs := theme.BaseStyle()

	hs := theme.BaseStyle().
		Foreground(lipgloss.Color(theme.Current.Primary))

	var items []string

	for index, em := range c.items {
		if em.Header {
			row := hs.Render(em.Key)

			if c.selection.IsSelected(index) {
				row = srs.Render(em.Key)
			}

			items = append(items, row)
			continue
		}

		name := em.Key + ": "
		value := em.Value
		value = strings.ReplaceAll(value, "\n", "\\n")

		vwidth := c.width - lipgloss.Width(name) - 20

		if len(value) >= vwidth {
			value = value[:vwidth-3] + "..."
		}

		row := ks.Render(name) + vs.Render(value)

		if c.selection.IsSelected(index) {
			row = srs.Render(name + value)
		}

		items = append(items, row)
	}

	content := lipgloss.JoinVertical(lipgloss.Left, items...)

	content = c.viewport.SetContent(content).Render()

	return content
}
