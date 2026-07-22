package action

import (
	"fmt"
	"maps"

	"github.com/sidekick-coder/atlas/internal/action/handlers/group"
	"github.com/sidekick-coder/atlas/internal/action/handlers/shell"
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

func (m *Manager) Execute(id string, context map[string]any) (map[string]any, error) {
	handlerId := id
	options := make(map[string]any)

	data, exists := m.data[id]

	if exists {
		handlerId = data.Type
		options = data.Options
	}

	def, exists := m.definitions[handlerId]

	if !exists {
		return nil, fmt.Errorf("action handler %s not found", handlerId)
	}

	ctx := make(map[string]any)

	maps.Copy(ctx, context)

	if handlerId == "group" {
		ctx["actions"] = options["actions"]
	}

	if handlerId != "group" {
		opt, err := template.EvaluateMap(options, context)

		if err != nil {
			return nil, err
		}

		maps.Copy(ctx, opt)
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
