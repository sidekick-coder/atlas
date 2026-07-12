package key

import (
	"log"
	"os"
	"slices"

	tea "charm.land/bubbletea/v2"
)

type Manager struct {
	debug      bool
	leader     string
	pending    []string
	registered []Binding
}

func NewManager() Manager {
	return Manager{
		debug:      os.Getenv("DEBUG") == "true",
		leader:     "space", // default leader key is space
		pending:    []string{},
		registered: []Binding{},
	}
}

func (m *Manager) Register(bindings ...Binding) {
	m.registered = append(manager.registered, bindings...)

	if !m.debug {
		return
	}

	for _, b := range bindings {
		for _, k := range b.keys {
			log.Printf("Register binding: %s -> %s (%v)\n", k.tokens, b.GetDescription(), b.GetTags())
		}
	}
}

func (m *Manager) Unregister(bindings ...Binding) {
	for _, b := range bindings {
		for i := range m.registered {
			if m.registered[i].id == b.id {
				m.registered = append(m.registered[:i], m.registered[i+1:]...)

				if m.debug {
					log.Printf("Unregistered binding: %s -> %s (%v)\n", b.keys, b.GetDescription(), b.GetTags())
				}
				break
			}
		}
	}
}

func normalize(km tea.KeyMsg) string {
	switch km.String() {
	case manager.leader:
		return "<leader>"
	default:
		return km.String()
	}
}

func (m *Manager) hasPossibleMatchForBindingKey(b BindingKey) bool {
	tokens := b.tokens

	if len(m.pending) > len(tokens) {
		return false
	}

	for i := range m.pending {
		if m.pending[i] != tokens[i] {
			return false
		}
	}

	return true

}

func (m *Manager) hasPossibleMatchForBinding(b Binding) bool {
	return slices.ContainsFunc(b.keys, m.hasPossibleMatchForBindingKey)
}

func (m *Manager) hasPossibleMatch() (Binding, bool) {
	for _, b := range m.registered {
		if m.hasPossibleMatchForBinding(b) {
			return b, true
		}
	}

	return Binding{}, false
}

func MatchBiningKey(k BindingKey) bool {

	if len(manager.pending) != len(k.tokens) {
		return false
	}

	for i := range k.tokens {
		if manager.pending[i] != k.tokens[i] {
			return false
		}
	}

	return true
}

func Matches(b Binding) bool {
	mached := slices.ContainsFunc(b.keys, MatchBiningKey)

	if mached {
		manager.pending = nil
		return true
	}

	return false
}

func (m *Manager) HandleKeypress(msg tea.Msg) tea.Cmd {
	km, ok := msg.(tea.KeyMsg)

	if !ok {
		return nil
	}

	// forse quit
	if km.String() == "ctrl+c" {
		return tea.Quit
	}

	normalized := normalize(km)

	manager.pending = append(manager.pending, normalized)

	b, hasPossibleMatch := manager.hasPossibleMatch()

	if m.debug {
		log.Printf("Key pressed: %s, pending: %v, possible match: %s\n", normalized, manager.pending, b.GetDescription())
	}

	if !hasPossibleMatch {
		manager.pending = nil
	}

	return nil
}

func (m *Manager) GetBindings() []Binding {
	return m.registered
}
func (m *Manager) GetBindingsByTags(tags ...string) []Binding {
	var bindings []Binding

	for _, b := range m.registered {
		if slices.ContainsFunc(b.GetTags(), func(tag string) bool {
			return slices.Contains(tags, tag)
		}) {
			bindings = append(bindings, b)
		}
	}

	return bindings
}

func (m *Manager) ClearBindings() {
	m.registered = []Binding{}
}
