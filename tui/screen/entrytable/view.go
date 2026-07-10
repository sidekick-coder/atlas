package entrytable

import (
	"maps"
	"strconv"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/table"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (s *Screen) SetSize(width, height int) {
	s.width = width
	s.height = height
	s.table.SetSize(width-8, height)
	s.container.SetSize(width-6, height).SetMargin(0, 2, 0, 2).SetBorder(theme.Current.Primary)
}

func (s *Screen) LoadColumns() tea.Cmd {
	columns := []*table.Column{}

	columns = append(columns, &table.Column{Label: "ID", Field: "id", Width: 10})
	columns = append(columns, &table.Column{Label: "Name", Field: "basename", Width: 120})
	columns = append(columns, &table.Column{Label: "Path", Field: "path"})

	s.table.SetColumns(columns)

	return nil
}

func (s *Screen) Render() string {

	var items []table.Item

	for _, entry := range s.loader.GetEntries() {
		values := map[string]string{}

		maps.Copy(values, entry.Metas)

		values["id"] = strconv.FormatInt(entry.ID, 10)
		values["path"] = entry.Path

		item := table.Item{
			Values: values,
		}

		items = append(items, item)
	}

	s.table.SetItems(items)

	table := s.table.Render()

	s.container.SetContent(table)

	return s.container.Render()
}
