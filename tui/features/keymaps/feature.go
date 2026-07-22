package keymaps

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/utils"
	"github.com/sidekick-coder/atlas/internal/utils/maputil"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/action"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

type Manager struct {
	keymaps  []config.Keymap
	bindings []key.Binding
	groups   map[string]ManagerGroup
}

type ManagerGroup struct {
	ID string
	Values []string
}

var manager *Manager = &Manager{
	keymaps:  []config.Keymap{},
	bindings: []key.Binding{},
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
		ID: id,
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

	for _, action := range manager.keymaps {
		if action.HasGroup(groups...) == false {
			continue
		}

		b := key.CreateBinding(action.Keys...).
			SetDescription(action.Description).
			SetTags("user").
			SetHelp(action.Keys[0]).
			SetMeta("action", action.Action).
			SetID(action.ID)

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
			actionId := b.GetMeta("action")

			if actionId == nil {
				return toast.Error("No action defined for key binding: " + b.GetDescription())
			}

			return action.Execute(actionId.(string))
		}
	}

	return nil
}

func Update(msg tea.Msg) tea.Cmd {
	return chain.Update(msg, chain.OnKey(HandleBinding))
}
