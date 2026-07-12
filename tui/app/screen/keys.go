package screen

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Next  key.Binding
	Prev  key.Binding
	Add   key.Binding
	Close key.Binding
}

var tags = []string{"global", "screen"}

var Bindings = Keymap{
	Next: key.CreateBinding("<leader>n").
		SetTags(tags...).
		SetHelp("<leader>n").
		SetDescription("Next screen"),
	Prev: key.CreateBinding("<leader>p").
		SetTags(tags...).
		SetHelp("<leader>p").
		SetDescription("Previous screen"),
	Add: key.CreateBinding("<leader>a").
		SetTags(tags...).
		SetHelp("<leader>a").
		SetDescription("Add new screen"),
	Close: key.CreateBinding("<leader>x").
		SetTags(tags...).
		SetHelp("<leader>x").
		SetDescription("Close screen"),
}

func (f *Feature) GetBindings() []key.Binding {
	return []key.Binding{
		Bindings.Next,
		Bindings.Prev,
		Bindings.Add,
		Bindings.Close,
	}
}

func (f *Feature) LoadBindings() {
	key.Register(f.GetBindings()...)
}

func (f *Feature) UnloadBindings() {
	key.Unregister(f.GetBindings()...)
}

func (f *Feature) HandleBinding(km tea.KeyMsg) tea.Cmd {
	if key.Matches(Bindings.Next) {
		f.Next()
	}

	if key.Matches(Bindings.Prev) {
		f.Prev()
	}

	if key.Matches(Bindings.Add) {
		return Add("emtpy")
	}

	if key.Matches(Bindings.Close) {
		return Remove(f.Selection.GetCursor())
	}

	for i, b := range f.bindings {
		if key.Matches(b) {
			err := f.SetCurrent(i)

			if err != nil {
				return toast.Error(err.Error())
			}
		}
	}

	return nil
}
