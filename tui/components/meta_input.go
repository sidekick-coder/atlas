package components

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

// MetaInputSubmitMsg is emitted when the user confirms a value in the input overlay.
type MetaInputSubmitMsg struct {
	EntryID int64
	Name    string
	Value   string
}

// MetaOpenEditorMsg is emitted when the user requests to edit a value in $EDITOR.
type MetaOpenEditorMsg struct {
	EntryID      int64
	Name         string
	CurrentValue string
}

const inputBoxWidth = 58 // total outer width

var (
	inputBorderColor = lipgloss.Color("12")
	inputDimColor    = lipgloss.Color("240")

	inputBorderStyle = lipgloss.NewStyle().Foreground(inputBorderColor)
	inputDimStyle    = lipgloss.NewStyle().Foreground(inputDimColor)
	inputLabelStyle  = lipgloss.NewStyle().Bold(true).Foreground(inputBorderColor)
	inputTextStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	inputCursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("0")).Background(inputBorderColor)
)

// buildTopBorder returns  ╭─ label ──...──╮
func buildTopBorder(label string, width int) string {
	inner := width - 2 // subtract ╭ and ╮
	labelPart := "─ " + label + " "
	remaining := inner - len([]rune(labelPart))
	if remaining < 1 {
		remaining = 1
	}
	return inputBorderStyle.Render("╭"+labelPart+strings.Repeat("─", remaining)+"╮")
}

// buildBottomBorder returns  ╰──...── hint ─╯
func buildBottomBorder(hint string, width int) string {
	inner := width - 2
	hintPart := " " + hint + " ─"
	// use lipgloss.Width so ANSI codes in hint don't skew the count
	remaining := inner - lipgloss.Width(hintPart)
	if remaining < 1 {
		remaining = 1
	}
	return inputBorderStyle.Render("╰"+strings.Repeat("─", remaining)+hintPart+"╯")
}

// buildRow returns  │ content padded to width │
func buildRow(content string, innerWidth int) string {
	// pad content to innerWidth with spaces
	contentWidth := lipgloss.Width(content)
	pad := innerWidth - contentWidth
	if pad < 0 {
		pad = 0
	}
	return inputBorderStyle.Render("│") + " " + content + strings.Repeat(" ", pad) + " " + inputBorderStyle.Render("│")
}

// buildEmptyRow returns an empty padded row.
func buildEmptyRow(innerWidth int) string {
	return buildRow("", innerWidth)
}

// renderCursor builds the text line with a block cursor.
func renderCursor(buf []rune, cursor int) string {
	before := inputTextStyle.Render(string(buf[:cursor]))
	var cur string
	if cursor < len(buf) {
		cur = inputCursorStyle.Render(string(buf[cursor]))
	} else {
		cur = inputCursorStyle.Render(" ")
	}
	after := ""
	if cursor < len(buf) {
		after = inputTextStyle.Render(string(buf[cursor+1:]))
	}
	return before + cur + after
}

// ── MetaInput ────────────────────────────────────────────────────────────────

// MetaInput is a centered text input overlay for editing a meta value.
type MetaInput struct {
	entryID int64
	name    string
	buf     []rune
	cursor  int
	width   int
	height  int
	active  bool
}

func NewMetaInput() *MetaInput {
	return &MetaInput{}
}

func (m *MetaInput) SetSize(width, height int) {
	m.width = width
	m.height = height
}

func (m *MetaInput) Open(entryID int64, name, initialValue string) {
	m.entryID = entryID
	m.name = name
	m.buf = []rune(initialValue)
	m.cursor = len(m.buf)
	m.active = true
}

func (m *MetaInput) Close() {
	m.active = false
	m.buf = nil
	m.cursor = 0
}

func (m *MetaInput) Active() bool {
	return m.active
}

