package group

import (
	"errors"
	"fmt"
	"maps"

	"github.com/sidekick-coder/atlas/internal/utils/maputil"
)

type GroupAction struct {
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

	for _, action := range actions {
		ga, err := ParseGroupAction(action)

		if err != nil {
			return result, err
		}

		groupActions = append(groupActions, ga)
	}

	currentCtx := map[string]any{}

	for index, ga := range groupActions {
		maps.Copy(currentCtx, ga.Options)
		maps.Copy(currentCtx, ctx)

		actionResult, err := h.ExecuteAction(ga.Type, currentCtx)

		if err != nil {
			return result, err
		}

		maps.Copy(currentCtx, actionResult)
		result[fmt.Sprintf("%d", index)] = actionResult
	}

	result["$is_group"] = true

	return result, nil
}
