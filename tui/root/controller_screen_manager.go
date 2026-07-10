package root

import (
	"fmt"
	"log"
	"maps"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/messages"
	"github.com/sidekick-coder/atlas/tui/models"
	"github.com/sidekick-coder/atlas/tui/screen/empty"
	"github.com/sidekick-coder/atlas/tui/screen/entry"
	"github.com/sidekick-coder/atlas/tui/screen/entrysingle"
	"github.com/sidekick-coder/atlas/tui/screen/entrytable"
	"github.com/sidekick-coder/atlas/tui/screen/syncer"
)

func (m *model) LoadUserScreen(screen config.Screen) (models.ScreenFactory, error) {
	original := m.availableScreens[screen.Type]

	if original == nil {
		return nil, fmt.Errorf("invalid screen type: %s", screen.Type)
	}

	fac := func(p models.ScreenPayload) (models.Screen, error) {
		maps.Copy(p.Options, screen.Options)

		return original(p)
	}

	return fac, nil

}

func (m model) LoadScreenRegistry() tea.Cmd {
	m.availableScreens["empty"] = empty.Create
	m.availableScreens["entry_list"] = entry.Create
	m.availableScreens["entry_table"] = entrytable.Create
	m.availableScreens["entry_single"] = entrysingle.Create
	m.availableScreens["syncer"] = syncer.Create

	us, err := m.app.Config().GetScreens()

	if err != nil {
		return messages.ToastErrorCmd(fmt.Sprintf("Failed to load screens from config: %v", err), 3*1000)
	}

	for _, s := range us {
		fac, err := m.LoadUserScreen(s)

		if err != nil {
			return messages.ToastErrorCmd(err.Error())
		}

		m.availableScreens[s.ID] = fac
	}

	loaded := []string{}

	for k := range m.availableScreens {
		loaded = append(loaded, k)
	}

	log.Printf("Loaded screen registry: %v", loaded)

	return nil
}

func (m *model) GetCurrentScreen() (models.Screen, bool) {
	if len(m.screens) == 0 || m.currentIndex < 0 || m.currentIndex >= len(m.screens) {
		return nil, false
	}

	return m.screens[m.currentIndex], true
}

func (m *model) SetCurrentScreen(index int) tea.Cmd {
	if index < 0 || index >= len(m.screens) {
		return nil
	}

	m.currentIndex = index

	m.LoadBindings()
	m.LoadTabs()

	s := m.screens[index]

	s.SetSize(m.width, m.screenHeight)

	key.ClearBindings()

	return s.Init()
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

	log.Printf("Creating screen instance for name: %s, factory: %v", name, fac)

	if fac == nil {
		return nil, fmt.Errorf("screen factory not found for name: %s", name)
	}

	if Program == nil {
		return nil, fmt.Errorf("program is not initialized")
	}

	p := models.ScreenPayload{
		App:     m.app,
		Options: options,
		Program: Program,
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
		return messages.ToastErrorCmd(fmt.Sprintf("Failed to create screen: %v", err))
	}

	m.screens = append(m.screens, s)
	index := len(m.screens) - 1
	return m.SetCurrentScreen(index)
}

func (m *model) ReplaceScreen(index int, name string, options map[string]any) tea.Cmd {

	s, err := m.CreateScreenInstance(name, options)

	if err != nil {
		return messages.ToastErrorCmd(fmt.Sprintf("Failed to create screen: %v", err), 3*1000)
	}

	if index < 0 || index >= len(m.screens) {
		return messages.ToastErrorCmd(fmt.Sprintf("Invalid screen index: %d", index), 3*1000)
	}

	m.screens[index] = s
	return m.SetCurrentScreen(index)
}

func (m *model) RemoveScreen(index int) error {
	if index < 0 || index >= len(m.screens) {
		return fmt.Errorf("invalid screen index: %d", index)
	}

	m.screens = append(m.screens[:index], m.screens[index+1:]...)

	if m.currentIndex >= len(m.screens) {
		m.currentIndex = len(m.screens) - 1
	}

	if len(m.screens) > 0 {
		m.SetCurrentScreen(m.currentIndex)
	} else {
		m.currentIndex = -1
	}

	return nil
}

func (m *model) AddScreenEmpty() tea.Cmd {
	entries := []empty.Entry{}

	entries = append(entries, empty.Entry{
		ID:      "entry_list",
		Options: map[string]any{},
	})

	entries = append(entries, empty.Entry{
		ID:      "entry_table",
		Options: map[string]any{},
	})

	us, err := m.app.Config().GetScreens()

	if err != nil {
		return messages.ToastErrorCmd(fmt.Sprintf("Failed to load screens from config: %v", err), 3*1000)
	}

	for _, s := range us {
		entries = append(entries, empty.Entry{
			ID:      s.ID,
			Options: s.Options,
		})
	}

	options := map[string]any{
		"entries": entries,
	}

	return m.AddScreen("empty", options)
}
