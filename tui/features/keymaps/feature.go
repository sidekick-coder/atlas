package keymaps

import (
	"fmt"
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
			a, ok := b.GetMeta("action").(config.Action)

			if !ok {
				return toast.Error("No action defined for key binding: " + b.GetDescription())
			}

			t, ok := GetTriggerByID(b.GetMeta("trigger").(string))

			if !ok {
				return toast.Error("No trigger defined for key binding: " + b.GetDescription())
			}

			c, ok := context.GetById(t.ContextID)

			if !ok {
				return toast.Error(fmt.Sprintf("No context defined for key binding: %s, context: %s", b.GetDescription(), t.ContextID))
			}

			ctx := c.GetEntriesMap()

			maps.Copy(ctx, a.Options)

			return action.Execute(a.Type, ctx)
		}
	}

	return nil
}

func Update(msg tea.Msg) tea.Cmd {
	return chain.Update(msg, chain.OnKey(HandleBinding))
}
