package components

import (
	"image/color"
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

const inputBoxWidth = 58

// GlobalInput is the single shared floating text input for the entire TUI.
// Any component can open it; the root model handles rendering the overlay.
var GlobalInput = NewInput()

// Input is a generic single-line text input that renders as a floating dialog box.
// Configure it with the fluent setters, call Open() to activate, and read results
// via the OnSubmit callback — no screen coupling required.
type Input struct {
	title    string
	color    color.Color
	buf      []rune
	cursor   int
	width    int
	height   int
	active   bool
	onSubmit func(value string) tea.Cmd
}

func NewInput() *Input {
	return &Input{color: lipgloss.Color("12")}
}

func (i *Input) SetTitle(title string) *Input {
	i.title = title
	return i
}

func (i *Input) SetColor(c color.Color) *Input {
	i.color = c
	return i
}

func (i *Input) SetOnSubmit(fn func(string) tea.Cmd) *Input {
	i.onSubmit = fn
	return i
}

func (i *Input) SetSize(width, height int) {
	i.width = width
	i.height = height
}

func (i *Input) Open(initialValue string) {
	i.buf = []rune(initialValue)
	i.cursor = len(i.buf)
	i.active = true
}

func (i *Input) Close() {
	i.active = false
	i.buf = nil
	i.cursor = 0
}

func (i *Input) Active() bool { return i.active }

func (i *Input) Update(msg tea.Msg) tea.Cmd {
	if !i.active {
		return nil
	}
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.Code {
		case tea.KeyEscape:
			i.Close()
		case tea.KeyEnter:
			value := string(i.buf)
			i.Close()
			if i.onSubmit != nil {
				return i.onSubmit(value)
			}
		case tea.KeyBackspace:
			if i.cursor > 0 {
				i.buf = append(i.buf[:i.cursor-1], i.buf[i.cursor:]...)
				i.cursor--
			}
		case tea.KeyDelete:
			if i.cursor < len(i.buf) {
				i.buf = append(i.buf[:i.cursor], i.buf[i.cursor+1:]...)
			}
		case tea.KeyLeft:
			if i.cursor > 0 {
				i.cursor--
			}
		case tea.KeyRight:
			if i.cursor < len(i.buf) {
				i.cursor++
			}
		case tea.KeyHome:
			i.cursor = 0
		case tea.KeyEnd:
			i.cursor = len(i.buf)
		default:
			if msg.Text != "" {
				i.buf = append(i.buf[:i.cursor], append([]rune(msg.Text), i.buf[i.cursor:]...)...)
				i.cursor += len([]rune(msg.Text))
			}
		}
	}
	return nil
}

// Box returns the raw dialog box string (no surrounding whitespace).
// Use this with PlaceOverlay to composite over a background.
func (i *Input) Box() string {
	border := lipgloss.NewStyle().Foreground(i.color)
	dim := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("0")).Background(i.color)
	text := lipgloss.NewStyle().Foreground(lipgloss.Color("252"))

	innerWidth := inputBoxWidth - 4

	// Top border with title.
	labelPart := "─ " + i.title + " "
	remaining := (inputBoxWidth - 2) - len([]rune(labelPart))
	if remaining < 1 {
		remaining = 1
	}
	top := border.Render("╭" + labelPart + strings.Repeat("─", remaining) + "╮")

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

// View renders the box centered on the terminal with no background fill.
func (i *Input) View() string {
	return centerBox(i.Box(), i.width, i.height)
}

// centerBox positions a pre-rendered box in the center of a w×h terminal
// using leading newlines + per-line left padding, avoiding any background fill.
func centerBox(box string, w, h int) string {
	lines := strings.Split(box, "\n")
	boxH := len(lines)
	boxW := 0
	for _, l := range lines {
		if lw := lipgloss.Width(l); lw > boxW {
			boxW = lw
		}
	}
	topPad := (h - boxH) / 2
	if topPad < 0 {
		topPad = 0
	}
	leftPad := (w - boxW) / 2
	if leftPad < 0 {
		leftPad = 0
	}
	pad := strings.Repeat(" ", leftPad)
	var sb strings.Builder
	for i := 0; i < topPad; i++ {
		sb.WriteByte('\n')
	}
	for idx, l := range lines {
		sb.WriteString(pad + l)
		if idx < len(lines)-1 {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}
