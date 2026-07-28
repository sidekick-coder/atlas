package keymaps

import (
	"slices"

	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/tui/features/keymaps/trigger"
)

type Trigger = trigger.Trigger

func CreateTrigger() Trigger {
	return trigger.Create()
}

func AddTrigger(t trigger.Trigger) {
	manager.triggers = append(manager.triggers, t)

	LoadBindings()
}

func GetTriggerByID(id string) (Trigger, bool) {
	for _, t := range manager.triggers {
		if t.ID == id {
			return t, true
		}
	}

	return Trigger{}, false
}

func RemoveTriggerByID(id string) {
	manager.triggers = slices.DeleteFunc(manager.triggers, func(t trigger.Trigger) bool {
		return t.ID == id
	})

	LoadBindings()
}

func RemoveTriggerByContextID(contextID string) {
	manager.triggers = slices.DeleteFunc(manager.triggers, func(t trigger.Trigger) bool {
		return t.ContextID == contextID
	})

	LoadBindings()
}

func GetKeymapTrigger(km config.Keymap) (Trigger, bool) {
	if len(manager.triggers) == 0 {
		return Trigger{}, false
	}

	for _, t := range manager.triggers {
		if t.Test(km) {
			return t, true
		}
	}

	return Trigger{}, false
}
