package keymaps

import (
	"slices"

	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/logger"
	"github.com/sidekick-coder/atlas/tui/features/keymaps/trigger"
)

type Trigger = trigger.Trigger

var log = logger.Child("package", "keymaps")

func CreateTrigger() Trigger {
	return trigger.Create()
}

func AddTrigger(t trigger.Trigger) {
	log.Debug("Adding trigger", "id", t.ID, "context_id", t.ContextID)

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
	UnloadBindings()

	manager.triggers = slices.DeleteFunc(manager.triggers, func(t trigger.Trigger) bool {
		if t.ID == id {
			log.Debug("Removing trigger", "id", t.ID, "context_id", t.ContextID)
			return true
		}

		return false
	})

	LoadBindings()
}

func RemoveTriggerByContextID(contextID string) {
	UnloadBindings()

	manager.triggers = slices.DeleteFunc(manager.triggers, func(t trigger.Trigger) bool {
		if t.ContextID == contextID {
			log.Debug("Removing trigger", "id", t.ID, "context_id", t.ContextID)
			return true
		}

		return false
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
