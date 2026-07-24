package theme

import "charm.land/lipgloss/v2"

type Theme struct {
	Primary   string
	Secondary string
	Accent    string

	Foreground string
	Muted      string
	Background string

	Success string
	Warning string
	Error   string
	Info    string

	Border      string
	Selection   string
	Active      string
	Placeholder string
}

var Current = CatppuccinMocha

// var Current = Theme{
// 	Primary:   "#1E90FF",
// 	Secondary: "#FF69B4",
// 	Accent:    "#32CD32",
//
// 	Foreground: "#FFFFFF",
// 	Muted:      "#A9A9A9",
// 	Background: "#000000",
//
// 	Success: "#00FF00",
// 	Warning: "#FFFF00",
// 	Error:   "#FF0000",
// 	Info:    "#00FFFF",
//
// 	Border:      "#808080",
// 	Selection:   "#FFD700",
// 	Highlight:   "#FFA500",
// 	Placeholder: "#696969",
// }

func BaseStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(lipgloss.Color(Current.Background)).
		Foreground(lipgloss.Color(Current.Foreground)).
		BorderBackground(lipgloss.Color(Current.Background)).
		BorderForeground(lipgloss.Color(Current.Border)).
		MarginBackground(lipgloss.Color(Current.Background))
}

func SetTheme(theme Theme) {
	Current = theme
}
