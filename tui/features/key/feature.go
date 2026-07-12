package key

import tea "charm.land/bubbletea/v2"

var manager = NewManager()

func Register(bindings ...Binding) {
	manager.Register(bindings...)
}

func Unregister(bindings ...Binding) {
	manager.Unregister(bindings...)
}

func HandleKeypress(msg tea.Msg) tea.Cmd {
	return manager.HandleKeypress(msg)
}

func GetBindings() []Binding {
	return manager.GetBindings()
}

func GetPendingTokens() []string {
	return manager.GetPendingTokens()
}

func GetBindingsByTags(tags ...string) []Binding {
	return manager.GetBindingsByTags(tags...)
}

func ClearBindings() {
	manager.ClearBindings()
}
