package components

import (
	"fmt"
	"path/filepath"
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	"charm.land/bubbles/v2/key"
	"github.com/sidekick-coder/atlas/internal/models"
)

var extIcons = map[string]string{
	// Text / docs
	".md":   "󰍔 ",
	".txt":  " ",
	".pdf":  " ",
	".doc":  "󰈬 ",
	".docx": "󰈬 ",
	// Code
	".go":   " ",
	".js":   " ",
	".ts":   " ",
	".jsx":  " ",
	".tsx":  " ",
	".py":   " ",
	".rb":   " ",
	".rs":   " ",
	".c":    " ",
	".cpp":  " ",
	".h":    " ",
	".java": " ",
	".cs":   "󰌛 ",
	".php":  " ",
	".sh":   " ",
	".bash": " ",
	".zsh":  " ",
	".fish": " ",
	".lua":  " ",
	".vim":  " ",
	// Config / data
	".json":  " ",
	".yaml":  " ",
	".yml":   " ",
	".toml":  " ",
	".xml":   "󰗀 ",
	".env":   " ",
	".ini":   " ",
	".conf":  " ",
	".lock":  " ",
	// Images
	".png":  " ",
	".jpg":  " ",
	".jpeg": " ",
	".gif":  " ",
	".svg":  "󰜡 ",
	".ico":  " ",
	// Audio / video
	".mp3": " ",
	".mp4": " ",
	".mkv": " ",
	".wav": " ",
	// Archives
	".zip": " ",
	".tar": " ",
	".gz":  " ",
	".rar": " ",
	// Web
	".html": " ",
	".css":  " ",
	".scss": " ",
	// Database
	".sql":    " ",
	".sqlite": " ",
	".db":     " ",
}

func fileIcon(path string) string {
	// Directory (no extension, or ends with /)
	if strings.HasSuffix(path, "/") {
		return "󰉋 "
	}
	ext := strings.ToLower(filepath.Ext(path))
	if icon, ok := extIcons[ext]; ok {
		return icon
	}
	if ext == "" {
		return " "
	}
	return " "
}

var (
	listItemStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Foreground(lipgloss.Color("244"))

	listItemSelectedStyle = lipgloss.NewStyle().
				Padding(0, 1).
				Bold(true).
				Foreground(lipgloss.Color("12"))

	listContainerStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder())

	listTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("12")).
			Padding(0, 1)
)

type EntrySelectedMsg struct {
	Entry models.Entry
}

type EntryList struct {
	entries    []models.Entry
	cursor     int
	focused    bool
	width      int
	height     int
	page       int
	pageSize   int
	totalCount int
}

func NewEntryList() *EntryList {
	return &EntryList{pageSize: 30}
}

func (c *EntryList) SetEntries(entries []models.Entry, totalCount int) {
	c.entries = entries
	c.totalCount = totalCount
	c.cursor = 0
}

func (c *EntryList) Page() int    { return c.page }
func (c *EntryList) PageSize() int { return c.pageSize }
func (c *EntryList) TotalPages() int {
	if c.pageSize <= 0 {
		return 1
	}
	t := (c.totalCount + c.pageSize - 1) / c.pageSize
	if t < 1 {
		return 1
	}
	return t
}

func (c *EntryList) SetFocused(focused bool) {
	c.focused = focused
}

func (c *EntryList) SetSize(width, height int) {
	c.width = width
	c.height = height
}

func (c *EntryList) SelectedEntry() *models.Entry {
	if len(c.entries) == 0 {
		return nil
	}
	return &c.entries[c.cursor]
}

// EntryPageChangedMsg is emitted when the user navigates to a different page.
type EntryPageChangedMsg struct {
	Page   int
	Offset int
	Limit  int
}

func (c *EntryList) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Up):
			if c.cursor > 0 {
				c.cursor--
				return func() tea.Msg { return EntrySelectedMsg{Entry: c.entries[c.cursor]} }
			}
		case key.Matches(msg, DefaultKeyMap.Down):
			if c.cursor < len(c.entries)-1 {
				c.cursor++
				return func() tea.Msg { return EntrySelectedMsg{Entry: c.entries[c.cursor]} }
			}
		case key.Matches(msg, DefaultKeyMap.PageNext):
			if c.page < c.TotalPages()-1 {
				c.page++
				c.cursor = 0
				return func() tea.Msg {
					return EntryPageChangedMsg{Page: c.page, Offset: c.page * c.pageSize, Limit: c.pageSize}
				}
			}
		case key.Matches(msg, DefaultKeyMap.PagePrev):
			if c.page > 0 {
				c.page--
				c.cursor = 0
				return func() tea.Msg {
					return EntryPageChangedMsg{Page: c.page, Offset: c.page * c.pageSize, Limit: c.pageSize}
				}
			}
		}
	}
	return nil
}

func (c *EntryList) View() string {
	innerWidth := c.width - 2 // account for border
	if innerWidth < 1 {
		innerWidth = 1
	}

	var rows []string
	for i, entry := range c.entries {
		icon := fileIcon(entry.Path)
		label := icon + entry.Path
		if len(label) > innerWidth-2 {
			label = label[:innerWidth-2]
		}
		if i == c.cursor {
			rows = append(rows, listItemSelectedStyle.Width(innerWidth).Render("▶ "+label))
		} else {
			rows = append(rows, listItemStyle.Width(innerWidth).Render("  "+label))
		}
	}

	if len(rows) == 0 {
		rows = append(rows, listItemStyle.Width(innerWidth).Render("No entries"))
	}

	borderColor := lipgloss.Color("24") // dim blue
	if c.focused {
		borderColor = lipgloss.Color("33") // bright blue — glowing
	}

	// Pagination footer: "page 2 / 5  ·  42 total"
	dimStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	accentStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("33")).Bold(true)
	paginationText := accentStyle.Render(fmt.Sprintf("%d", c.page+1)) +
		dimStyle.Render(fmt.Sprintf(" / %d  ·  %d total", c.TotalPages(), c.totalCount))
	pagination := dimStyle.Width(innerWidth).Align(lipgloss.Right).Render(paginationText)

	title := listTitleStyle.Render(" Entries")
	content := lipgloss.JoinVertical(lipgloss.Left, title, lipgloss.JoinVertical(lipgloss.Left, rows...), pagination)
	return listContainerStyle.
		BorderForeground(borderColor).
		Width(c.width - 2).
		Height(c.height - 2).
		Render(content)
}
