package entrytable

import (
	tea "charm.land/bubbletea/v2"

	tkey "charm.land/bubbles/v2/key"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

type Keymap struct {
	Up       key.Binding
	Down     key.Binding
	CtrlS    key.Binding
	AltA     key.Binding
	LeaderA  key.Binding
	LeaderBA key.Binding
	AA       key.Binding
}

var Bindings = Keymap{
	Up:       key.CreateBinding("<Up>", "k").SetHelp("↑/k").SetDescription("Move up"),
	CtrlS:    key.CreateBinding("<C-s>").SetHelp("Ctrl+S").SetDescription("Save"),
	AltA:     key.CreateBinding("<A-a>").SetHelp("Alt+A").SetDescription("Alt A action"),
	LeaderA:  key.CreateBinding("<leader>a").SetHelp("Leader+A").SetDescription("Leader A action"),
	LeaderBA: key.CreateBinding("<leader>ba").SetHelp("Leader+BA").SetDescription("Leader BA action"),
	AA:       key.CreateBinding("aa").SetHelp("AA").SetDescription("AA action"),
	Down:     key.CreateBinding("down").SetHelp("↓").SetDescription("Move down"),
}

func (s *Screen) GetBindings() []tkey.Binding {
	bindings := []tkey.Binding{}

	return bindings
}

func (s *Screen) RegisterBindings() tea.Cmd {
	key.Register(Bindings.Up, Bindings.Down, Bindings.CtrlS, Bindings.AltA, Bindings.LeaderA, Bindings.LeaderBA, Bindings.AA)
	return nil
}

func (s *Screen) HandleKeypress(msg tea.Msg) tea.Cmd {
	if key.Matches(Bindings.AA) {
		return messages.ToastSuccessCmd("AA key pressed")
	}

	if key.Matches(Bindings.Up) {
		return messages.ToastSuccessCmd("Up key pressed")
	}

	if key.Matches(Bindings.Down) {
		return messages.ToastSuccessCmd("Down key pressed")
	}

	if key.Matches(Bindings.CtrlS) {
		return messages.ToastSuccessCmd("Ctrl+S key pressed")
	}

	if key.Matches(Bindings.AltA) {
		return messages.ToastSuccessCmd("Alt+A key pressed")
	}

	if key.Matches(Bindings.LeaderA) {
		return messages.ToastSuccessCmd("Leader+A key pressed")
	}

	if key.Matches(Bindings.LeaderBA) {
		return messages.ToastSuccessCmd("Leader+BA key pressed")
	}

	return nil
}
