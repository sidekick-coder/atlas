package form

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (c *Component) InitRender() tea.Cmd {
	return nil
}

func (c *Component) Render() string {
	var lines []string

	for index, i := range c.inputs {
		f := c.fields[index]

		content := i.Render()

		field := c.fieldBorder.SetLabel(f.Label).SetContent(content).Render()

		if c.selection.IsSelected(index) {
			field = c.fieldBorderSelected.SetLabel(f.Label).SetContent(content).Render()
		}

		lines = append(lines, field)
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}
