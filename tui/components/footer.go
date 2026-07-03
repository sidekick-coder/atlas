package components

import (
	"fmt"

	lipgloss "charm.land/lipgloss/v2"
	"github.com/charmbracelet/bubbles/key"
)

var (
	footerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))

	footerKeyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("12")).
			Bold(true)
)

// Footer renders a one-line shortcuts bar.
type Footer struct {
	width  int
	keymap KeyMap
}

func NewFooter(km KeyMap) *Footer {
	return &Footer{keymap: km}
}

func (f *Footer) SetWidth(width int) {
	f.width = width
}

func (f *Footer) View() string {
	bindings := []key.Binding{
		f.keymap.Up,
		f.keymap.Down,
		f.keymap.Help,
		f.keymap.Quit,
	}

	var parts []string
	for _, b := range bindings {
		h := b.Help()
		parts = append(parts, fmt.Sprintf("%s %s", footerKeyStyle.Render(h.Key), footerStyle.Render(h.Desc)))
	}

	sep := footerStyle.Render("  ·  ")
	row := ""
	for i, p := range parts {
		if i > 0 {
			row += sep
		}
		row += p
	}

	return footerStyle.Width(f.width).Render(row)
}
