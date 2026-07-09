package chain

import tea "charm.land/bubbletea/v2"

type InitHandler func() tea.Cmd
type UpdateHandler func(tea.Msg) tea.Cmd
type KeyHandler func(tea.KeyMsg) tea.Cmd

func Init(handlers ...InitHandler) tea.Cmd {
	for _, h := range handlers {
		if cmd := h(); cmd != nil {
			return cmd
		}
	}
	return nil
}
func Update(msg tea.Msg, handlers ...UpdateHandler) tea.Cmd {
	for _, h := range handlers {
		if cmd := h(msg); cmd != nil {
			return cmd
		}
	}
	return nil
}


func Keypress(msg tea.KeyMsg, handlers ...KeyHandler) tea.Cmd {
	for _, h := range handlers {
		if cmd := h(msg); cmd != nil {
			return cmd
		}
	}
	return nil
}
