package form

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)


func (c *Component) DisableInputs() {
	for _, input := range c.inputs {
		input.Disable()
		input.UnloadBindings()
	}
}

func (c *Component) Refresh(){
	cursor := c.selection.GetCursor()

	for index, input := range c.inputs {
		input.Disable()
		input.LoadBindings()

		if cursor == index {
			input.Enable()
		}
	}
}

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(msg, chain.OnKey(c.HandleBindings), chain.OnEntity(c.inputs))
}
