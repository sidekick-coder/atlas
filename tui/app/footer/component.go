package footer

import (
	"path/filepath"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/components/helpdialog"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

type Component struct {
	width  int
	label  string
	dialog *helpdialog.Component
}

func Create(app *app.App) *Component {
	config := app.Config()

	path, _ := config.Get("workspace.path")

	label := filepath.Base(path)

	if t, ok := config.Get("tui.title"); ok {
		label = t
	}

	return &Component{
		dialog: helpdialog.Create(),
		label:  label,
	}
}

func (f *Component) SetWidth(width int) {
	f.width = width
}

func (f *Component) SetLabel(label string) {
	f.label = label
}

func (f *Component) View() string {
	return f.Render()
}

func (f *Component) Init() tea.Cmd {
	return chain.Init(f.LoadBindings, f.dialog.Init)
}
