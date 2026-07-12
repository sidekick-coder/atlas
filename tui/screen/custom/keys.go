package custom

import (
	"log/slog"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

type KeyMap struct {
	Up   key.Binding
	Down key.Binding
}

var tags = []string{"screen", "custom"}

var Bindings = KeyMap{
	Up: key.CreateBinding("k", "<up>").
		SetTags(tags...).
		SetDescription("component up").
		SetHelp("k/up"),
	Down: key.CreateBinding("j", "<down>").
		SetTags(tags...).
		SetDescription("component down").
		SetHelp("j/down"),
}

func (s *Screen) GetBindings() []key.Binding {
	return []key.Binding{
		Bindings.Up,
		Bindings.Down,
	}
}

func (s *Screen) LoadBindings() tea.Cmd {
	key.Register(s.GetBindings()...)
	return nil
}

func (s *Screen) UnloadBindings() tea.Cmd {
	slog.Info("Unloading custom screen bindings")
	key.Unregister(s.GetBindings()...)
	return nil
}

func (s *Screen) HandleBinding(msg tea.KeyMsg) tea.Cmd {
	if key.Matches(Bindings.Up) {
		s.selection.Next()
	}

	if key.Matches(Bindings.Down) {
		s.selection.Prev()
	}

	return nil
}
