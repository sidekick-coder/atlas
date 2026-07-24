package entrycontroller

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/app/program"
)

type UpdatedMsg struct {
	Path string
}

func updated(path string) tea.Cmd {
	return program.Command(UpdatedMsg{Path: path})
}
