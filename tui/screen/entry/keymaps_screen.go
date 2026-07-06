package entry

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

func (s *Screen) HandleScreenKeymaps(msg tea.Msg) tea.Cmd {
	keyMsg, ok := msg.(tea.KeyMsg)

	if !ok {
		return nil
	}

	if key.Matches(keyMsg, ScreenBindings.Up) {
		s.List.MoveUp()
		return nil
	}

	if key.Matches(keyMsg, ScreenBindings.Down) {
		s.List.MoveDown()
		return nil
	}

	if key.Matches(keyMsg, ScreenBindings.Reload) {
		s.Load()

		return messages.ToastSuccessCmd("Reloaded entries", 3*1000)
	}

	if key.Matches(keyMsg, ScreenBindings.Next) {
		s.NextPage()
		return nil
	}

	if key.Matches(keyMsg, ScreenBindings.Prev) {
		s.PreviousPage()
		return nil
	}

	if key.Matches(keyMsg, ScreenBindings.Enter) {
		name := "entry_single"
		entry, selected := s.List.SelectedEntry()

		options := map[string]any{}

		if selected {
			options["path"] = entry.Path
		}

		return messages.AddScreenCmd(messages.AddScreen{
			Name:    name,
			Options: options,
		})
	}

	if key.Matches(keyMsg, ScreenBindings.Search) {
		cb := func(input string) tea.Cmd {
			s.Search(input)

			return messages.SkipCmd()
		}

		return messages.InputCmd(messages.Input{
			Title:        "Search",
			InitialValue: s.Query,
			Callback:     cb,
		})
	}

	return nil
}
