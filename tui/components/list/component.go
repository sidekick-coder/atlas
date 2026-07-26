package list

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/selection"
)

type Component struct {
	width    int
	height   int
	items    []string

	selection *selection.Feature
}

func Create() *Component {
	return &Component{
		width:  100,
		height: 100,
		items:  []string{},
		selection: selection.Create(),
	}
}

func (c *Component) Init() tea.Cmd {
	return nil
}

func (c *Component) Dispose() tea.Cmd {
	c.UnloadBindings()
	return nil
}

func (c *Component) Activate() tea.Cmd {
	c.LoadBindings()
	return nil
}

func (c *Component) Deactive() tea.Cmd {
	c.UnloadBindings()
	return nil
}

func (c *Component) SetItems(items []string) {
	c.items = items

	c.selection.SetTotal(len(items))
}

func (c *Component) SetSelection(selection *selection.Feature) {
	c.selection = selection
}

func (c *Component) Focus() {
	c.LoadBindings()
}

func (c *Component) Blur() {
	c.UnloadBindings()
}
