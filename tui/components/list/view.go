package list

import (
	lipgloss "charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/features/theme"
	"github.com/charmbracelet/x/ansi"
)

func (c *Component) SetSize(w, h int) *Component {
	c.width = w
	c.height = h
	return c
}

func (c *Component) SetWidth(w int) *Component {
	c.width = w
	return c
}

func (c *Component) SetHeight(h int) *Component {
	c.height = h
	return c
}

func (c *Component) Render() string {
	normal := theme.BaseStyle().
	    Height(1).
		MaxHeight(1).
		Width(c.width)

	focus := theme.BaseStyle().
		Width(c.width).
		Height(1).
		MaxHeight(1).
		Background(lipgloss.Color(theme.Current.Selection))

	var items []string

	for index, item := range c.items {
		if c.selection.IsSelected(index) {
			unstyled := ansi.Strip(item)
			result := focus.Render(unstyled)

			items = append(items, result)
			continue
		}

		items = append(items, normal.Render(item))
	}

	if len(items) == 0 {
		return normal.Render("No items")
	}

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}
