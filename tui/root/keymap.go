package root

import (
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/bubbles/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

func (m *model) HandleGlobalKeyMaps(msg tea.Msg) tea.Cmd {
	keyMsg, ok := msg.(tea.KeyMsg)

	if !ok {
		return nil
	}

	if key.Matches(keyMsg, GlobalBindings.Quit) {
		return tea.Quit
	}

	if key.Matches(keyMsg, GlobalBindings.OpenEntry) {
		m.AddScreen("entry", nil)
		return  nil
	}

	if key.Matches(keyMsg, GlobalBindings.NextScreen) {
		nextIndex := (m.currentIndex + 1) % len(m.screens)
		m.currentIndex = nextIndex
		m.tabBar.SetCurrent(nextIndex)
		return nil
	}

	if key.Matches(keyMsg, GlobalBindings.PrevScreen) {
		prevIndex := (m.currentIndex - 1 + len(m.screens)) % len(m.screens)
		m.currentIndex = prevIndex
		m.tabBar.SetCurrent(prevIndex)

		return nil
	}

	if (key.Matches(keyMsg, GlobalBindings.SyncAll)) {
		m.app.Syncer().All()

		return messages.ToastSuccessCmd("Syncing all entries...", 3 * 1000)
	}

	return nil
}
