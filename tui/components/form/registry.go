package form

import (
	"github.com/sidekick-coder/atlas/tui/components/form/field"
	"github.com/sidekick-coder/atlas/tui/components/textfield"
)


type Registry struct {
	entries map[string] field.DefinitionFunc
}

func CreateRegistry() *Registry {
	return &Registry{
		entries: make(map[string] field.DefinitionFunc),
	}
}

func Register[T field.Definition](r *Registry, name string, def func() T) {
	r.Register(name, func() field.Definition {
		return def()
	})
}

func (r *Registry) Load() {
	Register(r, "textfield", textfield.Create)
}

func (r *Registry) Register(name string, field field.DefinitionFunc) {
	r.entries[name] = field
}

func (r *Registry) Get(name string) (field.DefinitionFunc, bool) {
	field, ok := r.entries[name]

	return field, ok
}
