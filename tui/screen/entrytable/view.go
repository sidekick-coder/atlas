package entrytable

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/table"
)

func (s *Screen) SetSize(width, height int) {
	s.width = width
	s.height = height
	s.table.SetSize(width-6, height)
	s.container.SetSize(width-4, height).SetMargin(0, 2, 0, 2).SetBorder("12")
}

func (s *Screen) LoadColumns() tea.Cmd {
	s.table.SetColumns([]table.Column{
		{Label: "ID", Field: "id", Width: 10},
		{Label: "Name", Field: "name", Width: 20},
		{Label: "Description", Field: "description", Width: 30},
	})

	return nil
}

func (s *Screen) Render() string {
	table := s.table.Render()

	s.container.SetContent(table)

	return s.container.Render()
}

