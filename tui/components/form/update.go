package form

import (
	"maps"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/template"
	"github.com/sidekick-coder/atlas/tui/action"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(msg, chain.OnKey(c.HandleBindings), c.UpdateField)
}

func (c *Component) UpdateField(msg tea.Msg) tea.Cmd {
	index := c.focus.GetIndex()

	if index < 0 || index >= len(c.fields) {
		return nil
	}

	field := c.fields[index]

	return field.Update(msg)
}

func (c *Component) submit() tea.Cmd {
	values := map[string]any{}

	for _, field := range c.fields {
		values[field.GetName()] = field.GetValue()
	}

	c.Events.Submit.Emit()

	id := c.action.Type
	opts := map[string]any{}

	maps.Copy(opts, c.action.Options)

	computed, err := template.EvaluateMap(opts, values)

	if err != nil {
		return toast.Error(err.Error())
	}

	return action.Execute(id, computed)
}
