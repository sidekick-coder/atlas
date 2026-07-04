package components

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/x/ansi"
	"github.com/sidekick-coder/atlas/internal/models"
)

var (
	metaRowStyle = lipgloss.NewStyle().
			Padding(0, 1)

	metaRowSelectedStyle = lipgloss.NewStyle().
				Padding(0, 1).
				Bold(true).
				Foreground(lipgloss.Color("12"))

	metasContainerStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("240"))
)

type EntryMetas struct {
	metas    []models.EntryMeta
	entryID  int64
	cursor   int
	focused  bool
	input    *MetaInput
	addInput *MetaAddInput
	width    int
	height   int
}

func NewEntryMetas() *EntryMetas {
	return &EntryMetas{
		input:    NewMetaInput(),
		addInput: NewMetaAddInput(),
	}
}

func (c *EntryMetas) SetEntryID(id int64) {
	c.entryID = id
}

func (c *EntryMetas) SetMetas(metas []models.EntryMeta) {
	// Preserve cursor position by keeping the same meta name selected.
	var selectedName string
	if c.cursor < len(c.metas) {
		selectedName = c.metas[c.cursor].Name
	}

	c.metas = metas

	// Try to find the previously selected meta by name.
	c.cursor = 0
	if selectedName != "" {
		for i, m := range metas {
			if m.Name == selectedName {
				c.cursor = i
				break
			}
		}
	}
}

func (c *EntryMetas) SetFocused(focused bool) {
	c.focused = focused
}

func (c *EntryMetas) SetSize(width, height int) {
	c.width = width
	c.height = height
}

// SetOverlaySize passes the full terminal dimensions to the input overlays.
func (c *EntryMetas) SetOverlaySize(width, height int) {
	c.input.SetSize(width, height)
	c.addInput.SetSize(width, height)
}

func (c *EntryMetas) InputActive() bool {
	return c.input.Active() || c.addInput.Active()
}

func (c *EntryMetas) Update(msg tea.Msg) tea.Cmd {
	// While input overlay is open, route everything to it.
	if c.input.Active() {
		return c.input.Update(msg)
	}
	if c.addInput.Active() {
		return c.addInput.Update(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Up):
			if c.cursor > 0 {
				c.cursor--
			}
		case key.Matches(msg, DefaultKeyMap.Down):
			if c.cursor < len(c.metas)-1 {
				c.cursor++
			}
		case key.Matches(msg, DefaultKeyMap.MetaAdd):
			c.addInput.Open(c.entryID)
		case key.Matches(msg, DefaultKeyMap.MetaReplace):
			if len(c.metas) > 0 {
				m := c.metas[c.cursor]
				c.input.Open(m.EntryID, m.Name, "")
			}
		case key.Matches(msg, DefaultKeyMap.MetaUpdate):
			if len(c.metas) > 0 {
				m := c.metas[c.cursor]
				c.input.Open(m.EntryID, m.Name, m.Value)
			}
		case key.Matches(msg, DefaultKeyMap.MetaEditor):
			if len(c.metas) > 0 {
				m := c.metas[c.cursor]
				return func() tea.Msg {
					return MetaOpenEditorMsg{
						EntryID:      m.EntryID,
						Name:         m.Name,
						CurrentValue: m.Value,
					}
				}
			}
		}
	}
	return nil
}

// sanitizeValue collapses newlines/tabs into spaces for single-line display.
func sanitizeValue(v string) string {
	v = strings.ReplaceAll(v, "\r\n", " ")
	v = strings.ReplaceAll(v, "\n", " ")
	v = strings.ReplaceAll(v, "\t", " ")
	return strings.TrimSpace(v)
}

func (c *EntryMetas) ActiveOverlay() string {
	if c.input.Active() {
		return c.input.View()
	}
	if c.addInput.Active() {
		return c.addInput.View()
	}
	return ""
}

func (c *EntryMetas) View() string {
	// 2 for border, 2 for padding (left+right from row style)
	innerWidth := c.width - 2
	if innerWidth < 4 {
		innerWidth = 4
	}
	rowContentWidth := innerWidth - 2 // subtract row padding

	var rows []string
	for i, meta := range c.metas {
		selected := i == c.cursor && c.focused

		nameStr := meta.Name + ":"
		valueStr := sanitizeValue(meta.Value)

		nameLen := len(nameStr) + 1 // name + space separator
		maxValueLen := rowContentWidth - nameLen
		if maxValueLen < 1 {
			maxValueLen = 1
		}

		if ansi.StringWidth(valueStr) > maxValueLen {
			valueStr = ansi.Truncate(valueStr, maxValueLen-1, "…")
		}

		line := nameStr + " " + valueStr
		if selected {
			rows = append(rows, metaRowSelectedStyle.Width(innerWidth).Render("▶ "+line))
		} else {
			rows = append(rows, metaRowStyle.Width(innerWidth).Render("  "+line))
		}
	}

	if len(rows) == 0 {
		rows = append(rows, metaRowStyle.Width(innerWidth).Render("No metadata"))
	}

	borderColor := lipgloss.Color("240")
	if c.focused {
		borderColor = lipgloss.Color("12")
	}

	content := lipgloss.JoinVertical(lipgloss.Left, rows...)
	return metasContainerStyle.
		BorderForeground(borderColor).
		Width(c.width - 2).
		Height(c.height - 2).
		Render(content)
}


