package form

import "github.com/sidekick-coder/atlas/tui/features/event"

type Events struct {
	Cancel  event.VoidEvent
	Submit  event.VoidEvent
}

func CreateEvents() *Events {
	return &Events{
		Cancel:  event.VoidEvent{},
		Submit:  event.VoidEvent{},
	}
}
