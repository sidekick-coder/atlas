package borderlabel

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (c *Component) Render() string {
	border := theme.BaseStyle().Foreground(lipgloss.Color(c.color))
	text := theme.BaseStyle()
	maxWidth := c.width + 4

	boxWidth := c.width + 4                                 // 2 for padding on each side
	boxWidth = max(boxWidth, lipgloss.Width(c.label)+6, 50) // 2 for corners, 4 for padding
	boxWidth = min(boxWidth, maxWidth)

	// Top border with title.
	labelPart := "─ " + c.label + " "

	if c.label == "" {
		labelPart = ""
	}

	empty := text.Render(" ")

	topLen := boxWidth - lipgloss.Width(labelPart) - 2 // 2 for the corners
	topLen = max(topLen, 0)
	top := border.Render("╭" + labelPart + strings.Repeat("─", topLen) + "╮")

	inputContent := text.Render(c.content)

	lines := strings.Split(inputContent, "\n")

	availableHeight := c.height - 2

	if len(lines) < availableHeight {
		for i := len(lines); i < availableHeight; i++ {
			lines = append(lines, "")
		}
	}

	rowParts := make([]string, 0, len(lines))

	for _, line := range lines {
		pad := max(0, boxWidth - lipgloss.Width(line) - 4) // 4 for the corners and padding

		part := border.Render("│") + empty + line + strings.Repeat(empty, pad) + empty + border.Render("│")

		rowParts = append(rowParts, part)
	}

	row := strings.Join(rowParts, "\n")

	bottomLen := max(boxWidth-2, 0) // 2 for the corners
	bottom := border.Render("╰" + strings.Repeat("─", bottomLen) + "╯")

	return lipgloss.JoinVertical(lipgloss.Left, top, row, bottom)
}
