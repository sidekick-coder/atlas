package form

import (
	"fmt"

	"github.com/sidekick-coder/atlas/tui/components/input"
)

type Field struct {
	Label    string
	FielName string
}

func CreateFieldFromMap(payload any) (Field, error) {
	f := Field{}

	m, ok := payload.(map[string]any)

	if !ok {
		return f, fmt.Errorf("invalid field type: %T", payload)
	}

	if label, ok := m["label"].(string); ok {
		f.Label = label
	}

	if name, ok := m["name"].(string); ok {
		f.FielName = name
	}

	return f, nil
}

func CreateFieldsFromArray(payload any) ([]Field, error) {
	fields := []Field{}

	array, ok := payload.([]any)

	if !ok {
		return fields, fmt.Errorf("invalid fields type: %T", payload)
	}

	for _, item := range array {
		field, err := CreateFieldFromMap(item)

		if err != nil {
			return fields, err
		}

		fields = append(fields, field)
	}

	return fields, nil
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
		input := input.Create()
		input.SetWidth(c.width - 4) // 4 padding

		inputs = append(inputs, input)
	}

	c.inputs = inputs
}

