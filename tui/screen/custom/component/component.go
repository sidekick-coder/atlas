package component

import (
	"fmt"

	"github.com/sidekick-coder/atlas/internal/utils"
	"github.com/sidekick-coder/atlas/internal/utils/maputil"
)

type Component struct {
	Type string
	Cols int
	Rows int
	X    int
	Y    int
	Definition Definition
}

func Create() *Component {
	return &Component{
		Type: "unknown",
		Cols: 0,
		Rows: 0,
		X:    0,
		Y:    0,
	}
}

func CreateFromMap(payload any) (*Component, error) {
	component := Create()

	cm, ok := payload.(map[string]any)

	if !ok {
		return nil, fmt.Errorf("invalid component type: %T", payload)
	}

	t, ok := cm["type"].(string)

	if !ok {
		return nil, fmt.Errorf("invalid component type: %T", cm["type"])
	}

	definition, ok := Definitions[t]

	if !ok {
		return nil, fmt.Errorf("component definition not found: %s", t)
	}

	component.Type = t
	component.Definition = definition(DefinitionPayload{
		Width:   component.Cols,
		Height:  component.Rows,
		Options: maputil.Except(cm, "type", "cols", "rows", "x", "y"),
	})


	if cols, ok := utils.ParseInt(cm["cols"]); ok {
		component.Cols = cols
	}

	if rows, ok := utils.ParseInt(cm["rows"]); ok {
		component.Rows = rows
	}

	if x, ok := utils.ParseInt(cm["x"]); ok {
		component.X = x
	}

	if y, ok := utils.ParseInt(cm["y"]); ok {
		component.Y = y
	}

	return component, nil
}
