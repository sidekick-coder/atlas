package keymaps

import (
	"log/slog"
	"maps"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/utils"
	"github.com/sidekick-coder/atlas/internal/utils/maputil"
	"github.com/sidekick-coder/atlas/tui/action"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/context"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/features/keymaps/trigger"
)

type Manager struct {
	keymaps  []config.Keymap
	bindings []key.Binding
	triggers []trigger.Trigger
	groups   map[string]ManagerGroup
}

type ManagerGroup struct {
	ID     string
	Values []string
}

var manager *Manager = &Manager{
	keymaps:  []config.Keymap{},
	bindings: []key.Binding{},
	triggers: []trigger.Trigger{},
	groups:   map[string]ManagerGroup{},
}

func MapToGroups(m map[string]any) []string {
	groups := []string{}

	flat := utils.FlattenMap(m, "")

	sm := maputil.String(flat)

	for k, v := range sm {
		g := k + "=" + v

		groups = append(groups, g)
	}

	return groups
}

func AddGroup(id string, groups []string) {
	UnloadBindings()

	mg := ManagerGroup{
		ID:     id,
		Values: groups,
	}

	manager.groups[id] = mg

	LoadBindings()
}

func RemoveGroup(id string) {
	UnloadBindings()

	delete(manager.groups, id)

	LoadBindings()
}

func LoadConfigKeymaps(config *config.Config) {
	manager.keymaps = config.GetKeymaps()
}

func UnloadBindings() {
	if len(manager.bindings) == 0 {
		return
	}

	key.Unregister(manager.bindings...)

	manager.bindings = []key.Binding{}
}

func LoadBindings() {
	bindings := []key.Binding{}
	groups := []string{}

	for _, g := range manager.groups {
		groups = append(groups, g.Values...)
	}

	for _, km := range manager.keymaps {
		t, ok := GetKeymapTrigger(km)

		if !ok {
			continue
		}

		b := key.CreateBinding(km.Keys...).
			SetDescription(km.Description).
			SetTags("user").
			SetHelp(km.Keys[0]).
			SetMeta("action", km.Action).
			SetMeta("options", km.ActionOptions).
			SetMeta("trigger", t.ID).
			SetID(km.ID)

		bindings = append(bindings, b)
	}

	key.Register(bindings...)
	manager.bindings = bindings
}

func HandleBinding(msg tea.KeyMsg) tea.Cmd {
	if len(manager.groups) == 0 {
		return nil
	}

	for _, b := range manager.bindings {
		if key.Matches(b) {
			actionId, ok := b.GetMeta("action").(string)

			if !ok {
				return toast.Error("No action defined for key binding: " + b.GetDescription())
			}

			t, ok := GetTriggerByID(b.GetMeta("trigger").(string))

			if !ok {
				return toast.Error("No trigger defined for key binding: " + b.GetDescription())
			}

			c := context.GetById(t.ContextID)

			ctx := c.GetEntriesMap()

			if opts, ok := b.GetMeta("options").(map[string]any); ok {
				maps.Copy(ctx, opts)
			}

			slog.Info("Executing action", "action", actionId, "context", t.ContextID)

			return action.Execute(actionId, ctx)
		}
	}

	return nil
}

func Update(msg tea.Msg) tea.Cmd {
	return chain.Update(msg, chain.OnKey(HandleBinding))
}
