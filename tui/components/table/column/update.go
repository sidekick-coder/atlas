package column

import (
	"fmt"
	"strconv"
)

func (f *Feature) UpdateColumnByIndex(index int, payload Column) error {
	_, ok := f.GetColumn(index)

	if !ok {
		return fmt.Errorf("column index %d out of range", index)
	}


	newColumn := &Column{
		Field: payload.Field,
		Label: payload.Label,
		Width: payload.Width,
	}

	f.columns[index] = newColumn

	f.SetColumns(f.columns)

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

func (f *Feature) RemoveColumnByIndex(index int) error {
	if index < 0 || index >= len(f.columns) {
		return fmt.Errorf("column index %d out of range", index)
	}

	newColumns := append(f.columns[:index], f.columns[index+1:]...)

	f.SetColumns(newColumns)

	if f.Selection.GetCursor() >= len(f.columns) {
		f.Selection.SetCursor(len(f.columns) - 1)
	}

	return nil
}

func (f *Feature) RemoveSelectedColumn() error {
	index := f.Selection.GetCursor()

	if index < 0 || index >= len(f.columns) {
		return fmt.Errorf("no column selected")
	}

	return f.RemoveColumnByIndex(index)
}

func (f *Feature) AddColumn(payload Column) {
	newColumns := append(f.columns, &payload)

	f.SetColumns(newColumns)
}
