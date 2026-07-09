package entry

import (
	"github.com/sidekick-coder/atlas/internal/config"
)

func (s *Screen) GetEntryKeymapGroups() []string {
	groups := []string{}
	entry, selected := s.List.SelectedEntry()

	if !selected {
		return groups
	}

	entryMetas, err := s.App.EntryMetaRepo().ListByEntryPath(entry.Path)

	if err != nil {
		return groups
	}

	for _, meta := range entryMetas {
		groups = append(groups, meta.Name+"="+meta.Value)
	}

	return groups
}

func (s *Screen) GetUserKeymaps() []config.Keymap {
	groups := []string{"entry_list"}

	groups = append(groups, s.GetEntryKeymapGroups()...)

	keymaps := []config.Keymap{}

	for _, group := range groups {
		keymaps = append(keymaps, s.App.Config().GetKeymapsByGroup(group)...)
	}

	return keymaps
}
