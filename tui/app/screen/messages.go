package screen

import (
	"log/slog"

	tea "charm.land/bubbletea/v2"
)

type SizeMsg struct {
	Width  int
	Height int
}

func (f *Feature) Size() tea.Msg {
	slog.Info("sending size message", slog.Int("width", f.windowWidth), slog.Int("height", f.windowHeight-7))
	return SizeMsg{
		Width:  f.windowWidth,
		Height: f.windowHeight - 7,
	}
}

type AddMsg struct {
	Name    string
	Options map[string]any
}

func Add(name string, options ...map[string]any) tea.Cmd {
	o := map[string]any{}

	if len(options) > 0 {
		o = options[0]
	}
	return func() tea.Msg {
		return AddMsg{
			Name:    name,
			Options: o,
		}
	}
}

type RemoveMsg struct {
	Index int
}

func Remove(index int) tea.Cmd {
	return func() tea.Msg {
		return RemoveMsg{
			Index: index,
		}
	}
}

type ReplaceMsg struct {
	Index   int
	Name    string
	Options map[string]any
}

func Replace(index int, name string, options map[string]any) tea.Cmd {
	return func() tea.Msg {
		return ReplaceMsg{
			Index:   index,
			Name:    name,
			Options: options,
		}
	}
}

type ReplaceCurrentMsg struct {
	Name    string
	Options map[string]any
}

func ReplaceCurrent(name string, options map[string]any) tea.Cmd {
	return func() tea.Msg {
		return ReplaceCurrentMsg{
			Name:    name,
			Options: options,
		}
	}
}
