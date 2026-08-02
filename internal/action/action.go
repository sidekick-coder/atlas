package action

import (

	"github.com/sidekick-coder/atlas/internal/action/handlers/group"
	"github.com/sidekick-coder/atlas/internal/action/handlers/shell"
)

type ActionData struct {
	ID      string
	Type    string
	Options map[string]any
}

type ActionHandler struct {
	ID      string
	Execute func(map[string]any) (map[string]any, error)
}

type Manager struct {
	data        map[string]ActionData
	definitions map[string]ActionHandler
}

func Create() *Manager {
	m := &Manager{
		data:        make(map[string]ActionData),
		definitions: make(map[string]ActionHandler),
	}

	g := group.Create(m.Execute)

	m.AddDefinition("group", g.Execute)

	s := shell.Create()

	m.AddDefinition("shell", s.Execute)

	return m
}

func (m *Manager) Add(id string, actionType string, options map[string]any) ActionData {
	ad := ActionData{
		ID:      id,
		Type:    actionType,
		Options: options,
	}

	m.data[id] = ad

	return ad
}

func (m *Manager) AddDefinition(id string, execute func(map[string]any) (map[string]any, error)) {
	m.definitions[id] = ActionHandler{
		ID:      id,
		Execute: execute,
	}
}

func (m *Manager) List() ([]ActionData, error) {
	ad := make([]ActionData, 0, len(m.data))

	for _, action := range m.data {
		ad = append(ad, action)
	}

	return ad, nil
}
