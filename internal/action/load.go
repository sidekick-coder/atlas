package action

import (
	"github.com/sidekick-coder/atlas/internal/config"
)

func (m *Manager) LoadConfigActions(config *config.Config) error {
	actions, err := config.GetActions()

	if err != nil {
		return err
	}

	for _, action := range actions {
		m.Add(action.ID, action.Type, action.Options)
	}

	return nil
}
