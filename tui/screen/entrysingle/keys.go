package entrysingle

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

type Keymap struct {
	Up     key.Binding
	Down   key.Binding
	Edit   key.Binding
	Add    key.Binding
	Delete key.Binding
	Reload key.Binding
	Sync   key.Binding
}

var Bindings = Keymap{
	Up: key.CreateBinding("k", "up").
		SetDescription("up").
		SetHelp("k/up"),
	Down: key.CreateBinding("j", "down").
		SetDescription("down").
		SetHelp("j/down"),
	Edit: key.CreateBinding("e").
		SetDescription("edit").
		SetHelp("e"),
	Add: key.CreateBinding("a", "add").
		SetDescription("add").
		SetHelp("a"),
	Delete: key.CreateBinding("d", "delete").
		SetDescription("delete").
		SetHelp("d"),
	Reload: key.CreateBinding("<leader>r", "<leader><r>").
		SetDescription("reload").
		SetHelp("<leader>r"),
	Sync: key.CreateBinding("<leader>s", "<leader><s>").
		SetDescription("sync").
		SetHelp("<leader>s"),
}

func (s *Screen) GetBindigs() []key.Binding {
	return []key.Binding{
		Bindings.Up,
		Bindings.Down,
		Bindings.Edit,
		Bindings.Reload,
		Bindings.Sync,
		Bindings.Add,
		Bindings.Delete,
	}
}

func (s *Screen) LoadBindings() tea.Cmd {
	key.Register(s.GetBindigs()...)
	return nil
}

func (s *Screen) UnloadBindings() tea.Cmd {
	key.Unregister(s.GetBindigs()...)
	return nil
}

func (s *Screen) HandleBindings(msg tea.KeyMsg) tea.Cmd {
	if key.Matches(Bindings.Up) {
		s.EntryMetaComponent.MoveUp()
		return nil
	}

	if key.Matches(Bindings.Down) {
		s.EntryMetaComponent.MoveDown()
		return nil
	}

	if key.Matches(Bindings.Reload) {
		s.Load()

		return messages.ToastSuccessCmd("Reloaded")
	}

	if key.Matches(Bindings.Sync) {
		err := s.Sync()

		if err != nil {
			return messages.ToastErrorCmd(err.Error())
		}

		return messages.ToastSuccessCmd("Synced")
	}

	if key.Matches(Bindings.Edit) {
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
			Title:        "Edit",
			InitialValue: selected.Value,
			Callback:     cb,
		})
	}

	if key.Matches(Bindings.Add) {
		cb := func(input string) tea.Cmd {
			err := s.SetMeta(input, "")

			if err != nil {
				return messages.ToastErrorCmd(err.Error())
			}

			return messages.SkipCmd()
		}

		return messages.InputCmd(messages.Input{
			Title:    "Add",
			Callback: cb,
		})
	}

	if key.Matches(Bindings.Delete) {
		err := s.UnsetMetaSelected()

		if err != nil {
			return messages.ToastErrorCmd(err.Error())
		}

		return messages.SkipCmd()
	}

	return nil
}
