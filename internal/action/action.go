package action

import "fmt"

type ActionData struct {
	ID      string
	Type    string
	Options map[string]any
}

type ActionDefinition struct {
	ID      string
	Execute func(map[string]string) (map[string]any, error)
}

type Manager struct {
	data        map[string]ActionData
	definitions map[string]ActionDefinition
}

func Create() *Manager {
	return &Manager{
		data:        make(map[string]ActionData),
		definitions: make(map[string]ActionDefinition),
	}
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

func (m *Manager) AddDefinition(def ActionDefinition) {
	m.definitions[def.ID] = def
}

func (m *Manager) Execute(id string, ctx map[string]string) (map[string]any, error) {
	data, exists := m.data[id]

	if !exists {
		return nil, fmt.Errorf("action with ID %s not found", id)
	}

	def, exists := m.definitions[data.Type]

	if !exists {
		return nil, fmt.Errorf("action definition for type %s not found", data.Type)
	}

	return def.Execute(ctx)
}

func (m *Manager) List() ([]ActionData, error) {
	ad := make([]ActionData, 0, len(m.data))

	for _, action := range m.data {
		ad = append(ad, action)
	}

	return ad, nil
}
