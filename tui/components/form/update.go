package form

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

func (c *Component) DisableInputs() {
	for _, input := range c.inputs {
		input.Disable()
	}
}

func (c *Component) Refresh() {
	cursor := c.selection.GetCursor()

	c.DisableInputs()

	for index, input := range c.inputs {
		if cursor == index {
			input.Enable()
		}
	}
}

func (c *Component) submit() tea.Cmd {
	if c.onSubmit == nil {
		return toast.Error("No submit handler defined")
	}

	values := map[string]any{}

	for index, field := range c.fields {
		input := c.inputs[index]

		values[field.Name] = input.GetValue()
	}

	c.onSubmit(values)

	return nil
}

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(msg, chain.OnKey(c.HandleBindings), chain.OnEntity(c.inputs))
}
