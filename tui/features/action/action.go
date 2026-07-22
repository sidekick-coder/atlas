package action

import (
	"maps"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/action"
	tc "github.com/sidekick-coder/atlas/tui/components/toast"
	ta "github.com/sidekick-coder/atlas/tui/features/action/actions/toast"
)

type ActionContext struct {
	ID      string
	Context map[string]any
}

type Manager struct {
	action  *action.Manager
	context map[string]ActionContext
}

var manager *Manager = &Manager{
	context: map[string]ActionContext{},
	action:  nil,
}

func SetManager(m *action.Manager) {
	manager.action = m
}

func Init() tea.Cmd {
	t := ta.Create()

	manager.action.AddDefinition(t.ID, t.Execute)

	return nil
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

	resultMsg, ok := result["tea_message"].(tea.Msg)

	if ok {
		return func() tea.Msg {
			return resultMsg
		}
	}

	return nil
}
