package column

import (
	"fmt"
	"strconv"
)

func (f *Feature) UpdateColumnByIndex(index int, payload Column) error {
	column, ok := f.GetColumn(index)

	if !ok {
		return fmt.Errorf("column index %d out of range", index)
	}

	newColumn := &Column{
		Field: payload.Field,
		Label: payload.Label,
		Width: column.Width,
	}

	f.columns[index] = newColumn

	return nil
}

func (f *Feature) ParseMapToColumn(payload map[string]string) (Column, error) {
	field, ok := payload["field"]

	if !ok {
		return Column{}, fmt.Errorf("field is required")
	}

	label, ok := payload["label"]

	if !ok {
		return Column{}, fmt.Errorf("label is required")
	}

	column := Column{
		Field: field,
		Label: label,
	}

	widthStr, ok := payload["width"]

	if !ok {
		return column, nil
	}

	if widthStr == "auto" {
		return column, nil
	}

	width, err := strconv.Atoi(widthStr)

	if err != nil {
		return Column{}, fmt.Errorf("invalid width: %v", err)
	}

	column.Width = width

	return column, nil
}

func (f *Feature) UpdateSelectedColumn(payload Column) error {
	index := f.Selection.GetCursor()

	if index < 0 || index >= len(f.columns) {
		return fmt.Errorf("no column selected")
	}

	return f.UpdateColumnByIndex(index, payload)
}

func (f *Feature) AddColumn(payload Column) {
	newColumns := append(f.columns, &payload)

	f.SetColumns(newColumns)
}
