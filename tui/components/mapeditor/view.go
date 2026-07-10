package mapeditor

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (c *Component) InitRender() tea.Cmd {
	c.dialog.
		OnRender(c.Render).
		SetWidth(50)

	return nil
}

func (c *Component) Render() string {
	var lines []string

	for index, f := range c.fields {
		content := ""

		value, ok := c.values[f.FielName]

		if ok {
			content = value
		}

		field := c.fieldBorder.SetLabel(f.Label).SetContent(content).Render()

		if c.selection.IsSelected(index) {
			field = c.fieldBorderSelected.SetLabel(f.Label).SetContent(content).Render()
		}

		lines = append(lines, field)
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}
