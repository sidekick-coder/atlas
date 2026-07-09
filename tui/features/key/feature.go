package key

import tea "charm.land/bubbletea/v2"

var manager = NewManager()

func Register(bindings ...Binding) {
	manager.Register(bindings...)
}

func HandleKeypress(msg tea.Msg) tea.Cmd {
	return manager.HandleKeypress(msg)
}

func GetBindings() []Binding {
	return manager.GetBindings()
}

func ClearBindings() {
	manager.ClearBindings()
}
