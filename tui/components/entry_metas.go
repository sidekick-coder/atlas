package components

import (
	lipgloss "charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/internal/models"
)

var (
	metaNameStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("12"))

	metaValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252"))

	metaRowStyle = lipgloss.NewStyle().
			Padding(0, 1)

	metasContainerStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("240"))
)

type EntryMetas struct {
	metas  []models.EntryMeta
	width  int
	height int
}

func NewEntryMetas() *EntryMetas {
	return &EntryMetas{}
}

func (c *EntryMetas) SetMetas(metas []models.EntryMeta) {
	c.metas = metas
}

func (c *EntryMetas) SetSize(width, height int) {
	c.width = width
	c.height = height
}

func (c *EntryMetas) View() string {
	innerWidth := c.width - 2
	if innerWidth < 1 {
		innerWidth = 1
	}

	var rows []string
	for _, meta := range c.metas {
		name := metaNameStyle.Render(meta.Name+":")
		value := metaValueStyle.Render(meta.Value)
		rows = append(rows, metaRowStyle.Width(innerWidth).Render(name+" "+value))
	}

	if len(rows) == 0 {
		rows = append(rows, metaRowStyle.Width(innerWidth).Render("No metadata"))
	}

	content := lipgloss.JoinVertical(lipgloss.Left, rows...)
	return metasContainerStyle.Width(c.width - 2).Height(c.height - 2).Render(content)
}
