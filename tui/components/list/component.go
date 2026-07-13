package list

import tea "charm.land/bubbletea/v2"

type Component struct {
	width    int
	height   int
	cursor int
	onSelect func(cursor int) tea.Cmd
	items    []string
}

func Create() *Component {
	return &Component{
		width:    100,
		height:   100,
		cursor: 0,
		items:    []string{},
	}
}

func (c *Component) OnSelect(f func(cursor int) tea.Cmd) *Component {
	c.onSelect = f
	return c
}

func (c *Component) Init() tea.Cmd {
	c.LoadBindings()
	return nil
}

func (c *Component) Dispose() tea.Cmd {
	c.UnloadBindings()
	return nil
}
