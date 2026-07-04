package tui

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/bubbles/key"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/components"
	"github.com/sidekick-coder/atlas/tui/screen"
)

type model struct {
	screens      []screen.Screen
	currentIndex int
	tabBar       *components.TabBar
	width        int
	height       int
	initialized  []bool // tracks whether Init has been called per tab
}

type GlobalKeyMap struct {
	Up          key.Binding
	Down        key.Binding
	PageNext    key.Binding
	PagePrev    key.Binding
	TabNext     key.Binding
	TabPrev     key.Binding
	FocusNext   key.Binding
	MetaReplace key.Binding
	MetaUpdate  key.Binding
	MetaEditor  key.Binding
	MetaAdd     key.Binding
	SyncEntry   key.Binding
	Help        key.Binding
	Quit        key.Binding
}

var GlobalBindings = GlobalKeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

func NewModel(a *app.App) model {
	tabBar := components.NewTabBar()
	screens := []screen.Screen{}

	titles := []string{}
	tabBar.SetTitles(titles)
	tabBar.SetCurrent(0)
	initialized := []bool{}

	return model{
		screens:     screens,
		tabBar:      tabBar,
		initialized: initialized,
	}
}

func (m model) Init() tea.Cmd {
	if len(m.screens) == 0 {
		return nil
	}

	m.initialized[0] = true

	return m.screens[0].Init()
}

func (m model) HandleKeyMessage(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	if key.Matches(msg, GlobalBindings.Quit) {
		return m, tea.Quit
	}

	return m, nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		components.GlobalInput.SetSize(msg.Width, msg.Height)
		components.GlobalToast.SetSize(msg.Width, msg.Height)
		m.tabBar.SetWidth(msg.Width)

	case components.ToastExpiredMsg:
		components.GlobalToast.Hide()
		return m, nil

	case components.ToastShowMsg:
		cmd := components.GlobalToast.Show(msg.Message, msg.Level, 2*time.Second)
		return m, cmd

	case components.BatchMsg:
		cmds := make([]tea.Cmd, len(msg.Msgs))
		for i, inner := range msg.Msgs {
			inner := inner
			cmds[i] = func() tea.Msg { return inner }
		}
		return m, tea.Batch(cmds...)
	}

	// Route all input to GlobalInput while it's active; block everything else.
	if components.GlobalInput.Active() {
		cmd := components.GlobalInput.Update(msg)
		return m, cmd
	}

	keyMsg, ok := msg.(tea.KeyPressMsg) 

	if ok {
		return m.HandleKeyMessage(keyMsg)
	}

	// Tab switching — always available.
	if kp, ok := msg.(tea.KeyPressMsg); ok {
		switch {
		case key.Matches(kp, components.DefaultKeyMap.TabNext):
			if m.currentIndex < len(m.screens)-1 {
				m.currentIndex++
				m.tabBar.SetCurrent(m.currentIndex)
				return m, m.initTab(m.currentIndex)
			}
			return m, nil
		case key.Matches(kp, components.DefaultKeyMap.TabPrev):
			if m.currentIndex > 0 {
				m.currentIndex--
				m.tabBar.SetCurrent(m.currentIndex)
				return m, m.initTab(m.currentIndex)
			}
			return m, nil
		}
	}

	if len(m.screens) == 0 {
		return m, nil
	}

	cmd := m.screens[m.currentIndex].Update(msg)
	return m, cmd
}

// initTab calls Init on a tab the first time it becomes active.
func (m *model) initTab(i int) tea.Cmd {
	if m.initialized[i] {
		return nil
	}
	m.initialized[i] = true
	return m.screens[i].Init()
}

func (m model) View() tea.View {

	if len(m.screens) == 0 {
		return tea.NewView("No screen available")
	}

	const tabBarHeight = 1

	content := m.screens[m.currentIndex].Render()

	if components.GlobalInput.Active() {
		content = components.PlaceOverlay(components.GlobalInput.Box(), content, m.width, m.height)
	}
	if components.GlobalToast.Active() {
		content = components.PlaceOverlay(components.GlobalToast.Box(), content, m.width, m.height)
	}

	// Tab bar sits above the screen content.
	_ = tabBarHeight
	full := m.tabBar.View() + "\n" + content

	v := tea.NewView(full)
	v.AltScreen = true
	return v
}

func Run() error {
	a, err := app.Create()
	if err != nil {
		return err
	}

	p := tea.NewProgram(NewModel(a))
	_, err = p.Run()
	return err
}
