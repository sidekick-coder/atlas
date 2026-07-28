package entrylist

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/internal/template"
	"github.com/sidekick-coder/atlas/internal/utils/maputil"
	"github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (c *Component) SetSize(width, height int) {
	c.list.SetSize(width, height)

	limit := 20

	limit = max(limit, height-2)

	c.loader.SetLimit(limit)

	c.loader.Load()
	c.LoadItems()
}

func Color(payload any, color string) string {
	value := fmt.Sprintf("%v", payload)
	c := color

	pallete := map[string]string{
		"primary":   theme.Current.Primary,
		"secondary": theme.Current.Secondary,
		"muted":     theme.Current.Muted,
		"error":     theme.Current.Error,
		"success":   theme.Current.Success,
		"warning":   theme.Current.Warning,
	}

	if v, ok := pallete[color]; ok {
		c = v
	}

	return lipgloss.NewStyle().Foreground(lipgloss.Color(c)).Render(value)
}

func Concat(payload ...string) string {
	return strings.Join(payload, "")
}

func Width(payload any, width int) string {
	value := fmt.Sprintf("%v", payload)

	return lipgloss.NewStyle().Width(width).Render(value)
}

func Align(payload any, alignment string) string {
	value := fmt.Sprintf("%v", payload)

	switch alignment {
	case "left":
		return lipgloss.NewStyle().Align(lipgloss.Left).Render(value)
	case "right":
		return lipgloss.NewStyle().Align(lipgloss.Right).Render(value)
	case "center":
		return lipgloss.NewStyle().Align(lipgloss.Center).Render(value)
	default:
		return value
	}
}

func style(value any, payload string) string {
	s := lipgloss.NewStyle()
	options := maputil.FromString(payload)

	// parse options...
	if w, ok := maputil.GetInt(options, "width"); ok {
		s = s.Width(w)
	}

	if a, ok := maputil.GetString(options, "align"); ok {
		switch a {
		case "left":
			s = s.Align(lipgloss.Left)
		case "right":
			s = s.Align(lipgloss.Right)
		case "center":
			s = s.Align(lipgloss.Center)
		}
	}

	return s.Render(fmt.Sprint(value))
}

func ColorMap(colorMap map[string]any, args ...string) string {
	colors := maputil.String(colorMap)

	key := ""
	value := ""

	if len(args) > 0 {
		key = args[0]
		value = args[0]
	}

	if len(args) > 1 {
		value = args[1]
	}

	color := theme.Current.Foreground

	if c, ok := colors[key]; ok {
		color = c
	}

	return Color(value, color)
}

func (c *Component) LoadItems() tea.Cmd {
	var items []string

	config := program.GetConfig()

	labelKey := ""
	labelTemplate := ""

	if l, ok := c.props["label_key"].(string); ok {
		labelKey = l
	}

	if lt, ok := c.props["label_template"].(string); ok {
		labelTemplate = lt
	}

	muted := theme.BaseStyle().Foreground(lipgloss.Color(theme.Current.Muted))

	for index, entry := range c.loader.GetEntries() {
		ctx := entry.ToMap()

		value := muted.Render("N/A")

		ctx["vars"] = config.GetMap("vars")

		if c.selection.IsSelected(index) {
			value = muted.Background(lipgloss.Color(theme.Current.Selection)).Render("N/A")
		}

		if labelKey != "" {
			v, ok := maputil.GetString(ctx, labelKey)

			if ok {
				value = v
			}
		}

		if labelTemplate != "" {
			v, err := template.Render(labelTemplate, ctx, template.RenderOptions{
				AllowMissingKeys: true,
				FuncMap: map[string]any{
					"color":     Color,
					"color_map": ColorMap,
					"concat":    Concat,
					"style":     style,
				},
			})

			if err == nil {
				value = v
			}

			if err != nil {
				value = muted.Render("N/A")
			}
		}

		items = append(items, value)
	}

	c.list.SetItems(items)

	return nil
}

func (c *Component) Render() string {

	return c.list.Render()
}
