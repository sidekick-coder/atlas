package registry

import "github.com/sidekick-coder/atlas/tui/features/component/defintion"

type RegistryItem struct {
	Name       string
	CreateFunc func() defintion.Component
}

type Registry struct {
	items map[string]RegistryItem
}

func Create() *Registry {
	return &Registry{
		items: map[string]RegistryItem{},
	}
}

func (r *Registry) Register(name string, createFn func() defintion.Component) {
	r.items[name] = RegistryItem{
		Name:       name,
		CreateFunc: createFn,
	}
}

func (r *Registry) Get(name string) (defintion.Component, bool) {
	item, exists := r.items[name]

	if !exists {
		return nil, false
	}

	return item.CreateFunc(), true
}

