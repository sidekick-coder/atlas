package screen

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

func (f *Feature) HandleMessages(msg tea.Msg) tea.Cmd {
	if wm, ok := msg.(tea.WindowSizeMsg); ok {
		f.windowWidth = wm.Width
		f.windowHeight = wm.Height

		height := f.windowHeight - 7 // footer, hedaer,tabs height
		width := f.windowWidth

		return func() tea.Msg {
			return SizeMsg{
				Width:  width,
				Height: height,
			}
		}
	}

	if s, ok := f.GetCurrent(); ok {
		return s.Update(msg)
	}

	return nil

}

func (f *Feature) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(
		msg,
		f.HandleMessages,
		chain.OnKey(f.HandleBinding),
	)
}
