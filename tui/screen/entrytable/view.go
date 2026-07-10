package entrytable

import (
	"maps"
	"strconv"

	"github.com/sidekick-coder/atlas/tui/components/table"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (s *Screen) SetSize(width, height int) {
	s.width = width
	s.height = height
	s.table.SetSize(width-6, height)
	s.container.SetSize(width-4, height).SetMargin(0, 2, 0, 2).SetBorder(theme.Current.Primary)
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
