package screen

import (
	"os"
	"os/exec"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	"github.com/charmbracelet/bubbles/key"
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

type focusedPanel int

const (
	focusList focusedPanel = iota
	focusMetas
)

type BrowserScreen struct {
	app        *app.App
	entryList  *components.EntryList
	entryMetas *components.EntryMetas
	footer     *components.Footer
	help       *components.HelpScreen
	showHelp   bool
	focus      focusedPanel
	width      int
	height     int
}

func NewBrowserScreen(a *app.App) *BrowserScreen {
	km := components.DefaultKeyMap
	s := &BrowserScreen{
		app:        a,
		entryList:  components.NewEntryList(),
		entryMetas: components.NewEntryMetas(),
		footer:     components.NewFooter(),
		help:       components.NewHelpScreen(km),
		focus:      focusList,
	}
	s.entryList.SetFocused(true)
	s.updateFooter()
	return s
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
				m.entryMetas.SetEntryID(entry.ID)
				return m.loadMetas(entry.ID)
			}
		}

	case metasLoadedMsg:
		if msg.err == nil {
			m.entryMetas.SetMetas(msg.metas)
		}

	case components.EntrySelectedMsg:
		m.entryMetas.SetEntryID(msg.Entry.ID)
		return m.loadMetas(msg.Entry.ID)

	case components.MetaInputSubmitMsg:
		return m.saveMeta(msg.EntryID, msg.Name, msg.Value)

	case components.MetaOpenEditorMsg:
		return m.openEditor(msg.EntryID, msg.Name, msg.CurrentValue)

	case metaEditorDoneMsg:
		if msg.err == nil {
			return m.saveMeta(msg.entryID, msg.name, msg.value)
		}

	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, components.DefaultKeyMap.Quit):
			return tea.Quit
		case key.Matches(msg, components.DefaultKeyMap.Help):
			m.showHelp = !m.showHelp
			return nil
		case key.Matches(msg, components.DefaultKeyMap.FocusNext):
			m.cycleFocus()
			return nil
		}
		if m.focus == focusList {
			return m.entryList.Update(msg)
		}
		return m.entryMetas.Update(msg)
	}

	return nil
}

type metaEditorDoneMsg struct {
	entryID int64
	name    string
	value   string
	err     error
}

func (m *BrowserScreen) saveMeta(entryID int64, name, value string) tea.Cmd {
	return func() tea.Msg {
		_, err := m.app.EntryMetaRepo().UpsertMany(entryID, map[string]string{name: value})
		if err != nil {
			return metasLoadedMsg{err: err}
		}
		metas, err := m.app.EntryMetaRepo().ListByEntryID(entryID)
		return metasLoadedMsg{metas: metas, err: err}
	}
}

func (m *BrowserScreen) openEditor(entryID int64, name, currentValue string) tea.Cmd {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	tmpFile, err := os.CreateTemp("", "atlas-meta-*.txt")
	if err != nil {
		return nil
	}
	tmpPath := tmpFile.Name()
	_, _ = tmpFile.WriteString(currentValue)
	tmpFile.Close()

	return tea.ExecProcess(exec.Command(editor, tmpPath), func(err error) tea.Msg {
		if err != nil {
			_ = os.Remove(tmpPath)
			return metaEditorDoneMsg{err: err}
		}
		data, readErr := os.ReadFile(tmpPath)
		_ = os.Remove(tmpPath)
		if readErr != nil {
			return metaEditorDoneMsg{err: readErr}
		}
		return metaEditorDoneMsg{entryID: entryID, name: name, value: string(data)}
	})
}

func (m *BrowserScreen) cycleFocus() {
	if m.focus == focusList {
		m.focus = focusMetas
	} else {
		m.focus = focusList
	}
	m.entryList.SetFocused(m.focus == focusList)
	m.entryMetas.SetFocused(m.focus == focusMetas)
	m.updateFooter()
}

func (m *BrowserScreen) updateFooter() {
	km := components.DefaultKeyMap
	shared := []key.Binding{km.Up, km.Down, km.FocusNext, km.Help, km.Quit}
	if m.focus == focusMetas {
		metaBindings := []key.Binding{km.MetaAdd, km.MetaReplace, km.MetaUpdate, km.MetaEditor}
		m.footer.SetBindings(append(shared[:3], append(metaBindings, shared[3:]...)...)...)
	} else {
		m.footer.SetBindings(shared...)
	}
}

func (m *BrowserScreen) updateSizes() {
	const footerHeight = 1
	contentHeight := m.height - footerHeight

	listWidth := m.width / 3
	metasWidth := m.width - listWidth

	m.entryList.SetSize(listWidth, contentHeight)
	m.entryMetas.SetSize(metasWidth, contentHeight)
	m.footer.SetWidth(m.width)
	m.help.SetSize(m.width, m.height)
}

// Render returns the screen content as a plain string (no tea.View wrapper).
// The root model uses this to composite overlays before creating the final View.
func (m *BrowserScreen) Render() string {
	const footerHeight = 1
	contentHeight := m.height - footerHeight

	mainArea := lipgloss.NewStyle().Width(m.width).Height(contentHeight).Render(
		lipgloss.JoinHorizontal(lipgloss.Top, m.entryList.View(), m.entryMetas.View()),
	)
	bg := lipgloss.JoinVertical(lipgloss.Left, mainArea, m.footer.View())

	if m.showHelp {
		return m.help.View()
	}
	return bg
}

func (m *BrowserScreen) View() tea.View {
	v := tea.NewView(m.Render())
	v.AltScreen = true
	return v
}

