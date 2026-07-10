package input

import (
	lipgloss "charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (i *Input) Render() string {
	cursorStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Current.Foreground)).
		Background(lipgloss.Color(theme.Current.Primary))

	text := lipgloss.NewStyle().Foreground(lipgloss.Color(theme.Current.Foreground))

	if (!i.enabled) {
		cursorStyle = text
	}

	// Input row with block cursor.
	before := text.Render(string(i.buf[:i.cursor]))

	var cur string

	if i.cursor < len(i.buf) {
		cur = cursorStyle.Render(string(i.buf[i.cursor]))
	} else {
		cur = cursorStyle.Render(" ")
	}

	after := ""

	if i.cursor < len(i.buf) {
		after = text.Render(string(i.buf[i.cursor+1:]))
	}

	return before + cur + after
}
