package entrylist

import (
	"strings"

	tea "charm.land/bubbletea/v2"

	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/entrycontroller"
	key "github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	NextPage key.Binding
	PrevPage key.Binding
	Search   key.Binding
	Reload   key.Binding
	Sync     key.Binding
}

var tags = []string{"screen:entry_table"}

var Bindings = Keymap{
	NextPage: key.CreateBinding("n", "l").
		SetTags(tags...).
		SetDescription("next page").
		SetHelp("l"),
	PrevPage: key.CreateBinding("p", "h").
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

func (c *Component) GetBindings() []key.Binding {
	return []key.Binding{
		Bindings.NextPage,
		Bindings.PrevPage,
		Bindings.Search,
		Bindings.Reload,
		Bindings.Sync,
	}
}

func (c *Component) LoadBindings() tea.Cmd {
	key.Register(c.GetBindings()...)
	return nil
}

func (c *Component) UnloadBindings() tea.Cmd {
	key.Unregister(c.GetBindings()...)
	return nil
}

func (c *Component) HadleBinding(msg tea.KeyMsg) tea.Cmd {
	if key.Matches(Bindings.NextPage) {
		c.loader.Next()
	}

	if key.Matches(Bindings.PrevPage) {
		c.loader.Prev()
	}

	if key.Matches(Bindings.Search) {
		c.dialog.SetContent(strings.Join(c.loader.GetQuery(), " "))
		c.dialog.Open()
	}

	if key.Matches(Bindings.Reload) {
		c.loader.Load()

		return toast.Success("Reloaded")
	}

	if key.Matches(Bindings.Sync) {
		current, exists := c.GetCurrent()

		if exists {
			return entrycontroller.Sync(current.Path)
		}
	}

	return nil
}
