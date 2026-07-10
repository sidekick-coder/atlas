package entrytable

import (
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
		{Label: "Path", Field: "path"},
	})

	return nil
}

func (s *Screen) Render() string {

	var items []table.Item

	for _, entry := range s.loader.GetEntries() {
		items = append(items, table.Item{
			Values: map[string]string{
				"path": entry.Path,
			},
		})
	}

	s.table.SetItems(items)

	table := s.table.Render()

	s.container.SetContent(table)

	return s.container.Render()
}
