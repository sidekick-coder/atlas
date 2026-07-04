package screen

import (
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

// PlaceholderScreen is a simple stub used for screens not yet implemented.
type PlaceholderScreen struct {
	title  string
	icon   string
	width  int
	height int
}

func NewPlaceholderScreen(icon, title string) *PlaceholderScreen {
	return &PlaceholderScreen{title: title, icon: icon}
}

func (s *PlaceholderScreen) Title() string { return s.icon + " " + s.title }

func (s *PlaceholderScreen) Init() tea.Cmd { return nil }

func (s *PlaceholderScreen) Update(msg tea.Msg) tea.Cmd {
	if ws, ok := msg.(tea.WindowSizeMsg); ok {
		s.width = ws.Width
		s.height = ws.Height
	}
	return nil
}

func (s *PlaceholderScreen) Render() string {
	label := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Render(s.icon + "  " + s.title + " — coming soon")

	return lipgloss.NewStyle().
		Width(s.width).
		Height(s.height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(label)
}
