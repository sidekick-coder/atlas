package command

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (c *Component) InitView() tea.Cmd {
	w, h := c.dialog.GetContentSize()
	c.textfield.Resize(w, 1)
	c.container.SetSize(w, h)
	return nil
}

func (c *Component) Render() string {
	var items []string

	for _, cmd := range c.commands {
		items = append(items, cmd.Render())
	}

	items = append([]string{c.textfield.Render()}, items...)

	if len(items) == 1 {
		items = append(items, theme.BaseStyle().Render("No commands"))
	}

	content := lipgloss.JoinVertical(lipgloss.Left, items...)

	return c.container.SetContent(content).Render()
}
