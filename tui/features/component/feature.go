package component

import (
	"github.com/sidekick-coder/atlas/tui/components/entrymeta"
	"github.com/sidekick-coder/atlas/tui/components/text"
	"github.com/sidekick-coder/atlas/tui/features/component/defintion"
	"github.com/sidekick-coder/atlas/tui/features/component/registry"
)

type Component = defintion.Component
type Registry = registry.Registry

func CreateRegistry() *Registry {
	r := registry.Create()

	LoadDefaultComponents(r)

	return r
}

func Register[T defintion.Component](r *Registry, name string, createFn func() T) {
	r.Register(name, func() defintion.Component {
		return createFn()
	})
}

func LoadDefaultComponents(r *Registry) {
	Register(r, "metas", entrymeta.Create)
	Register(r, "text", text.Create)
}
