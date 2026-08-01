package form

import (
	"github.com/sidekick-coder/atlas/tui/components/form/field"
)

type FieldData = field.Data
type FieldDefinition = field.Definition
type Field = field.Field

func (c *Component) AddField(data field.Data) {
	field := field.Field{}

	field.Data = data

	if defFn, ok := c.registry.Get(data.Type); ok {
		def := defFn()
		def.Resize(c.width, c.height)
		def.SetProps(data.Options)
		c.focus.Add(def)

		field.Definition = def
	}

	c.fields = append(c.fields, field)
}

func (c *Component) SetFields(payload []field.Data) {
	c.focus.Clear()
	c.fields = []field.Field{}

	for _, data := range payload {
		c.AddField(data)
	}
}
