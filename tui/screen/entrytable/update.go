package entrytable

import (
	"fmt"
	"strconv"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/table"
	"github.com/sidekick-coder/atlas/tui/messages"
)

func (s *Screen) LoadDefaultColumn() tea.Cmd {
	columns := []*table.Column{}

	columns = append(columns, &table.Column{Label: "ID", Field: "id", Width: 10})
	columns = append(columns, &table.Column{Label: "Name", Field: "basename", Width: 120})
	columns = append(columns, &table.Column{Label: "Path", Field: "path"})

	s.table.SetColumns(columns)

	return nil
}

func (s *Screen) ParseOptionColumn(option map[string]any) (*table.Column, error) {
	column := &table.Column{}

	l, ok := option["label"].(string)

	if !ok {
		return nil, fmt.Errorf("missing or invalid 'label' in column option: %v", option)
	}

	f, ok := option["field"].(string)

	if !ok {
		return nil, fmt.Errorf("missing or invalid 'field' in column option: %v", option)
	}

	column.Label = l
	column.Field = f

	if w, ok := option["width"].(string); ok {
		width, err := strconv.Atoi(w)

		if err != nil {
			return nil, fmt.Errorf("invalid 'width' in column option: %v", option)
		}

		column.Width = width
	}

	return column, nil
}

func (s *Screen) LoadColumns() tea.Cmd {
	columns := []*table.Column{}

	oc, ok := s.options["columns"].([]any)


	if !ok {
		return s.LoadDefaultColumn()
	}

	for _, c := range oc {
		if colMap, ok := c.(map[string]any); ok {
			column, err := s.ParseOptionColumn(colMap)

			if err != nil {
				return messages.ToastErrorCmd(fmt.Sprintf("Failed to parse column option: %v", err))
			}

			columns = append(columns, column)
		}

	}

	s.table.SetColumns(columns)

	return nil
}
