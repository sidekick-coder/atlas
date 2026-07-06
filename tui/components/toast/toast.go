package toast

import (
	lipgloss "charm.land/lipgloss/v2"
	"image/color"
)

type Component struct {
	Title        string
	Content      string
	Color        color.Color
	Width        int
	Height       int
	ScreenWidth  int
	ScreenHeight int
	Active       bool
}

func New() *Component {
	return &Component{
		Title:   "Toast",
		Content: "This is a toast message.",
		Color:   lipgloss.Color("12"),
		Width:   80,
		Height:  3,
	}
}

func (i *Component) SetTitle(title string) *Component {
	i.Title = title
	return i
}

func (i *Component) SetContent(content string) *Component {
	i.Content = content
	return i
}

func (i *Component) SetColor(c string) *Component {
	i.Color = lipgloss.Color(c)
	return i
}

func (i *Component) SetSize(width, height int) {
	i.Width = width
	i.Height = height
}

func (i *Component) SetScreenSize(width, height int) {
	i.ScreenWidth = width
	i.ScreenHeight = height
}

func (i *Component) SetActive(active bool) {
	i.Active = active
}
