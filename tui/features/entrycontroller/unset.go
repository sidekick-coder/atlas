package entrycontroller

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/components/toast"
)

type UnsetMsg struct {
	Path  string
	Name  string
}

type UnsetEndMsg struct {
	Path  string
	Name  string
}

func Unset(path, name string) tea.Cmd {
	return program.Command(UnsetMsg{
		Path:  path,
		Name:  name,
	})
}

func unsetEnd(path, name string) tea.Cmd {
	return program.Command(UnsetEndMsg{
		Path:  path,
		Name:  name,
	})
}

func (f *Feature) HandleUnset(msg tea.Msg) tea.Cmd {
	if m, ok := msg.(UnsetMsg); ok {
		err := f.app.Syncer().One(m.Path) // Sync the entry before updating

		if err != nil {
			return toast.Error(fmt.Sprintf("Failed to sync entry: %v", err))
		}

		err = f.app.UnsetEntryMeta(m.Path, m.Name)

		if err != nil {
			return toast.Error(fmt.Sprintf("Failed to update entry meta: %v", err))
		}

		return tea.Batch(
			unsetEnd(m.Path, m.Name),
			updated(m.Path),
		)

	}

	return nil
}
