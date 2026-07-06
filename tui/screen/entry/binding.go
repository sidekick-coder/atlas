package entry

import (
	"charm.land/bubbles/v2/key"
)

func (s *Screen) GetBindings() []key.Binding {
	bindings := []key.Binding{}

	bindings = append(bindings, s.GetScreenBindigs()...)
	bindings = append(bindings, s.GetUserKeymapBindings()...)

	return bindings
}
