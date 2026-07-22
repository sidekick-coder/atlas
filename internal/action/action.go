package action

import (
	"fmt"
	"maps"

	"github.com/sidekick-coder/atlas/internal/template"
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
	return &Manager{
		data:        make(map[string]ActionData),
		definitions: make(map[string]ActionHandler),
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

func (m *Manager) AddDefinition(id string, execute func(map[string]any) (map[string]any, error)) {
	m.definitions[id] = ActionHandler{
		ID:      id,
		Execute: execute,
	}
}

func (m *Manager) Execute(id string, context map[string]any) (map[string]any, error) {
	data, exists := m.data[id]

	if !exists {
		return nil, fmt.Errorf("action with ID %s not found", id)
	}

	def, exists := m.definitions[data.Type]

	if !exists {
		return nil, fmt.Errorf("action definition for type %s not found", data.Type)
	}

	opt, err := template.EvaluateMap(data.Options, context)

	if err != nil {
		return nil, err
	}

	ctx := make(map[string]any)

	maps.Copy(ctx, context)
	maps.Copy(ctx, opt)

	return def.Execute(ctx)
}

func (m *Manager) List() ([]ActionData, error) {
	ad := make([]ActionData, 0, len(m.data))

	for _, action := range m.data {
		ad = append(ad, action)
	}

	return ad, nil
}
