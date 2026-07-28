package trigger

import (
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/utils"
)

type Trigger struct {
	ID        string
	ContextID string
	Test      func(km config.Keymap) bool
	Metas     map[string]any
}

func Create() Trigger {
	id, err := utils.CreateID()

	if err != nil {
		panic(err)
	}

	return Trigger{
		ID:        id,
		ContextID: "",
		Test:      func(km config.Keymap) bool { return false },
		Metas: map[string]any{},
	}
}

func (t *Trigger) Set(key string, value any) {
	t.Metas[key] = value
}
