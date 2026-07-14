package actionmanager

import (
	"fmt"

	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/utils"
)

type ActionManager struct {
	actions map[string]Action
	config  *config.Config
}

func New(c *config.Config) (*ActionManager, error) {
	configActions := c.GetMap("actions")
	actions := make(map[string]Action)

	for key, action := range configActions {
		actionMap, ok := action.(map[string]any)

		if !ok {
			return nil, fmt.Errorf("invalid action format")
		}

		actionType, ok := actionMap["type"].(string)

		if !ok {
			return nil, fmt.Errorf("action type is not a string")
		}

		actionID, ok := actionMap["id"].(string)

		if !ok {
			actionID = key
		}

		if (actionType == "shell") {
			flat := utils.FlattenMap(actionMap, "")
			flatString := utils.StringifyMap(flat)

			a := ShellAction{
				Options: flatString,
			}

			actions[actionID] = a
		}

	}

	am := &ActionManager{
		actions: actions,
		config:  c,
	}

	return am, nil
}

func (am *ActionManager) Register(name string, action Action) error {
	am.actions[name] = action 
	return nil
}

func (am *ActionManager) Execute(name string, ctx ActionContext) error {
	action, exists := am.actions[name]

	if !exists {
		return fmt.Errorf("action %s not found", name)
	}

	err := action.Execute(ctx)

	if err != nil {
		return fmt.Errorf("error executing action %s: %v", name, err)
	}

	return nil
}
