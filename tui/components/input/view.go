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

	if !i.enabled {
		cursorStyle = text
	}

	// Keep cursor visible.
	if i.cursor < i.offset {
		i.offset = i.cursor
	}
	if i.cursor >= i.offset+i.width {
		i.offset = i.cursor - i.width + 1
	}

	start := i.offset
	end := min(len(i.buf), start+i.width)

	visible := i.buf[start:end]

	cursor := i.cursor - start

	before := ""
	if cursor > 0 {
		before = text.Render(string(visible[:cursor]))
	}

	var cur string
	if cursor >= 0 && cursor < len(visible) {
		cur = cursorStyle.Render(string(visible[cursor]))
	} else {
		cur = cursorStyle.Render(" ")
	}

	after := ""
	if cursor+1 < len(visible) {
		after = text.Render(string(visible[cursor+1:]))
	}

	rendered := before + cur + after

	// Pad so the rendered width is always i.width.
	if len(visible) < i.width {
		for j := 0; j < i.width-len(visible); j++ {
			rendered += text.Render(" ")
		}
	}

	return rendered
}
