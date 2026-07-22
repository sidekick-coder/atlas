package action

import (
	"maps"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/action"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/action/actions"
	tc "github.com/sidekick-coder/atlas/tui/components/toast"
)

type ActionContext struct {
	ID      string
	Context map[string]any
}

type Manager struct {
	app     *app.App
	action  *action.Manager
	context map[string]ActionContext
}

var manager *Manager = &Manager{
	context: map[string]ActionContext{},
}

func Load(a *app.App) {
	manager.action = a.Action
	manager.app = a
	manager.action.LoadConfigActions(a.Config())

	manager.action.AddDefinition("entry-sync", actions.EntrySyncAction)
}

func AddDefinition(id string, fn func(map[string]any) (map[string]any, error)) {
	manager.action.AddDefinition(id, fn)
}

func AddContext(id string, context map[string]any) {
	ac := ActionContext{
		ID:      id,
		Context: context,
	}

	manager.context[id] = ac
}

func RemoveContext(id string) {
	delete(manager.context, id)
}

func Execute(id string) tea.Cmd {
	ctx := make(map[string]any)

	for _, c := range manager.context {
		maps.Copy(ctx, c.Context)
	}

	result, err := manager.action.Execute(id, ctx)

	if err != nil {
		return tc.Error(err.Error())
	}

	resultList := make([]map[string]any, 0)

	isGroup, ok := result["$is_group"].(bool)

	if ok && isGroup {
		for _, v := range result {
			if vMap, ok := v.(map[string]any); ok {
				resultList = append(resultList, vMap)
			}
		}
	}

	if !ok || !isGroup {
		resultList = append(resultList, result)
	}

	for _, r := range resultList {
		if rmsg, ok := r["tea_message"].(tea.Msg); ok {
			return func() tea.Msg {
				return rmsg
			}
		}
	}

	return nil
}
