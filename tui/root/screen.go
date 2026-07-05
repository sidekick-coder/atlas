package root

import "github.com/sidekick-coder/atlas/tui/models"


func (m *model) GetCurrentScreen() (models.Screen, bool) {
	if len(m.screens) == 0 || m.currentIndex < 0 || m.currentIndex >= len(m.screens) {
		return nil, false
	}

	return m.screens[m.currentIndex], true
}
