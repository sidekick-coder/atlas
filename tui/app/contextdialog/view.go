package contextdialog

import (
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (f *Component) render() string {
	content := f.kv.Render()

	w, h := f.dialog.GetContentSize()

	return theme.BaseStyle().
		Border(lipgloss.NormalBorder()).
		Width(w).
		Height(h).
		Render(content)
}
