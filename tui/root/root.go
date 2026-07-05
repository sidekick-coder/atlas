package root

import (
	"fmt"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/components"
	"github.com/sidekick-coder/atlas/tui/messages"
	"github.com/sidekick-coder/atlas/tui/screen"
	"github.com/sidekick-coder/atlas/tui/screen/entry"
	"github.com/sidekick-coder/atlas/tui/screen/entrysingle"
)

type model struct {
	app          *app.App
	screens      []screen.Screen
	currentIndex int
	width        int
	height       int

	tabBar  *components.TabBar
	toolbar *components.Toolbar
	footer  *components.Footer
}

func New(a *app.App) model {
	screens := []screen.Screen{}

	entryScreen := entry.Create(a)
	screens = append(screens, entryScreen)

	tabBar := components.NewTabBar()
	tabBar.Add("0: " + entryScreen.Title())

	toolbar := components.NewToolbar()
	toolbar.SetTitle("󰉋 " + a.WorkspacePath())

	footer := components.NewFooter()

	m := model{
		app:          a,
		currentIndex: 0,
		screens:      screens,
		tabBar:       tabBar,
		toolbar:      toolbar,
		footer:       footer,
	}

	m.SetCurrentScreen(0)
	m.SetBindings()

	return m
}

func (m *model) SetBindings() {
	bindings := m.GetBindings()

	m.footer.SetBindings(bindings...)
}

func (m *model) SetCurrentScreen(index int) {
	if index < 0 || index >= len(m.screens) {
		return
	}

	m.currentIndex = index
	m.tabBar.SetCurrent(index)
	m.footer.SetBindings(m.GetBindings()...)

	m.SetSize(m.width, m.height)
}

func (m *model) AddScreen(name string, options map[string]any) {
	if name == "entry" {
		s := entry.Create(m.app)
		m.screens = append(m.screens, s)
		index := len(m.screens) - 1
		m.tabBar.Add(fmt.Sprintf("%d: %s", index, s.Title()))

		m.SetCurrentScreen(index)
		return
	}

	if name == "entry-single" {
		entryID, ok := options["entry_id"].(int64)

		if !ok {
			fmt.Println("entryID not provided or not an int64")
			return
		}

		s := entrysingle.Create(m.app, entryID)
		m.screens = append(m.screens, s)
		index := len(m.screens) - 1
		m.tabBar.Add(fmt.Sprintf("%d: %s", index, s.Title()))

		m.SetCurrentScreen(index)
		return
	}

	toastMsg := components.ToastShowMsg{
		Message: fmt.Sprintf("Unknown screen: %s", name),
		Level:   1,
	}

	m.Update(toastMsg)
}

func (m *model) SetSize(width int, height int) {
	m.width = width
	m.height = height

	components.GlobalInput.SetSize(width, height)
	components.GlobalToast.SetSize(width, height)

	m.tabBar.SetWidth(width)
	m.toolbar.SetWidth(width)
	m.footer.SetWidth(width)

	toolbarHeight := 1
	tabBarHeight := 1
	footerHeight := 1
	contentHeight := height - toolbarHeight - tabBarHeight - footerHeight

	for _, s := range m.screens {
		s.SetSize(width, contentHeight)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	handlers := []func(tea.Msg) tea.Cmd{}

	handlers = append(handlers,
		m.HandleActions,
		m.actionBindingMessageHandler,
		m.HandleGlobalKeyMaps,
	)

	for _, handler := range handlers {
		cmd := handler(msg)

		if cmd != nil {
			return m, cmd
		}
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
	case components.ToastExpiredMsg:
		components.GlobalToast.Hide()
		return m, nil

	case components.ToastShowMsg:
		cmd := components.GlobalToast.Show(msg.Message, msg.Level, 2*time.Second)
		return m, cmd
	case messages.Toast:
		cmd := components.GlobalToast.Show(msg.Message, components.ToastInfo, 2*time.Second)
		return m, cmd
	case messages.AddScreen:
		m.AddScreen(msg.Name, msg.Options)

		return m, nil
	}

	if len(m.screens) == 0 {
		return m, nil
	}

	cmd := m.screens[m.currentIndex].Update(msg)

	return m, cmd
}

func (m model) View() tea.View {
	if len(m.screens) == 0 {
		return tea.NewView("No screen available")
	}

	currentScreen := m.screens[m.currentIndex]

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		m.toolbar.Render(),
		m.tabBar.Render(),
		currentScreen.Render(),
		m.footer.Render(),
	)

	if components.GlobalInput.Active() {
		content = components.PlaceOverlay(components.GlobalInput.Box(), content, m.width, m.height)
	}

	if components.GlobalToast.Active() {
		content = components.PlaceOverlay(components.GlobalToast.Box(), content, m.width, m.height)
	}

	const tabBarHeight = 1
	const toolbarHeight = 1

	v := tea.NewView(content)
	v.AltScreen = true

	return v
}
