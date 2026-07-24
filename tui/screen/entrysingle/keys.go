package entrysingle

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

type Keymap struct {
	Reload key.Binding
	Sync   key.Binding
}

var Bindings = Keymap{
	Reload: key.CreateBinding("<leader>r", "<leader><r>").
		SetDescription("reload").
		SetHelp("<leader>r"),
	Sync: key.CreateBinding("<leader>s", "<leader><s>").
		SetDescription("sync").
		SetHelp("<leader>s"),
}

func (s *Screen) GetBindigs() []key.Binding {
	return []key.Binding{
		Bindings.Reload,
		Bindings.Sync,
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

	return nil
}
