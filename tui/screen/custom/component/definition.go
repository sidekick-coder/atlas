package component

var Definitions = map[string]func(DefinitionPayload) Definition{}

type Definition interface {
	Render() string
}

type DefinitionPayload struct {
	Width   int
	Height  int
	Options map[string]any
}

type DefinitionFactory func(payload DefinitionPayload) Definition

func RegisterDefinition(name string, factory DefinitionFactory) {
	Definitions[name] = factory
}
