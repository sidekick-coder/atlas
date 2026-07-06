package entrysingle 

import (
	"charm.land/bubbles/v2/key"
)

func (s *Screen) GetBindings() []key.Binding {
	bindings := []key.Binding{}

	bindings = append(bindings, s.GetScreenBindigs()...)

	return bindings
}
