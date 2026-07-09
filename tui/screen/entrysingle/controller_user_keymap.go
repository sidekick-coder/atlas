package entrysingle

import "github.com/sidekick-coder/atlas/internal/config"

func (s *Screen) GetEntryKeymapGroups() []string {
	groups := []string{}

	for _, meta := range s.EntryMetaComponent.Metas {
		groups = append(groups, meta.Name+"="+meta.Value)
	}

	return groups
}

func (s *Screen) GetUserKeymaps() []config.Keymap {
	groups := []string{"entry_single"}

	groups = append(groups, s.GetEntryKeymapGroups()...)

	keymaps := []config.Keymap{}

	for _, group := range groups {
		keymaps = append(keymaps, s.App.Config().GetKeymapsByGroup(group)...)
	}

	return keymaps
}

