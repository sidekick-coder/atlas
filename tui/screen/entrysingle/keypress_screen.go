package entrysingle

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
		s.EntryMetaComponent.MoveUp()
		return nil
	}

	if key.Matches(keyMsg, ScreenBindings.Down) {
		s.EntryMetaComponent.MoveDown()
		return nil
	}

	if key.Matches(keyMsg, ScreenBindings.Reload) {
		s.Load()

		return messages.ToastSuccessCmd("Reloaded")
	}

	if key.Matches(keyMsg, ScreenBindings.Sync) {
		err := s.Sync()

		if err != nil {
			return messages.ToastErrorCmd(err.Error())
		}

		return messages.ToastSuccessCmd("Synced")
	}

	if (key.Matches(keyMsg, ScreenBindings.Edit)) {
		cb := func(input string) tea.Cmd {
			err := s.SetValue(input)

			if err != nil {
				return messages.ToastErrorCmd(err.Error())
			}

			return messages.SkipCmd()
		}

		selected, exists := s.EntryMetaComponent.GetSelected()

		if !exists {
			return messages.ToastErrorCmd("No meta selected")
		}

		return messages.InputCmd(messages.Input{
			Title:         "Edit",
			InitialValue:  selected.Value,
			Callback:      cb,
		})

	}

	return nil
}
