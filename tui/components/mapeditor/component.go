package mapeditor

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/borderlabel"
	"github.com/sidekick-coder/atlas/tui/components/dialog"
	"github.com/sidekick-coder/atlas/tui/components/input"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/selection"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

type Component struct {
	fields []Field
	values map[string]string
	width  int
	height int

	onClose  func()
	onOpen   func()
	onSubmit func(values map[string]string)

	dialog    *dialog.Component
	selection *selection.Feature
	inputs    []*input.Component

	fieldBorder         *borderlabel.Component
	fieldBorderSelected *borderlabel.Component
}

func Create(args ...map[string]any) (*Component, error) {

	c := &Component{
		fields: []Field{},
		width:  40,
		height: 20,
		values: map[string]string{},

		dialog:    dialog.Create().SetTitle("Map Editor"),
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

	c.fieldBorder.SetWidth(c.width - 6)
	c.fieldBorderSelected.SetWidth(c.width - 6) // 4 padding

	return c, nil
}

func (c *Component) OnOpen(fn func()) {
	c.onOpen = fn
}

func (c *Component) OnSubmit(fn func(values map[string]string)) {
	c.onSubmit = fn
}

func (c *Component) OnClose(fn func()) {
	c.onClose = fn
}

func (c *Component) IsOpen() bool {
	return c.dialog.IsOpen()
}

func (c *Component) GetFields() []Field {
	return c.fields
}

func (c *Component) GetField(index int) (Field, bool) {
	if index < 0 || index >= len(c.fields) {
		return Field{}, false
	}

	return c.fields[index], true
}

func (c *Component) GetFieldSelected() (Field, bool) {
	index := c.selection.GetCursor()

	if index < 0 || index >= len(c.fields) {
		return Field{}, false
	}

	return c.fields[index], true
}

func (c *Component) SetFields(fields []Field) {
	c.fields = fields
	c.selection.SetTotal(len(fields))
	c.selection.SetCursor(0)

	inputs := []*input.Component{}

	for range fields {
		input := input.New()
		input.SetWidth(c.width - 4) // 4 padding

		inputs = append(inputs, input)
	}

	c.inputs = inputs
}

func (c *Component) GetValues() map[string]string {
	return c.values
}

func (c *Component) SetValues(values map[string]string) {
	c.values = values

	for index, field := range c.fields {
		if value, ok := values[field.FielName]; ok {
			c.inputs[index].SetInitialValue(value)
		}
	}
}

func (c *Component) Init() tea.Cmd {
	c.dialog.Events.Close.On(func() {
		c.DisableInputs()
	})

	return chain.Init(c.dialog.Init, c.InitRender)
}

func (c *Component) Dispose() tea.Cmd {
	return chain.Dispose(c.dialog.Dispose)
}

func (c *Component) OnFocus() {
	c.LoadBindings()
}

func (c *Component) OnBlur() {

}
