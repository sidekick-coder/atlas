package entry

import (
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/bubbles/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

type ScreenKeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Enter  key.Binding
	Next   key.Binding 
	Prev   key.Binding
	Sync   key.Binding
	Reload key.Binding
	Search  key.Binding
}

var ScreenBindings = ScreenKeyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("k/up", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("j/down", "down"),
	),
	Next: key.NewBinding(
		key.WithKeys("l", "right"),
		key.WithHelp("l/right", "next"),
	),
	Prev: key.NewBinding(
		key.WithKeys("h", "left"),
		key.WithHelp("h/left", "prev"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Sync: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "sync selected entry"),
	),
	Reload: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "reload"),
	),
	Search: key.NewBinding( 
		key.WithKeys("/"),
		key.WithHelp("/", "search"),
	),
}

func (s *Screen) GetScreenBindigs() []key.Binding {
	bindings := []key.Binding{}

	bindings = append(bindings,
		ScreenBindings.Up,
		ScreenBindings.Down,
		ScreenBindings.Next,
		ScreenBindings.Prev,
		ScreenBindings.Sync,
		ScreenBindings.Enter,
		ScreenBindings.Reload,
		ScreenBindings.Search,
	)

	return bindings
}

func (s *Screen) HandleScreenKeymaps(msg tea.Msg) tea.Cmd {
	keyMsg, ok := msg.(tea.KeyMsg)

	if !ok {
		return nil
	}

	if key.Matches(keyMsg, ScreenBindings.Up) {
		s.List.MoveUp()
		return nil
	}

	if key.Matches(keyMsg, ScreenBindings.Down) {
		s.List.MoveDown()
		return nil
	}

	if key.Matches(keyMsg, ScreenBindings.Reload) {
		s.Load()

		return messages.ToastSuccessCmd("Reloaded entries", 3*1000)
	}

	if key.Matches(keyMsg, ScreenBindings.Next) {
		s.NextPage()
		return nil
	}

	if key.Matches(keyMsg, ScreenBindings.Prev) {
		s.PreviousPage()
		return nil
	}

	if key.Matches(keyMsg, ScreenBindings.Enter) {
		name := "entry-single"
		entry := s.List.SelectedEntry()

		options := map[string]any{}
		options["path"] = entry.Path

		return messages.AddScreenCmd(name, options)
	}

	if (key.Matches(keyMsg, ScreenBindings.Search)) {
		return messages.InputCmd("Search", func(input string) tea.Cmd {
			s.Search(input)

			return messages.SkipCmd()
		})
	}

	return nil
}
