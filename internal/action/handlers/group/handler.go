package action

import (
	"errors"
	"github.com/sidekick-coder/atlas/internal/template"
	"os/exec"
	"strings"
)

type GroupAction struct {
	ID	  string 
	Options map[string]string
}

type Handler struct {
	ExecuteAction func(ctx map[string]any) (map[string]any, error)
	Options map[string]string
}

func ParseGroupAction(action map[string]any) (GroupAction, error) {
	id, ok := action["id"].(string)

	if !ok {
		return GroupAction{}, errors.New("invalid or missing 'id' in group action")
	}

	options, ok := action["options"].(map[string]string)
	if !ok {
		return GroupAction{}, errors.New("invalid or missing 'options' in group action")
	}

	ga := GroupAction{
		ID:      id,
		Options: options,
	}

	return ga, nil
}

func (c Handler) Execute(ctx map[string]any) (map[string]any, error) {
	result := make(map[string]any)

	actions, ok := ctx["actions"].([]map[string]any)
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

	return result, nil
}
