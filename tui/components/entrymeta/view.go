package entrymeta

import (
	"strings"

	lipgloss "charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (c *Component) SetSize(width, height int) *Component {
	c.Width = width
	c.Height = height

	return c
}

func (c *Component) Render() string {
	border := theme.BaseStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Width(c.Width-4).
		Height(c.Height).
		Margin(0, 2).
		BorderForeground(lipgloss.Color(theme.Current.Primary))

	if c.Focus {
		border = border.BorderForeground(lipgloss.Color("33"))
	}

	srs := theme.BaseStyle().
		Width(c.Width - 4).
		Background(lipgloss.Color(theme.Current.Primary)).
		Foreground(lipgloss.Color(theme.Current.Foreground))

	ks := theme.BaseStyle().
		Foreground(lipgloss.Color(theme.Current.Accent))

	vs := theme.BaseStyle().
		Foreground(lipgloss.Color(theme.Current.Foreground))

	var items []string

	for index, em := range c.Metas {
		name := em.Name + ":"
		value := em.Value
		value = strings.ReplaceAll(value, "\n", "\\n")

		if len(value) > 50 {
			value = value[:50] + "..."
		}

		pad := c.Width - 4 - len([]rune(name)) - len([]rune(value)) - 2

		pad = max(pad, 0)

		spaces := vs.Render(strings.Repeat(" ", pad))

		row := ks.Render(name) + spaces + vs.Render(value)

		if index == c.CurrentIndex {
			row = srs.Render(name + strings.Repeat(" ", pad) + value)
		}

		items = append(items, row)
	}

	content := lipgloss.JoinVertical(lipgloss.Left, items...)

	return border.Render(content)
}
