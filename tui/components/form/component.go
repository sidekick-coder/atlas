package form

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/tui/components/form/field"
	"github.com/sidekick-coder/atlas/tui/features/focusmanager"
	"github.com/sidekick-coder/atlas/tui/features/selection"
)

type Component struct {
	values  map[string]string
	active bool
	width   int
	height  int
	fields  []field.Field
	action    config.Action

	selection *selection.Feature
	focus     *focusmanager.Feature
	registry  *Registry

	Events *Events
}

func Create() *Component {
	c := &Component{
		fields: []field.Field{},
		width:  40,
		height: 20,
		values: map[string]string{},

		selection: selection.Create(),
		focus:    focusmanager.Create(),

		registry: CreateRegistry(),

		Events: CreateEvents(),
	}

	c.registry.Load()

	return c
}

func (c *Component) Init() tea.Cmd {
	return nil
}

func (c *Component) Dispose() tea.Cmd {
	return nil
}

func (c *Component) Activate() tea.Cmd {
	c.LoadBindings()
	c.active = true
	c.focus.First()
	return nil
}

func (c *Component) Deactivate() tea.Cmd {
	c.UnloadBindings()
	c.active = false
	return nil
}

func (c *Component) GetValues() map[string]string {
	return c.values
}

func (c *Component) SetValues(values map[string]string) {
	// c.values = values
	//
	// for index, field := range c.fields {
	// 	if value, ok := values[field.Name]; ok {
	// 		c.inputs[index].SetInitialValue(value)
	// 	}
	// }
}

func (c *Component) SetAction(action config.Action) {
	c.action = action
}
