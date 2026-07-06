package box

import (
	lipgloss "charm.land/lipgloss/v2"
)

type Box struct {
	Width  int
	Height int
	Content   string
	Style lipgloss.Style
}

func Create() *Box {
	return &Box{
		Width:  100,
		Height: 100,
		Content:   "",
		Style: lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("12")),
	}
}

func (b *Box) SetSize(width, height int) *Box {
	b.Width = width
	b.Height = height
	return b
}

func (b *Box) SetContent(content string) *Box {
	b.Content = content
	return b
}


func (b *Box) Render() string {
	border := b.Style.
		Width(b.Width - 4).
		Height(b.Height - 4).
		Margin(0, 2).
		BorderForeground(lipgloss.Color("12"))

	content := ""

	if b.Content != "" {
		content = b.Content
	}

	return border.Render(content)
}
