package input

import (
	// "image/color"
	"strings"
	lipgloss "charm.land/lipgloss/v2"
)

const inputBoxWidth = 58

func (i *Input) Render() string {
	border := lipgloss.NewStyle().Foreground(i.Color)
	dim := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("0")).Background(i.Color)
	text := lipgloss.NewStyle().Foreground(lipgloss.Color("252"))

	innerWidth := inputBoxWidth - 4

	// Top border with title.
	labelPart := "─ " + i.Title + " "
	remaining := (inputBoxWidth - 2) - len([]rune(labelPart))
	if remaining < 1 {
		remaining = 1
	}

	top := border.Render("╭" + labelPart + strings.Repeat("─", remaining) + "╮")

	// Input row with block cursor.
	before := text.Render(string(i.Buf[:i.Cursor]))

	var cur string

	if i.Cursor < len(i.Buf) {
		cur = cursorStyle.Render(string(i.Buf[i.Cursor]))
	} else {
		cur = cursorStyle.Render(" ")
	}

	after := ""
	if i.Cursor < len(i.Buf) {
		after = text.Render(string(i.Buf[i.Cursor+1:]))
	}

	inputContent := before + cur + after
	contentW := lipgloss.Width(inputContent)
	pad := innerWidth - contentW

	if pad < 0 {
		pad = 0
	}

	row := border.Render("│") + " " + inputContent + strings.Repeat(" ", pad) + " " + border.Render("│")

	// Bottom border with hint.
	hint := dim.Render("enter") + " · " + dim.Render("esc")
	hintPart := " " + hint + " ─"
	remaining = (inputBoxWidth - 2) - lipgloss.Width(hintPart)

	if remaining < 1 {
		remaining = 1
	}

	bottom := border.Render("╰" + strings.Repeat("─", remaining) + hintPart + "╯")

	return lipgloss.JoinVertical(lipgloss.Left, top, row, bottom)
}

func (i *Input) RenderLayer() *lipgloss.Layer {
	return lipgloss.
		NewLayer(i.Render()).
		X(i.ScreenWidth/2 - i.Width/2).
		Y(i.ScreenHeight/2 - i.Height/2)
}
