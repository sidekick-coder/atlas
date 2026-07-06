package entrymeta

import (
	"strings"

	lipgloss "charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/internal/models"
)

type Component struct {
	Width        int
	Height       int
	Focus        bool
	CurrentIndex int
	Metas        []models.EntryMeta
}

func Create() *Component {
	return &Component{
		Width:        100,
		Height:       100,
		CurrentIndex: 0,
		Focus:        false,
		Metas:        []models.EntryMeta{},
	}
}

func (c *Component) SetSize(width, height int) *Component {
	c.Width = width
	c.Height = height

	return c
}

func (c *Component) SetFocus(focus bool) *Component {
	c.Focus = focus
	return c
}

func (c *Component) SetMetas(metas []models.EntryMeta) {
	c.Metas = metas

	maxIndex := len(metas) - 1

	if c.CurrentIndex > maxIndex {
		c.CurrentIndex = maxIndex
	}
}

func (c *Component) Render() string {
	border := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Width(c.Width-4).
		Height(c.Height-4).
		Margin(0, 2).
		BorderForeground(lipgloss.Color("12"))

	if c.Focus {
		border = border.BorderForeground(lipgloss.Color("33"))
	}

	srs := lipgloss.NewStyle().
		Width(c.Width - 4).
		Background(lipgloss.Color("12")).
		Foreground(lipgloss.Color("0"))

	ks := lipgloss.NewStyle().
		Foreground(lipgloss.Color("12"))

	vs := lipgloss.NewStyle().
		Foreground(lipgloss.Color("252"))


	var items []string

	for index, em := range c.Metas {
		name := em.Name+":"
		value := em.Value
		value = strings.ReplaceAll(value, "\n", "\\n")

		if len(value) > 50 {
			value = value[:50] + "..."
		}

		pad := c.Width - 4 - len([]rune(name)) - len([]rune(value)) - 2

		if pad < 0 {
			pad = 0
		}

		row := ks.Render(name) + strings.Repeat(" ", pad) + vs.Render(value)

		if index == c.CurrentIndex {
			row = srs.Render(name + strings.Repeat(" ", pad) + value)
		}

		items = append(items, row)
	}

	content := lipgloss.JoinVertical(lipgloss.Left, items...)

	return border.Render(content)
}

func (c *Component) MoveUp() {
	if c.CurrentIndex > 0 {
		c.CurrentIndex--
	}
}

func (c *Component) MoveDown() {
	if c.CurrentIndex < len(c.Metas)-1 {
		c.CurrentIndex++
	}
}

func (c *Component) GetSelected() (models.EntryMeta, bool) {
	if c.CurrentIndex < 0 || c.CurrentIndex >= len(c.Metas) {
		return models.EntryMeta{}, false
	}

	return c.Metas[c.CurrentIndex], true
}
