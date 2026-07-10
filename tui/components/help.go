package components

import (
	"fmt"

	lipgloss "charm.land/lipgloss/v2"
	"charm.land/bubbles/v2/key"
)

var (
	helpTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("12")).
			MarginBottom(1)

	helpKeyColStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("12")).
			Bold(true).
			Width(12)

	helpDescColStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("252"))

	helpRowStyle = lipgloss.NewStyle().
			Padding(0, 2)

	helpContainerStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("240")).
				Padding(1, 2)
)

// HelpScreen renders a full-screen overlay listing all key bindings.
type HelpScreen struct {
	width  int
	height int
	keymap KeyMap
}

func NewHelpScreen(km KeyMap) *HelpScreen {
	return &HelpScreen{keymap: km}
}

func (h *HelpScreen) SetSize(width, height int) {
	h.width = width
	h.height = height
}

func (h *HelpScreen) View() string {
	bindings := []key.Binding{
		h.keymap.Up,
		h.keymap.Down,
		h.keymap.FocusNext,
		h.keymap.MetaAdd,
		h.keymap.MetaReplace,
		h.keymap.MetaUpdate,
		h.keymap.MetaEditor,
		h.keymap.Help,
		h.keymap.Quit,
	}

	title := helpTitleStyle.Render("Keyboard Shortcuts")

	var rows []string
	for _, b := range bindings {
		hh := b.Help()
		row := helpRowStyle.Render(
			fmt.Sprintf("%s%s", helpKeyColStyle.Render(hh.Key), helpDescColStyle.Render(hh.Desc)),
		)
		rows = append(rows, row)
	}

	body := lipgloss.JoinVertical(lipgloss.Left, rows...)
	content := lipgloss.JoinVertical(lipgloss.Left, title, body)

	boxWidth := 40
	boxHeight := len(bindings) + 6
	if boxWidth > h.width-4 {
		boxWidth = h.width - 4
	}

	box := helpContainerStyle.Width(boxWidth).Height(boxHeight).Render(content)

	return box
}
