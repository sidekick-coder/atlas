package entrylist

import (
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/internal/utils/maputil"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (c *Component) SetSize(width, height int) {
	c.list.SetSize(width, height)

	limit := 20

	limit = max(limit, height)

	c.loader.SetLimit(limit)

	c.loader.Load()
}

func (c *Component) Render() string {
	var items []string

	labelKey := "path"

	if l, ok := c.props["label_key"].(string); ok {
		labelKey = l
	}

	muted := theme.BaseStyle().Foreground(lipgloss.Color(theme.Current.Muted))

	for index, entry := range c.loader.GetEntries() {
		em := entry.ToMap()

		value := muted.Render("N/A")

		if c.selection.IsSelected(index) {
			value = muted.Background(lipgloss.Color(theme.Current.Selection)).Render("N/A")
		}

		v, ok := maputil.GetString(em, labelKey)

		if ok {
			value = v
		}

		items = append(items, value)
	}

	c.list.SetItems(items)

	return c.list.Render()
}
