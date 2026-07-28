package keyvalue

import (
	"slices"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/inputdialog"
	"github.com/sidekick-coder/atlas/tui/components/viewport"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/selection"
)

type Item struct {
	Key    string
	Value  string
	Header bool
}

type Component struct {
	width  int
	height int

	path  string
	items []Item

	selection *selection.Feature

	dialog   *inputdialog.Component
	viewport *viewport.Component
}

func Create() *Component {
	return &Component{
		width:  100,
		height: 100,

		items: []Item{},

		selection: selection.Create(),

		dialog: inputdialog.Create(),
		viewport: viewport.Create(),
	}
}

func (c *Component) Init() tea.Cmd {
	return chain.Init(c.dialog.Init)
}

func (c *Component) Activate() tea.Cmd {
	return chain.Cmd(c.LoadBindings)
}

func (c *Component) Deactivate() tea.Cmd {
	return chain.Cmd(c.UnloadBindings)
}

func (c *Component) Dispose() tea.Cmd {
	return chain.Dispose(c.Deactivate, c.dialog.Dispose)
}

func (c *Component) SetSelection(selection *selection.Feature) {
	c.selection = selection
}

func (c *Component) SetItems(items []Item) {
	c.items = items
	c.selection.SetTotal(len(items))
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
