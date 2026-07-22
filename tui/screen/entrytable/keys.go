package entrytable

import (
	"strings"

	tea "charm.land/bubbletea/v2"

	key "github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Next   key.Binding
	Prev   key.Binding
	Search key.Binding
	Reload key.Binding
	Sync   key.Binding
}

var tags = []string{"screen:entry_table"}

var Bindings = Keymap{
	Next: key.CreateBinding("n", "l").
		SetTags(tags...).
		SetDescription("next page").
		SetHelp("l"),
	Prev: key.CreateBinding("p", "h").
		SetTags(tags...).
		SetDescription("prev page").
		SetHelp("h"),
	Search: key.CreateBinding("/").
		SetTags(tags...).
		SetDescription("search").
		SetHelp("/"),
	Reload: key.CreateBinding("r").
		SetTags(tags...).
		SetDescription("reload").
		SetHelp("r"),
	Sync: key.CreateBinding("s").
		SetHelp("s").
		SetTags(tags...).
		SetDescription("sync"),
}

func (s *Screen) GetBindings() []key.Binding {
	return []key.Binding{
		Bindings.Next,
		Bindings.Prev,
		Bindings.Search,
		Bindings.Reload,
		Bindings.Sync,
	}
}

func (s *Screen) LoadBindings() tea.Cmd {
	key.Register(s.GetBindings()...)
	return nil
}

func (s *Screen) UnloadBindings() tea.Cmd {
	key.Unregister(s.GetBindings()...)
	return nil
}

func (s *Screen) HadleBinding(msg tea.KeyMsg) tea.Cmd {
	if key.Matches(Bindings.Next) {
		s.loader.Next()
		s.loader.Load()
	}

	if key.Matches(Bindings.Prev) {
		s.loader.Prev()
		s.loader.Load()
	}

	if key.Matches(Bindings.Search) {
		s.dialog.SetContent(strings.Join(s.loader.GetQuery(), " "))
		s.dialog.Open()
	}

	if key.Matches(Bindings.Sync) {
		current:= s.selection.GetCursor()

		return s.sync(current)
	}

	if key.Matches(Bindings.Reload) {
		return Reload
	}


	return nil
}
