package components

import (
	"fmt"
	"strings"

	lipgloss "charm.land/lipgloss/v2"
	"charm.land/bubbles/v2/key"
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
	width    int
	bindings []key.Binding
}

func NewFooter() *Footer {
	return &Footer{}
}

func (f *Footer) SetWidth(width int) {
	f.width = width
}

func (f *Footer) SetBindings(bindings ...key.Binding) {
	f.bindings = bindings
}

func (f *Footer) Render() string {
	container := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Width(f.width - 4).
		Margin(0, 2).
		Padding(0, 2).
		BorderForeground(lipgloss.Color("12"))

	var parts []string

	for _, b := range f.bindings {
		h := b.Help()
		if h.Key == "" {
			continue
		}
		parts = append(parts, fmt.Sprintf("%s %s", footerKeyStyle.Render(h.Key), footerStyle.Render(h.Desc)))
	}

	sep := footerStyle.Render("  ·  ")
	row := strings.Join(parts, sep)

	return container.Render(row)
}

func (f *Footer) View() string {
	return f.Render()
}
