package column

import (
	"log"
	"strings"

	"github.com/sidekick-coder/atlas/tui/features/selection"
)

type Column struct {
	Label string 
	Field string
	Width int
}

type Feature struct {
	columns []*Column
	sizes []int
	width int
	height int

	Selection selection.Feature
}

func Create() *Feature {
	return &Feature{
		columns: []*Column{},
		sizes: []int{},
		width: 100,

		Selection: *selection.Create(),
	}
}

func (f *Feature) SetSize(w,h int) {
	f.width = w 
	f.height = h
}

func (f *Feature) GetColumns() []*Column {
	return f.columns
}

func (f *Feature) GetColumn(index int) (*Column, bool) {
	if index < 0 || index >= len(f.columns) {
		return nil, false
	}

	return f.columns[index], true
}

func (f *Feature) GetColumnSelected() (*Column, bool) {
	index := f.Selection.GetCursor()

	if index < 0 || index >= len(f.columns) {
		return nil, false
	}

	return f.columns[index], true
}

func (f *Feature) SetColumns(columns []*Column) {
	f.columns = columns

	f.Selection.SetTotal(len(f.columns))

	remaningWidth := f.width

	f.sizes = make([]int, len(columns))
	withWidthCount := 0

	for i, column := range columns {
		if column.Width > 0 {
			f.sizes[i] = column.Width
			remaningWidth -= column.Width
			withWidthCount++
		}
	}

	wihoutWidthWidth := remaningWidth / (len(columns) - withWidthCount)

	for i, column := range columns {
		if column.Width == 0 {
			f.sizes[i] = wihoutWidthWidth
		}
	}

	log.Printf("Column sizes: %v", f.sizes)
}

func (f *Feature) GetColumnIndex(column *Column) int {
	for i, col := range f.columns {
		if col.Field == column.Field {
			return i
		}
	}

	return -1
}

func (f *Feature) GetColumnSizes() []int {
	return f.sizes
}

func (f *Feature) ParseColumnText(column *Column,  text string) string {
	colIndex := f.GetColumnIndex(column)

	if colIndex == -1 {
		return text
	}

	sizes := f.GetColumnSizes()

	colw := sizes[colIndex]

	colw -= 2 // padding

	var result string

	if len(text) > colw {
		result = text[:colw-3] + "..."
	}

	if len(text) < colw {
		padding := colw - len(text)
		result = text + strings.Repeat("\u00A0", padding)
	}

	// add padding 
	result = strings.Repeat("\u00A0", 1) + result + strings.Repeat("\u00A0", 1)

	return result
}


