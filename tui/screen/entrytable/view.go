package entrytable

import (
	"log"
	"maps"
	"strconv"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/table"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (s *Screen) SetSize(width, height int) {
	s.width = width
	s.height = height
	s.table.SetSize(width-10, height)
	s.container.SetSize(width-4, height).SetMargin(0, 2, 0, 2).SetBorder(theme.Current.Primary).SetPadding(0, 2, 0, 2)
}

func (s *Screen) LoadColumns() tea.Cmd {
	s.table.SetColumns([]table.Column{
		{Label: "ID", Field: "id", Width: 10},
		{Label: "Name", Field: "basename", Width: 120},
		{Label: "Path", Field: "path"},
	})

	return nil
}

func (s *Screen) Render() string {

	var items []table.Item

	for _, entry := range s.loader.GetEntries() {
		values := map[string]string{}

		maps.Copy(values, entry.Metas)

		values["id"] = strconv.FormatInt(entry.ID, 10)
		values["path"] = entry.Path

		log.Printf("Rendering entry: %v", entry)

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
