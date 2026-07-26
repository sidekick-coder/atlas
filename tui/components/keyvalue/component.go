package keyvalue

import (
	"slices"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/inputdialog"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/selection"
)

type Item struct {
	Key   string
	Value string
}

type Component struct {
	width  int
	height int

	path  string
	items []Item

	selection *selection.Feature

	dialog *inputdialog.Component
}

func Create() *Component {
	return &Component{
		width:  100,
		height: 100,

		items: []Item{},

		selection: selection.Create(),

		dialog: inputdialog.Create(),
	}
}

func (c *Component) Init() tea.Cmd {
	return chain.Init(c.LoadBindings, c.dialog.Init)
}

func (c *Component) Dispose() tea.Cmd {
	return chain.Dispose(c.UnloadBindings, c.dialog.Dispose)
}

func (c *Component) SetSelection(selection *selection.Feature) {
	c.selection = selection
}

func (c *Component) SetItems(items []Item) {
	c.items = items
	c.selection.SetTotal(len(items))
	c.SortByKey()
}

func (c *Component) SortByKey() {
	slices.SortFunc(c.items, func(a, b Item) int {
		return strings.Compare(a.Key, b.Key)
	})
}

func (c *Component) Clear() {
	c.items = []Item{}
	c.selection.Clear()
}

func (c *Component) GetSelected() (Item, bool) {
	if len(c.items) == 0 {
		return Item{}, false
	}

	cursor := c.selection.GetCursor()

	if cursor < 0 || cursor >= len(c.items) {
		return Item{}, false
	}

	return c.items[cursor], true
}
