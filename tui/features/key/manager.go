package key

import (
	"slices"

	tea "charm.land/bubbletea/v2"
)

type Manager struct {
	leader     string
	pending    []string
	registered []Binding
}

func NewManager() Manager {
	return Manager{
		leader:     "space", // default leader key is space
		pending:    []string{},
		registered: []Binding{},
	}
}

func (m *Manager) Register(bindings ...Binding) {
	m.registered = append(manager.registered, bindings...)
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

func (m *Manager) hasPossibleMatch() bool {
	return slices.ContainsFunc(m.registered, m.hasPossibleMatchForBinding)
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

	normalized := normalize(km)

	manager.pending = append(manager.pending, normalized)

	hasPossibleMatch := manager.hasPossibleMatch()

	if !hasPossibleMatch {
		manager.pending = nil
	}

	return nil
}

func (m *Manager) GetBindings() []Binding {
	return m.registered
}

func (m *Manager) ClearBindings() {
	m.registered = []Binding{}
}
