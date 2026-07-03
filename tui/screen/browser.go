package screen

import (
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/tui/components"
)

type entriesLoadedMsg struct {
	entries []models.Entry
	err     error
}

type metasLoadedMsg struct {
	metas []models.EntryMeta
	err   error
}

type BrowserScreen struct {
	app        *app.App
	entryList  *components.EntryList
	entryMetas *components.EntryMetas
	width      int
	height     int
}

func NewBrowserScreen(a *app.App) *BrowserScreen {
	return &BrowserScreen{
		app:        a,
		entryList:  components.NewEntryList(),
		entryMetas: components.NewEntryMetas(),
	}
}

func (m *BrowserScreen) Init() tea.Cmd {
	return m.loadEntries()
}

func (m *BrowserScreen) loadEntries() tea.Cmd {
	return func() tea.Msg {
		entries, err := m.app.EntryRepo().List()
		return entriesLoadedMsg{entries: entries, err: err}
	}
}

func (m *BrowserScreen) loadMetas(entryID int64) tea.Cmd {
	return func() tea.Msg {
		metas, err := m.app.EntryMetaRepo().ListByEntryID(entryID)
		return metasLoadedMsg{metas: metas, err: err}
	}
}

func (m *BrowserScreen) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.updateSizes()

	case entriesLoadedMsg:
		if msg.err == nil {
			m.entryList.SetEntries(msg.entries)
			if entry := m.entryList.SelectedEntry(); entry != nil {
				return m.loadMetas(entry.ID)
			}
		}

	case metasLoadedMsg:
		if msg.err == nil {
			m.entryMetas.SetMetas(msg.metas)
		}

	case components.EntrySelectedMsg:
		return m.loadMetas(msg.Entry.ID)
	
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return tea.Quit
		}

	default:
		return m.entryList.Update(msg)
	}

	return nil
}

func (m *BrowserScreen) updateSizes() {
	listWidth := m.width / 3
	metasWidth := m.width - listWidth

	m.entryList.SetSize(listWidth, m.height)
	m.entryMetas.SetSize(metasWidth, m.height)
}

func (m *BrowserScreen) View() tea.View {
	listView := m.entryList.View()
	metasView := m.entryMetas.View()

	content := lipgloss.JoinHorizontal(lipgloss.Top, listView, metasView)

	v := tea.NewView(content)
	v.AltScreen = true

	return v
}
