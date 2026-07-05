package entry

import (
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/bubbles/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

func (s *Screen) Update(msg tea.Msg) tea.Cmd {
	handlers := []func(tea.Msg) tea.Cmd{}

	handlers = append(handlers, s.HandleKeyMaps)

	for _, handler := range handlers {
		cmd := handler(msg)

		if cmd != nil {
			return cmd
		}
	}

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if key.Matches(msg, Bindings.Up) {
			s.List.MoveUp()
			return nil
		}

		if key.Matches(msg, Bindings.Down) {
			s.List.MoveDown()
			return nil
		}

		if key.Matches(msg, Bindings.Enter) {
			name := "entry-single"
			entry := s.List.SelectedEntry()

			options := map[string]any{}
			options["entry_id"] = entry.ID

			return messages.AddScreenCmd(name, options)
		}

	}

	return nil
}
