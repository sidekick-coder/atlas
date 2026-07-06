package toast

import (
	// "image/color"
	"strings"
	lipgloss "charm.land/lipgloss/v2"
)

const BoxWidth = 58

func (c *Component) Render() string {
	border := lipgloss.NewStyle().Foreground(c.Color)
	text := lipgloss.NewStyle().Foreground(lipgloss.Color("252"))

	innerWidth := BoxWidth - 4

	// Top border with title.
	labelPart := "─ " + c.Title + " "

	remaining := (BoxWidth - 2) - len([]rune(labelPart))

	if remaining < 1 {
		remaining = 1
	}

	top := border.Render("╭" + labelPart + strings.Repeat("─", remaining) + "╮")

	inputContent := text.Render(c.Content)

	contentW := lipgloss.Width(inputContent)
	pad := innerWidth - contentW

	if pad < 0 {
		pad = 0
	}

	row := border.Render("│") + " " + inputContent + strings.Repeat(" ", pad) + " " + border.Render("│")

	// Bottom border with hint.
	if remaining < 1 {
		remaining = 1
	}

	bottom := border.Render("╰" + strings.Repeat("─", BoxWidth - 2) + "╯")

	return lipgloss.JoinVertical(lipgloss.Left, top, row, bottom)
}

func (i *Component) RenderLayer() *lipgloss.Layer {
	return lipgloss.
		NewLayer(i.Render()).
		X(i.ScreenWidth/2 - i.Width/2).
		Y(i.ScreenHeight/2 - i.Height/2)
}
