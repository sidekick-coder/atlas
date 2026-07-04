package components

import (
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

// ToastExpiredMsg is emitted by the toast timer — handled by the root model.
type ToastExpiredMsg struct{}

// ToastShowMsg is sent by any component to request a toast from the root model.
type ToastShowMsg struct {
	Message string
	Level   toastLevel
}

// BatchMsg carries multiple messages to be dispatched sequentially by the root model.
type BatchMsg struct {
	Msgs []tea.Msg
}

type toastLevel int

const (
	ToastInfo toastLevel = iota
	ToastSuccess
	ToastError
)

// Toast is a self-dismissing notification box rendered as a floating overlay.
// Use GlobalToast from anywhere; the root model handles rendering and expiry.
type Toast struct {
	message string
	level   toastLevel
	active  bool
	width   int
	height  int
}

var GlobalToast = &Toast{}

func (t *Toast) SetSize(width, height int) {
	t.width = width
	t.height = height
}

// Show activates the toast and returns a Cmd that fires ToastExpiredMsg after d.
func (t *Toast) Show(msg string, level toastLevel, d time.Duration) tea.Cmd {
	t.message = msg
	t.level = level
	t.active = true
	return func() tea.Msg {
		time.Sleep(d)
		return ToastExpiredMsg{}
	}
}

func (t *Toast) Hide() {
	t.active = false
}

func (t *Toast) Active() bool { return t.active }

// Box returns the raw rendered box for compositing with PlaceOverlay.
func (t *Toast) Box() string {
	borderColor := lipgloss.Color("33") // blue (info default)
	icon := " "
	switch t.level {
	case ToastSuccess:
		borderColor = lipgloss.Color("40") // green
		icon = "✓ "
	case ToastError:
		borderColor = lipgloss.Color("196") // red
		icon = "✗ "
	}

	border := lipgloss.NewStyle().Foreground(borderColor)
	msgStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("252"))

	text := icon + t.message
	innerWidth := len([]rune(text)) + 2 // padding on each side
	if innerWidth < 20 {
		innerWidth = 20
	}

	top := border.Render("╭" + strings.Repeat("─", innerWidth) + "╮")
	row := border.Render("│") + " " + msgStyle.Render(text) + " " + border.Render("│")
	bottom := border.Render("╰" + strings.Repeat("─", innerWidth) + "╯")

	return lipgloss.JoinVertical(lipgloss.Left, top, row, bottom)
}

// View renders the box centered on the terminal.
func (t *Toast) View() string {
	return centerBox(t.Box(), t.width, t.height)
}