func (m *MetaInput) Update(msg tea.Msg) tea.Cmd {
	if !m.active {
		return nil
	}
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.Code {
		case tea.KeyEscape:
			m.Close()
		case tea.KeyEnter:
			value := string(m.buf)
			m.Close()
			return func() tea.Msg {
				return MetaInputSubmitMsg{EntryID: m.entryID, Name: m.name, Value: value}
			}
		case tea.KeyBackspace:
			if m.cursor > 0 {
				m.buf = append(m.buf[:m.cursor-1], m.buf[m.cursor:]...)
				m.cursor--
			}
		case tea.KeyDelete:
			if m.cursor < len(m.buf) {
				m.buf = append(m.buf[:m.cursor], m.buf[m.cursor+1:]...)
			}
		case tea.KeyLeft:
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.KeyRight:
			if m.cursor < len(m.buf) {
				m.cursor++
			}
		case tea.KeyHome:
			m.cursor = 0
		case tea.KeyEnd:
			m.cursor = len(m.buf)
		default:
			if msg.Text != "" {
				m.buf = append(m.buf[:m.cursor], append([]rune(msg.Text), m.buf[m.cursor:]...)...)
				m.cursor += len([]rune(msg.Text))
			}
		}
	}
	return nil
}

func (m *MetaInput) View() string {
	innerWidth := inputBoxWidth - 4 // subtract borders (2) and side spaces (2)

	top := buildTopBorder(m.name, inputBoxWidth)
	inputRow := buildRow(renderCursor(m.buf, m.cursor), innerWidth)
	bottom := buildBottomBorder(inputDimStyle.Render("enter")+" · "+inputDimStyle.Render("esc"), inputBoxWidth)

	box := lipgloss.JoinVertical(lipgloss.Left, top, inputRow, bottom)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
}

// ── MetaAddInput ─────────────────────────────────────────────────────────────

// MetaAddInput asks only for a meta name and submits with an empty value.
type MetaAddInput struct {
	entryID int64
	buf     []rune
	cursor  int
	width   int
	height  int
	active  bool
}

func NewMetaAddInput() *MetaAddInput {
	return &MetaAddInput{}
}

func (m *MetaAddInput) SetSize(width, height int) {
	m.width = width
	m.height = height
}

func (m *MetaAddInput) Open(entryID int64) {
	m.entryID = entryID
	m.buf = nil
	m.cursor = 0
	m.active = true
}

func (m *MetaAddInput) Close() {
	m.active = false
}

func (m *MetaAddInput) Active() bool {
	return m.active
}

func (m *MetaAddInput) Update(msg tea.Msg) tea.Cmd {
	if !m.active {
		return nil
	}
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.Code {
		case tea.KeyEscape:
			m.Close()
		case tea.KeyEnter:
			name := string(m.buf)
			m.Close()
			return func() tea.Msg {
				return MetaInputSubmitMsg{EntryID: m.entryID, Name: name, Value: ""}
			}
		case tea.KeyBackspace:
			if m.cursor > 0 {
				m.buf = append(m.buf[:m.cursor-1], m.buf[m.cursor:]...)
				m.cursor--
			}
		case tea.KeyDelete:
			if m.cursor < len(m.buf) {
				m.buf = append(m.buf[:m.cursor], m.buf[m.cursor+1:]...)
			}
		case tea.KeyLeft:
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.KeyRight:
			if m.cursor < len(m.buf) {
				m.cursor++
			}
		case tea.KeyHome:
			m.cursor = 0
		case tea.KeyEnd:
			m.cursor = len(m.buf)
		default:
			if msg.Text != "" {
				m.buf = append(m.buf[:m.cursor], append([]rune(msg.Text), m.buf[m.cursor:]...)...)
				m.cursor += len([]rune(msg.Text))
			}
		}
	}
	return nil
}

func (m *MetaAddInput) View() string {
	innerWidth := inputBoxWidth - 4

	hint := inputDimStyle.Render("enter") + " · " + inputDimStyle.Render("esc")
	top := buildTopBorder("Add Metadata", inputBoxWidth)
	inputRow := buildRow(renderCursor(m.buf, m.cursor), innerWidth)
	bottom := buildBottomBorder(hint, inputBoxWidth)

	box := lipgloss.JoinVertical(lipgloss.Left, top, inputRow, bottom)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
}

