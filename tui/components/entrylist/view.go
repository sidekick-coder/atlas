package entrylist

import (
	"fmt"

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
	}

	if v, ok := pallete[color]; ok {
		c = v
	}

	return lipgloss.NewStyle().Foreground(lipgloss.Color(c)).Render(value)
}

func ColorMap(payload any, colorMap map[string]any) string {
	colors := maputil.String(colorMap)

	value := fmt.Sprintf("%v", payload)
	color := theme.Current.Foreground

	if c, ok := colors[value]; ok {
		color = c
	}

	return Color(payload, color)
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
