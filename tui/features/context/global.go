package context

var registry = []*Feature{}

func GetRegistry() []*Feature {
	return registry
}

func GetById(id string) (*Feature, bool) {
	for _, f := range registry {
		if f.id == id {
			return f, true
		}
	}

	return nil, false
}

