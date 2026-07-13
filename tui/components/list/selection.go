package list

import (
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

