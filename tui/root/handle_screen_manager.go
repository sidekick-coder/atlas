package root

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/messages"
	"github.com/sidekick-coder/atlas/tui/models"
)

func (m *model) AddScreen(msg messages.AddScreen) tea.Cmd {
	fac := m.availableScreens[msg.Name]

	if fac == nil {
		return messages.ToastErrorCmd("Error adding screen, screen not found: "+msg.Name, 3*1000)
	}

	p := models.ScreenPayload{
		App: m.app,
		Options: msg.Options,
	}

	s, err := fac(p)

	if err != nil {
		return messages.ToastErrorCmd("Error adding screen: "+err.Error(), 3*1000)
	}

	m.screens = append(m.screens, s)

	index := len(m.screens) - 1

	m.tabBar.Add(fmt.Sprintf("[%d]: %s", index, s.Title()))
	m.SetCurrentScreen(index)

	return  nil
}


func (m *model) HandleScreeManager(msg tea.Msg) tea.Cmd {
	if as, ok := msg.(messages.AddScreen); ok {
		return m.AddScreen(as)
	}

	return nil
}
