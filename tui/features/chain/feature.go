package chain

import tea "charm.land/bubbletea/v2"

type ReturnCommand func() tea.Cmd
type ReceiveMessageReturnCommand func(msg tea.Msg) tea.Cmd
type ReceiveKeyReturnCommand func(msg tea.KeyMsg) tea.Cmd

func Cmd(handlers ...ReturnCommand) tea.Cmd {
	for _, h := range handlers {
		if cmd := h(); cmd != nil {
			return cmd
		}
	}

	return nil
}

func Init(handlers ...ReturnCommand) tea.Cmd {
	return Cmd(handlers...)
}

func Dispose(handlers ...ReturnCommand) tea.Cmd {
	return Cmd(handlers...)
}

func OnKey(keyHandler ReceiveKeyReturnCommand) ReceiveMessageReturnCommand {
	return func(msg tea.Msg) tea.Cmd {
		if keyMsg, ok := msg.(tea.KeyMsg); ok {
			return keyHandler(keyMsg)
		}

		return nil
	}
}

func Update(msg tea.Msg, handlers ...func(msg tea.Msg) tea.Cmd) tea.Cmd {
	for _, h := range handlers {
		if cmd := h(msg); cmd != nil {
			return cmd
		}
	}
	return nil
}

func Keypress(msg tea.KeyMsg, handlers ...ReceiveKeyReturnCommand) tea.Cmd {
	for _, h := range handlers {
		if cmd := h(msg); cmd != nil {
			return cmd
		}
	}
	return nil
}
