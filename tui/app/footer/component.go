package footer

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/helpdialog"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

type Component struct {
	width    int
	dialog   *helpdialog.Component
}

func Create() *Component {
	return &Component{
		dialog: helpdialog.Create(),
	}
}

func (f *Component) SetWidth(width int) {
	f.width = width
}


func (f *Component) View() string {
	return f.Render()
}

func (f *Component) Init() tea.Cmd {
	return chain.Init(f.LoadBindings, f.dialog.Init)
}
