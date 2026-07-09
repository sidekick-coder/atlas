package list

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/messages"
)

func (c *Component) SetItems(items []string) {
	c.items = items

	maxIndex := len(items) - 1

	if c.cursor > maxIndex {
		c.cursor = maxIndex
	}
}


func (c *Component) Up() tea.Cmd {
	if c.cursor > 0 {
		c.cursor--
	}

	return messages.SkipCmd()
}

func (c *Component) Down() tea.Cmd {
	if c.cursor < len(c.items)-1 {
		c.cursor++
	}

	return messages.SkipCmd()
}

func (c *Component) GetCursor() (int, bool) {
	if c.cursor < 0 || c.cursor >= len(c.items) {
		return -1, false
	}

	return c.cursor, true
}

func (c *Component) HandleSelection(km tea.KeyMsg) tea.Cmd {
	if key.Matches(km, Binding.Up) {
		return c.Up()
	}

	if key.Matches(km, Binding.Down) {
		return c.Down()
	}
	if (key.Matches(km, Binding.Enter)) {
		index, ok := c.GetCursor()

		if ok && c.onSelect != nil {
			return c.onSelect(index)
		}

		return messages.SkipCmd()
	}

	return nil
}
