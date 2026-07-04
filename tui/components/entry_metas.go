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
			Padding(0, 1).
			Foreground(lipgloss.Color("244"))

	metaRowNameStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("75")) // soft blue for key

	metaRowSelectedStyle = lipgloss.NewStyle().
				Padding(0, 1).
				Bold(true).
				Foreground(lipgloss.Color("12"))

	metasContainerStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder())

	metasTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("12")).
			Padding(0, 1)
)

type EntryMetas struct {
	metas   []models.EntryMeta
	entryID int64
	cursor  int
	focused bool
	width   int
	height  int
}

type MetaInputSubmitMsg struct {
	EntryID int64
	Name    string
	Value   string
}

// MetaOpenEditorMsg is emitted when the user requests to open $EDITOR.
type MetaOpenEditorMsg struct {
	EntryID      int64
	Name         string
	CurrentValue string
}

func NewEntryMetas() *EntryMetas {
	return &EntryMetas{}
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


func (c *EntryMetas) Update(msg tea.Msg) tea.Cmd {
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
			GlobalInput.
				SetTitle("Add Metadata").
				SetOnSubmit(func(name string) tea.Cmd {
					entryID := c.entryID
					return func() tea.Msg {
						return MetaInputSubmitMsg{EntryID: entryID, Name: name, Value: ""}
					}
				}).
				Open("")
		case key.Matches(msg, DefaultKeyMap.MetaReplace):
			if len(c.metas) > 0 {
				m := c.metas[c.cursor]
				GlobalInput.
					SetTitle(m.Name).
					SetOnSubmit(func(value string) tea.Cmd {
						return func() tea.Msg {
							return MetaInputSubmitMsg{EntryID: m.EntryID, Name: m.Name, Value: value}
						}
					}).
					Open("")
			}
		case key.Matches(msg, DefaultKeyMap.MetaUpdate):
			if len(c.metas) > 0 {
				m := c.metas[c.cursor]
				GlobalInput.
					SetTitle(m.Name).
					SetOnSubmit(func(value string) tea.Cmd {
						return func() tea.Msg {
							return MetaInputSubmitMsg{EntryID: m.EntryID, Name: m.Name, Value: value}
						}
					}).
					Open(m.Value)
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

		if selected {
			line := nameStr + " " + valueStr
			rows = append(rows, metaRowSelectedStyle.Width(innerWidth).Render("▶ "+line))
		} else {
			// Color key and value separately for unselected rows.
			line := metaRowNameStyle.Render(nameStr) + " " + valueStr
			rows = append(rows, metaRowStyle.Width(innerWidth).Render("  "+line))
		}
	}

	if len(rows) == 0 {
		rows = append(rows, metaRowStyle.Width(innerWidth).Render("No metadata"))
	}

	borderColor := lipgloss.Color("24") // dim blue
	if c.focused {
		borderColor = lipgloss.Color("33") // bright blue — glowing
	}

	title := metasTitleStyle.Render(" Metadata")
	content := lipgloss.JoinVertical(lipgloss.Left, title, lipgloss.JoinVertical(lipgloss.Left, rows...))
	return metasContainerStyle.
		BorderForeground(borderColor).
		Width(c.width - 2).
		Height(c.height - 2).
		Render(content)
}


