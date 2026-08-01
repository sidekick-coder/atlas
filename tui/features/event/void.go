package event

type VoidEvent struct {
	handlers []func()
}

func CreateVoid() *VoidEvent {
	return &VoidEvent{
		handlers: []func(){},
	}
}

func (e *VoidEvent) On(fn func()) {
	e.handlers = append(e.handlers, fn)
}

func (e *VoidEvent) Emit() {
	for _, fn := range e.handlers {
		fn()
	}
}
