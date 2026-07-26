package entrycontroller

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/app/program"
)

type UpdateMsg struct {
	Path string
	Values map[string]string
}

type UpdatedMsg struct {
	Path string
}

func updated(path string) tea.Cmd {
	return program.Command(UpdatedMsg{Path: path})
}

func Update() tea.Cmd {
	return updated("")
}
