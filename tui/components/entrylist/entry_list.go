package entrylist

import (
	lipgloss "charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/internal/models"
)

type EntryList struct {
	Width  int
	Height int
	Focus  bool
	CurrentIndex int
	Entries []models.Entry
}

func New() *EntryList {
	return &EntryList{
		Width:  100,
		Height: 100,
		CurrentIndex: 0,
		Focus:  false,
		Entries: []models.Entry{},
	}
}

func (e *EntryList) SetSize(width, height int) {
	e.Width = width
	e.Height = height
}

func (e *EntryList) SetFocus(focus bool) {
	e.Focus = focus
}

func (e *EntryList) SetEntries(entries []models.Entry) {
	e.Entries = entries

	maxIndex := len(entries) - 1

	if e.CurrentIndex > maxIndex {
		e.CurrentIndex = maxIndex
	}
}

func (e *EntryList) Render() string {
	border := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Width(e.Width - 4).
		Height(e.Height - 4).
		Margin(0, 2).
		BorderForeground(lipgloss.Color("12"))

	if e.Focus {
		border = border.BorderForeground(lipgloss.Color("33"))
	}

	normal := lipgloss.NewStyle().
		Width(e.Width).
		Height(1).
		Padding(0, 1)
		

	selected := lipgloss.NewStyle().
		Background(lipgloss.Color("12")).
		Width(e.Width - 4).Padding(0, 1).
		Foreground(lipgloss.Color("0"))

	var items []string


	for index, entry := range e.Entries {
		if index == e.CurrentIndex {
			result := selected.Render(entry.Path)

			items = append(items, result)
			continue
		}

		items = append(items, normal.Render(entry.Path))
	}

	row := lipgloss.JoinVertical(lipgloss.Left, items...)

	return border.Render(row)
}

func (e *EntryList) MoveUp() {
	if e.CurrentIndex > 0 {
		e.CurrentIndex--
	}
}

func (e *EntryList) MoveDown() {
	if e.CurrentIndex < len(e.Entries)-1 {
		e.CurrentIndex++
	}
}

func (e *EntryList) SelectedEntry() (models.Entry, bool) {
	if e.CurrentIndex < 0 || e.CurrentIndex >= len(e.Entries) {
		return models.Entry{}, false
	}

	return e.Entries[e.CurrentIndex], true
}

func (e *EntryList) HasSeletion() bool {
	if len(e.Entries) == 0 {
		return false
	}

	return true
}

func (e *EntryList) SelectedEntryID() int64 {
	entry := e.Entries[e.CurrentIndex] 

	return entry.ID
}
