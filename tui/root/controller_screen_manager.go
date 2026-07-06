package root

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/messages"
	"github.com/sidekick-coder/atlas/tui/models"
	"github.com/sidekick-coder/atlas/tui/screen/empty"
	"github.com/sidekick-coder/atlas/tui/screen/entry"
	"github.com/sidekick-coder/atlas/tui/screen/entrysingle"
)

func (m model) LoadScreenRegistry() tea.Cmd {
	m.availableScreens["empty"] = empty.Create
	m.availableScreens["entry_list"] = entry.Create
	m.availableScreens["entry_single"] = entrysingle.Create

	return nil
}

func (m *model) GetCurrentScreen() (models.Screen, bool) {
	if len(m.screens) == 0 || m.currentIndex < 0 || m.currentIndex >= len(m.screens) {
		return nil, false
	}

	return m.screens[m.currentIndex], true
}

func (m *model) SetCurrentScreen(index int) {
	if index < 0 || index >= len(m.screens) {
		return
	}

	m.currentIndex = index

	m.LoadBindings()
	m.LoadTabs()

	m.SetSize(m.width, m.height)

	s := m.screens[index]

	s.Init()
}

func (m *model) LoadTabs() {
	m.tabBar.Clear()

	for i, s := range m.screens {
		m.tabBar.Add(fmt.Sprintf("[%d] %s", i, s.Title()))
	}

	m.tabBar.SetCurrent(m.currentIndex)
}

func (m *model) CreateScreenInstance(name string, options map[string]any) (models.Screen, error) {
	fac := m.availableScreens[name]

	if fac == nil {
		return nil, fmt.Errorf("screen factory not found for name: %s", name)
	}

	p := models.ScreenPayload{
		App: m.app,
		Options: options,
	}

	s, err := fac(p)

	if err != nil {
		return nil, fmt.Errorf("failed to create screen instance: %w", err)
	}

	return s, nil
}

func (m *model) AddScreen(name string, args ...map[string]any) tea.Cmd {
	options := map[string]any{}

	if len(args) > 0 {
		options = args[0]
	}

	s, err := m.CreateScreenInstance(name, options)

	if err != nil {
		return messages.ToastErrorCmd(fmt.Sprintf("Failed to create screen: %v", err), 3 * 1000)
	}

	m.screens = append(m.screens, s)
	index := len(m.screens) - 1
	m.SetCurrentScreen(index)
	m.LoadTabs()

	return  nil
}

func (m *model) ReplaceScreen(index int, name string, options map[string]any) tea.Cmd {

	s, err := m.CreateScreenInstance(name, options)

	if err != nil {
		return messages.ToastErrorCmd(fmt.Sprintf("Failed to create screen: %v", err), 3 * 1000)
	}

	if index < 0 || index >= len(m.screens) {
		return messages.ToastErrorCmd(fmt.Sprintf("Invalid screen index: %d", index), 3 * 1000)
	}

	m.screens[index] = s
	m.SetCurrentScreen(index)
	m.LoadTabs()

	return nil
}

