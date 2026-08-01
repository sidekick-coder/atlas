package input

import (
	tea "charm.land/bubbletea/v2"
)

type Component struct {
	buf     []rune
	cursor  int
	width   int
	height  int
	offset  int
	active bool
}

func Create() *Component {
	return &Component{
		width:   80,
		height:  3,
		active: false,
	}
}

func New() *Component {
	return Create()
}

func (i *Component) SetSize(width, height int) *Component {
	i.width = width
	i.height = height

	return i
}

func (i *Component) SetWidth(width int) *Component {
	i.width = width
	return i
}

func (i *Component) Activate() tea.Cmd {
	i.active = true
	i.LoadBindings()
	return nil
}

func (i *Component) Deactivate() tea.Cmd {
	i.active = false
	i.UnloadBindings()
	return nil
}

// Deprecated: Use Enable() instead.
func (i *Component) Enable() *Component {
	i.Activate()
	return i
}

// Deprecated: Use Deactivate() instead.
func (i *Component) Disable() *Component {
	// i.active = false
	i.UnloadBindings()
	return i
}

func (i *Component) SetValue(v string) {
	i.buf = []rune(v)
	i.cursor = len(i.buf)
}

func (i *Component) SetInitialValue(initialValue string) {
	i.buf = []rune(initialValue)
	i.cursor = len(i.buf)
}

func (i *Component) GetValue() string {
	return string(i.buf)
}
