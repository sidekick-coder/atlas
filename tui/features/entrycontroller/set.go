package entrycontroller

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/components/toast"
)

type SetMsg struct {
	Path  string
	Name  string
	Value string
}

type SetEndMsg struct {
	Path  string
	Name  string
	Value string
}

func Set(path, name, value string) tea.Cmd {
	return program.Command(SetMsg{
		Path:  path,
		Name:  name,
		Value: value,
	})
}

func setEnd(path, name, value string) tea.Cmd {
	return program.Command(SetEndMsg{
		Path:  path,
		Name:  name,
		Value: value,
	})
}

func (f *Feature) HandleSet(msg tea.Msg) tea.Cmd {
	if m, ok := msg.(SetMsg); ok {
		err := f.app.Syncer().One(m.Path) // Sync the entry before updating

		if err != nil {
			return toast.Error(fmt.Sprintf("Failed to sync entry: %v", err))
		}

		err = f.app.SetEntryMeta(m.Path, m.Name, fmt.Sprintf("%v", m.Value))

		if err != nil {
			return toast.Error(fmt.Sprintf("Failed to update entry meta: %v", err))
		}

		return tea.Batch(
			setEnd(m.Path, m.Name, m.Value),
			updated(m.Path),
		)
	}

	return nil
}
