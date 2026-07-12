package form

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/utils/maputil"
	"github.com/sidekick-coder/atlas/tui/components/borderlabel"
	"github.com/sidekick-coder/atlas/tui/components/input"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/selection"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

type Component struct {
	fields []Field
	values map[string]string
	focused bool
	width  int 
	height int

	selection *selection.Feature
	inputs    []*input.Input

	fieldBorder         *borderlabel.Component
	fieldBorderSelected *borderlabel.Component
}

func Create(args ...map[string]any) (*Component, error) {

	c := &Component{
		fields: []Field{},
		width: 40,
		height: 20,
		values: map[string]string{},

		selection: selection.Create(),

		fieldBorder: borderlabel.
			Create().
			SetColor(theme.Current.Muted),
		fieldBorderSelected: borderlabel.
			Create().
			SetColor(theme.Current.Primary),
	}

	props := map[string]any{}

	if len(args) > 0 {
		props = args[0]
	}

	if fp, ok := props["fields"]; ok {
		pf, err := CreateFieldsFromArray(fp)

		if err != nil {
			return nil, err
		}

		c.SetFields(pf)
	}

	if w, ok := props["width"].(int); ok {
		c.width = w
	}

	if h, ok := props["height"].(int); ok {
		c.height = h
	}

	if v, ok := props["values"].(map[string]any); ok {
		c.SetValues(maputil.String(v))
	}

	c.fieldBorder.SetWidth(c.width - 6)
	c.fieldBorderSelected.SetWidth(c.width - 6) // 4 padding

	return c, nil
}

func (c *Component) GetValues() map[string]string {
	return c.values
}

func (c *Component) SetValues(values map[string]string) {
	c.values = values

	for index, field := range c.fields {
		if value, ok := values[field.Name]; ok {
			c.inputs[index].SetInitialValue(value)
		}
	}
}

func (c *Component) Init() tea.Cmd {
	return chain.Init(c.InitRender)
}

func (c *Component) Dispose() tea.Cmd {
	return nil
}

func (c *Component) OnFocus() {
	c.LoadBindings()
	c.Refresh()
	c.focused = true
}

func (c *Component) OnBlur() {
	c.UnloadBindings()
	c.DisableInputs()
	c.focused = false
}
