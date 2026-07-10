package components

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/key"
	lipgloss "charm.land/lipgloss/v2"
	tkey "github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

var (
	footerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Current.Muted))

	footerKeyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Current.Primary)).
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
		Width(f.width-4).
		Margin(0, 2).
		Padding(0, 2).
		BorderForeground(lipgloss.Color(theme.Current.Primary))

	var parts []string

	for _, b := range f.bindings {
		h := b.Help()
		if h.Key == "" {
			continue
		}
		parts = append(parts, fmt.Sprintf("%s %s", footerKeyStyle.Render(h.Key), footerStyle.Render(h.Desc)))
	}

	for _, b := range tkey.GetBindings() {
		k := b.GetHelp()
		d := b.GetDescription()

		if k == "" || d == "" {
			continue
		}

		parts = append(parts, fmt.Sprintf("%s %s", footerKeyStyle.Render(k), footerStyle.Render(d)))
	}

	maxParts := 12

	if len(parts) > maxParts {
		remaning := len(parts) - maxParts
		parts = parts[:maxParts]
		parts = append(parts, footerStyle.Render(fmt.Sprintf("... and %d more", remaning)))
	}

	sep := footerStyle.Render(" · ")
	row := strings.Join(parts, sep)

	return container.Render(row)
}

func (f *Footer) View() string {
	return f.Render()
}

func (f *Footer) GetHeight() int {
	return lipgloss.Height(f.Render())
}

func (f *Footer) GetWidth() int {
	return lipgloss.Width(f.Render())
}
