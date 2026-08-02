package entrylist

import "github.com/sidekick-coder/atlas/tui/features/event"

type Events struct {
	Change event.VoidEvent
}

func CreateEvents() *Events {
	return &Events{
		Change:  event.VoidEvent{},
	}
}
