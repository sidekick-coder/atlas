package toast

import (
	// "image/color"
	lipgloss "charm.land/lipgloss/v2"
	"strings"
)

func (c *Component) Render() string {
	border := lipgloss.NewStyle().Foreground(c.Color)
	text := lipgloss.NewStyle().Foreground(lipgloss.Color("252"))

	boxWidth := lipgloss.Width(c.Content) + 4 // 2 for padding on each side

	// Top border with title.
	labelPart := "─ " + c.Title + " "

	remaining := (boxWidth) - len([]rune(labelPart))

	remaining = max(remaining, 1)

	topLen := boxWidth - lipgloss.Width(labelPart) - 2 // 2 for the corners

	top := border.Render("╭" + labelPart + strings.Repeat("─", topLen) + "╮")

	inputContent := text.Render(c.Content)

	pad := boxWidth - lipgloss.Width(inputContent) - 4

	pad = max(pad, 0)

	row := border.Render("│") + " " + inputContent + strings.Repeat(".", pad) + " " + border.Render("│")

	bottomLen := lipgloss.Width(inputContent) + 2 // 2 for the corners
	bottom := border.Render("╰" + strings.Repeat("─", bottomLen) + "╯")

	return lipgloss.JoinVertical(lipgloss.Left, top, row, bottom)
}

func (i *Component) RenderLayer() *lipgloss.Layer {
	return lipgloss.
		NewLayer(i.Render()).
		X(i.ScreenWidth/2 - i.Width/2).
		Y(i.ScreenHeight/2 - i.Height/2)
}
