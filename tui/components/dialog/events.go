package dialog

import "github.com/sidekick-coder/atlas/tui/features/event"

type Events struct {
	Open  event.VoidEvent
	Close event.VoidEvent
}

func CreateEvents() *Events {
	return &Events{
		Open:  event.VoidEvent{},
		Close: event.VoidEvent{},
	}
}
