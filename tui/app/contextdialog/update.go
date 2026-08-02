package contextdialog

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

func (f *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(
		msg,
		chain.OnKey(f.HandleBindings),
		f.dialog.Update,
		f.HandleMessages,
	)
}

func (f *Component) HandleMessages(msg tea.Msg) tea.Cmd {
	if m, ok := msg.(tea.WindowSizeMsg); ok {
		width := min(200, m.Width - 10)
		height := min(800, m.Height - 10)

		f.dialog.SetSize(width, height)
		f.kv.SetSize(width-2, height-2)
	}
	
	return nil
}
