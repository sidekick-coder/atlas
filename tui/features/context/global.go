package context

var registry = []*Feature{}

func GetRegistry() []*Feature {
	return registry
}

func GetById(id string) *Feature {
	for _, f := range registry {
		if f.id == id {
			return f
		}
	}

	return nil
}

