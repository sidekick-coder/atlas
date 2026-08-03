package textfield

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/borderlabel"
	"github.com/sidekick-coder/atlas/tui/components/input"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

type Component struct {
	width  int
	height int
	value  string

	input     *input.Component
	container *borderlabel.Component

	Events Events
}

func Create() *Component {
	return &Component{
		width:  40,
		height: 20,

		input:     input.Create(),
		container: borderlabel.Create(),

		Events:   *CreateEvents(),
	}
}

func (c *Component) Init() tea.Cmd {
	c.input.Events.Change.On(func() {
		c.Events.Change.Emit()
	})
	return nil
}

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(
		msg,
		c.input.Update,
	)
}

func (c *Component) Dispose() tea.Cmd {
	return chain.Dispose(
		c.input.Deactivate,
	)
}

func (c *Component) Activate() tea.Cmd {
	c.container.SetColor(theme.Current.Primary)
	c.input.Activate()
	return nil
}

func (c *Component) Deactivate() tea.Cmd {
	c.container.SetColor(theme.Current.Border)
	c.input.Deactivate()
	return nil
}

func (c *Component) SetProps(props map[string]any) {
	if l, ok := props["label"].(string); ok {
		c.container.SetLabel(l)
	}
}

func (c *Component) SetValue(content string) {
	c.input.SetValue(content)
}

func (c *Component) GetValue() string {
	return c.input.GetValue()
}
