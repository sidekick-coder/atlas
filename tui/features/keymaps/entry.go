package keymaps

import (
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/models"
)

func CreateEntryTrigger(entry models.Entry) Trigger {
	t := CreateTrigger()

	em := entry.ToMap()

	t.Test = func(km config.Keymap) bool {
		for k, v := range em {
			if kmv, ok := km.Options[k]; ok {
				return kmv == v
			} 
		}

		return false
	}

	return t
}
