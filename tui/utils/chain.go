package utils

import tea "charm.land/bubbletea/v2"

type Handler func(tea.Msg) tea.Cmd
type KeyHandler func(tea.KeyMsg) tea.Cmd

func Chain(msg tea.Msg, handlers ...Handler) tea.Cmd {
	for _, h := range handlers {
		if cmd := h(msg); cmd != nil {
			return cmd
		}
	}
	return nil
}

func ChainKeypress(msg tea.KeyMsg, handlers ...KeyHandler) tea.Cmd {
	for _, h := range handlers {
		if cmd := h(msg); cmd != nil {
			return cmd
		}
	}
	return nil
}
