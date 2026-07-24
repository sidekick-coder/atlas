package footer

import (
	"fmt"
	"strings"

	lipgloss "charm.land/lipgloss/v2"
	tkey "github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (f *Component) Render() string {
	bg := lipgloss.Color(theme.Current.Background)

	container := lipgloss.NewStyle().
		Width(f.width).
		Background(bg).
		Height(1).
		AlignVertical(lipgloss.Center)

	keyStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Current.Primary)).
		Background(bg).
		PaddingRight(1).
		Bold(true)

	textStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Current.Foreground)).
		Background(bg)

	parts := []string{}

	// icon
	remaningWidth := f.width

	for _, b := range tkey.GetBindings() {
		k := b.GetHelp()
		d := b.GetDescription()

		if f.dialog.IsOpen() && !b.HasTag("global:help") {
			continue
		}

		if k == "" || d == "" || b.IsHidden() {
			continue
		}

		key := keyStyle.Render(k)
		desc := textStyle.Render(d)

		if len(parts) == 0 {
			key = keyStyle.PaddingLeft(1).Render(k)
		}

		part := fmt.Sprintf("%s%s", key, desc)

		remaningWidth -= lipgloss.Width(part)

		parts = append(parts, part)

		if remaningWidth <= 80 {
			parts = append(parts, textStyle.Render(fmt.Sprintf("... and %d more", len(tkey.GetBindings())-len(parts))))
			break
		}

	}

	row := lipgloss.NewStyle().
		Background(lipgloss.Color(theme.Current.Primary)).
		Foreground(lipgloss.Color(theme.Current.Background)).
		Align(lipgloss.Center).
		Padding(0, 1).
		Render(strings.Join([]string{"󰆧", f.label}, " "))

	sep := textStyle.Render(" · ")

	row += strings.Join(parts, sep)

	return container.Render(row)
}
