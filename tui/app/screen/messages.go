package screen

import tea "charm.land/bubbletea/v2"

type SizeMsg struct {
	Width  int
	Height int
}

type AddMsg struct {
	Name    string
	Options map[string]any
}

func Add(name string, options map[string]any) tea.Cmd {
	return func() tea.Msg {
		return AddMsg{
			Name:    name,
			Options: options,
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
