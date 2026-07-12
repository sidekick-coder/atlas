package footer

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

func (f *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(
		msg,
		chain.OnKey(f.HandleBindings),
		f.dialog.Update,
	)
}
