package action

import (
	"fmt"
	"maps"

	"github.com/sidekick-coder/atlas/internal/template"
)

func (m *Manager) Execute(id string, context map[string]any, args ...template.EvaluateMapOptions) (map[string]any, error) {
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

	eo := template.NewEvaluateMapOptions()

	if len(args) > 0 {
		eo = args[0]
	}

	ctx := make(map[string]any)

	maps.Copy(ctx, context)

	if handlerId == "group" && options["actions"] != nil {
		ctx["actions"] = options["actions"]
	}

	if handlerId != "group" {
		opt, err := template.EvaluateMap(options, context, eo)

		if err != nil {
			return nil, err
		}

		maps.Copy(ctx, opt)
	}

	return def.Execute(ctx)

}
