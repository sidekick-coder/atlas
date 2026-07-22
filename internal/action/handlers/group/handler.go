package group

import (
	"errors"
	"fmt"
	"log/slog"
	"maps"

	"github.com/sidekick-coder/atlas/internal/template"
	"github.com/sidekick-coder/atlas/internal/utils/maputil"
)

type GroupAction struct {
	ID      string
	Type    string
	Options map[string]any
}

type Handler struct {
	ExecuteAction func(id string, ctx map[string]any) (map[string]any, error)
}

func Create(executeAction func(id string, ctx map[string]any) (map[string]any, error)) Handler {
	return Handler{
		ExecuteAction: executeAction,
	}
}

func ParseGroupAction(payload any) (GroupAction, error) {
	action, ok := payload.(map[string]any)

	if !ok {
		return GroupAction{}, errors.New("invalid group action format")
	}

	ga := GroupAction{}

	typeVal, ok := action["type"].(string)

	if !ok {
		return ga, errors.New("missing or invalid 'type' in group action")
	}

	if id, ok := action["id"].(string); ok {
		ga.ID = id
	}

	ga.Type = typeVal

	ga.Options = maputil.Except(action, "type")

	return ga, nil
}

func (h Handler) Execute(ctx map[string]any) (map[string]any, error) {
	result := make(map[string]any)

	actions, ok := ctx["actions"].([]any)

	if !ok {
		return result, errors.New("invalid or missing 'actions' in context")
	}

	groupActions := make([]GroupAction, 0, len(actions))

	for index, action := range actions {
		ga, err := ParseGroupAction(action)

		if err != nil {
			return result, err
		}

		if ga.ID == "" {
			ga.ID = fmt.Sprintf("%d", index)
		}

		groupActions = append(groupActions, ga)
	}

	currentCtx := map[string]any{}

	maps.Copy(currentCtx, maputil.Except(ctx, "actions"))

	for _, ga := range groupActions {
		if ga.Type == "group" {
			return result, errors.New("nested group actions are not allowed")
		}


		maps.Copy(currentCtx, result)

		opts, err := template.EvaluateMap(maputil.Except(ga.Options, "id"), currentCtx)

		if err != nil {
			return result, err
		}

		maps.Copy(currentCtx, opts)

		ar, err := h.ExecuteAction(ga.Type, currentCtx)

		if err != nil {
			return result, err
		}

		result[ga.ID] = ar

		slog.Info("Current ctx", "ctx", currentCtx)

	}

	slog.Info("Group action executed", "result", result)

	result["is_group"] = true

	return result, nil
}
