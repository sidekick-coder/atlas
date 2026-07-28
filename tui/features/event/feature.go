package event

type Event[T any] struct {
	handlers []func(T)
}

func Create[T any]() *Event[T] {
	return &Event[T]{
		handlers: []func(T){},
	}
}

func (e *Event[T]) On(fn func(T)) {
	e.handlers = append(e.handlers, fn)
}

func (e *Event[T]) Emit(value T) {
	for _, fn := range e.handlers {
		fn(value)
	}
}


