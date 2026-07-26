package focusmanager

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/app/program"
)

type FocusMsg struct {
	Item Focusable
}

type Focusable interface {
	Focus() tea.Cmd
	Blur() tea.Cmd
}

type Feature struct {
	items []Focusable
	index int
}

func Create() *Feature {
	return &Feature{
		items: []Focusable{},
		index: -1,
	}
}

func (f *Feature) Add(item Focusable) {
	f.items = append(f.items, item)
}

func (f *Feature) Remove(item Focusable) {
	for i, v := range f.items {
		if v == item {
			f.items = append(f.items[:i], f.items[i+1:]...)

			if f.index >= len(f.items) {
				f.index = 0
			}
			return
		}
	}
}

func (f *Feature) SetIndex(index int) tea.Cmd {
	if len(f.items) == 0 {
		return nil
	}

	cmds := []tea.Cmd{}

	if bi, ok := f.Current(); ok {
		cmds = append(cmds, bi.Blur())
	}

	f.index = index

	if ni, ok := f.Current(); ok {
		cmds = append(cmds, ni.Focus())
		cmds = append(cmds, program.Command(FocusMsg{Item: ni}))
	}

	return tea.Batch(cmds...)
}

func (f *Feature) Next() tea.Cmd {
	index := f.index + 1

	if index >= len(f.items) {
		return f.SetIndex(-1)
	}

	return f.SetIndex(index)
}

func (f *Feature) Prev() tea.Cmd {
	index := f.index - 1

	if index < 0 {
		return f.SetIndex(-1)
	}

	return f.SetIndex(index)
}

func (f Feature) Current() (Focusable, bool) {
	if f.index < 0 || f.index >= len(f.items) {
		return nil, false
	}

	return f.items[f.index], true
}

func (f Feature) IsFocused(item Focusable) bool {
	if item == nil {
		return false
	}

	if f.index < 0 || f.index >= len(f.items) {
		return false
	}

	return f.items[f.index] == item
}

func (f *Feature) Focus(item Focusable) {
	for i, v := range f.items {
		f.items[i].Blur()

		if v == item {
			f.index = i
			f.items[i].Focus()
			return
		}
	}
}
