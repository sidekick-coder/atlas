package input

import (
	"image/color"
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

type Input struct {
	Title    string
	Color    color.Color
	Buf      []rune
	Cursor   int
	Width    int
	Height   int
	ScreenWidth int
	ScreenHeight int
	Active   bool
	callback func(value string) tea.Cmd
}

func New() *Input {
	return &Input{Color: lipgloss.Color("12"), Width: 80, Height: 3}
}

func (i *Input) SetTitle(title string) *Input {
	i.Title = title
	return i
}

func (i *Input) SetColor(c color.Color) *Input {
	i.Color = c
	return i
}

func (i *Input) OnSubmit(fn func(string) tea.Cmd) *Input {
	i.callback = fn
	return i
}

func (i *Input) SetSize(width, height int) {
	i.Width = width
	i.Height = height
}

func (i *Input) SetScreenSize(width, height int) {
	i.ScreenWidth = width
	i.ScreenHeight = height
}

func (i *Input) Open(initialValue string) {
	i.Buf = []rune(initialValue)
	i.Cursor = len(i.Buf)
	i.Active = true
}

func (i *Input) Close() {
	i.Active = false
	i.Buf = nil
	i.Cursor = 0
}
