package helpdialog

import (
	"fmt"
	"maps"
	"slices"
	"sort"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/app/screen"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/features/layer"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (c *Component) LoadSize() {
	width := max(50, layer.ScreenWidth/2)
	height := max(10, layer.ScreenHeight-10)

	c.viewport.SetSize(width, height-2) // 2 padding
	c.dialog.SetSize(width, height)
}

func (c *Component) HandleView(msg tea.Msg) tea.Cmd {
	if _, ok := msg.(screen.SizeMsg); ok {
		c.LoadSize()
	}

	return nil
}

func prefix(tag string) string {
	if i := strings.IndexByte(tag, ':'); i != -1 {
		return tag[:i]
	}
	return tag
}

func (c *Component) Render() string {
	tagStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Current.Primary))

	textStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Current.Foreground))

	descStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Current.Muted))

	var parts []string

	groups := make(map[string][]key.Binding)

	for _, b := range key.GetBindings() {
		tags := b.GetTags()

		for _, tag := range tags {
			if tag == "help-dialog=false" {
				continue
			}

			groups[tag] = append(groups[tag], b)
		}
	}

	tags := slices.Collect(maps.Keys(groups))
	priority := map[string]int{
		"global":    0,
		"screen":    1,
		"component": 2,
	}

	sort.Slice(tags, func(i, j int) bool {
		pi := priority[prefix(tags[i])]
		pj := priority[prefix(tags[j])]

		if pi != pj {
			return pi < pj
		}

		// Same prefix: sort alphabetically.
		return tags[i] < tags[j]

	})

	for _, tag := range tags {
		parts = append(parts, tagStyle.Render(tag))
		bindings := groups[tag]

		if len(bindings) == 0 {
			continue
		}

		if tag == "screen" {
			parts = append(parts, descStyle.Render("Use <leader> + number to switch screens"))
		}

		for _, b := range bindings {
			d := b.GetDescription()
			k := ""

			for _, bk := range b.GetKeys() {
				if k != "" {
					k += ", "
				}

				k += bk.String()
			}

			part := fmt.Sprintf("%s %s", textStyle.Render(k), descStyle.Render(d))

			parts = append(parts, part)
		}
	}

	content := lipgloss.JoinVertical(lipgloss.Left, parts...)

	content = c.viewport.SetContent(content).Render()

	return content
}
