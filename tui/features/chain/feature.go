package chain

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/toast"
)

type ReceiveKeyReturnCommand func(msg tea.KeyMsg) tea.Cmd

type WithUpdate interface {
	Update(msg tea.Msg) tea.Cmd
}

func Cmd(handlers ...func() tea.Cmd) tea.Cmd {
	for _, h := range handlers {
		if cmd := h(); cmd != nil {
			return cmd
		}
	}

	return nil
}

func Init(handlers ...func() tea.Cmd) tea.Cmd {
	return Cmd(handlers...)
}

func Dispose(handlers ...func() tea.Cmd) tea.Cmd {
	return Cmd(handlers...)
}

func OnKey(keyHandler ReceiveKeyReturnCommand) func(msg tea.Msg) tea.Cmd {
	return func(msg tea.Msg) tea.Cmd {
		if keyMsg, ok := msg.(tea.KeyMsg); ok {
			return keyHandler(keyMsg)
		}

		return nil
	}
}

func OnVoid(h func()) func() tea.Cmd {
	return func() tea.Cmd {
		h()
		return nil
	}
}

func OnCondition(h func(tea.Msg) tea.Cmd, cond bool) func(msg tea.Msg) tea.Cmd {
	return func(msg tea.Msg) tea.Cmd {
		if cond {
			return h(msg)
		}

		return nil
	}
}

func OnError(fn func() error) func() tea.Cmd {
	return func() tea.Cmd {
		err := fn()

		if err != nil {
			return toast.Error(err.Error())
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

func OnEntity[T WithUpdate](entities []T) func(msg tea.Msg) tea.Cmd {
	return func(msg tea.Msg) tea.Cmd {
		for _, entity := range entities {
			cmd := entity.Update(msg)

			if cmd != nil {
				return cmd
			}
		}

		return nil
	}
}

func Keypress(msg tea.KeyMsg, handlers ...ReceiveKeyReturnCommand) tea.Cmd {
	for _, h := range handlers {
		if cmd := h(msg); cmd != nil {
			return cmd
		}
	}
	return nil
}
