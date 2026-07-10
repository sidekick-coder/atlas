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

type Field struct {
	Label  string
	FielName string
}

type Component struct {
	fields []Field
	values map[string]string

	onClose func()
	onOpen  func()
	onSubmit func(values map[string]string)

	dialog *dialog.Component
	selection *selection.Feature
	inputs []*input.Input

	fieldBorder *borderlabel.Component
	fieldBorderSelected *borderlabel.Component
}

func Create() *Component {
	fieldBorder := borderlabel.Create().SetColor(theme.Current.Muted)
	fieldBorderSelected := borderlabel.Create().SetColor(theme.Current.Primary)

	return &Component{
		fields: []Field{},
		values: map[string]string{},

		dialog: dialog.Create().SetTitle("Map Editor"),
		selection: selection.Create(),

		fieldBorder: fieldBorder,
		fieldBorderSelected: fieldBorderSelected,
	}
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

	inputs := []*input.Input{}

	for range fields {
		input := input.New()
		
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
	c.dialog.OnClose(func() {
		c.DisableInputs()
	})

	return chain.Init(c.dialog.Init, c.InitRender)
}

func (c *Component) Dispose() tea.Cmd {
	return chain.Dispose(c.dialog.Dispose)
}

