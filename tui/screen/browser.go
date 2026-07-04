package screen

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	"github.com/charmbracelet/bubbles/key"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/metadata"
	"github.com/sidekick-coder/atlas/internal/models"
	entryrepo "github.com/sidekick-coder/atlas/internal/repository/entry"
	"github.com/sidekick-coder/atlas/tui/components"
)

type entriesLoadedMsg struct {
	entries []models.Entry
	total   int
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
	app           *app.App
	entryList     *components.EntryList
	entryMetas    *components.EntryMetas
	footer        *components.Footer
	help          *components.HelpScreen
	showHelp      bool
	focus         focusedPanel
	width         int
	height        int
	workspacePath string
}

func NewBrowserScreen(a *app.App) *BrowserScreen {
	km := components.DefaultKeyMap
	s := &BrowserScreen{
		app:           a,
		entryList:     components.NewEntryList(),
		entryMetas:    components.NewEntryMetas(),
		footer:        components.NewFooter(),
		help:          components.NewHelpScreen(km),
		focus:         focusList,
		workspacePath: a.WorkspacePath(),
	}
	s.entryList.SetFocused(true)
	s.updateFooter()
	return s
}

func (m *BrowserScreen) Init() tea.Cmd {
	return m.loadEntries(0)
}

func (m *BrowserScreen) loadEntries(offset int) tea.Cmd {
	limit := m.entryList.PageSize()
	return func() tea.Msg {
		total, err := m.app.EntryRepo().Count()
		if err != nil {
			return entriesLoadedMsg{err: err}
		}
		entries, err := m.app.EntryRepo().List(entryrepo.ListOptions{Limit: limit, Offset: offset})
		return entriesLoadedMsg{entries: entries, total: total, err: err}
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
			m.entryList.SetEntries(msg.entries, msg.total)
			if entry := m.entryList.SelectedEntry(); entry != nil {
				m.entryMetas.SetEntryID(entry.ID)
				return m.loadMetas(entry.ID)
			}
		}

	case components.EntryPageChangedMsg:
		return m.loadEntries(msg.Offset)

	case metasLoadedMsg:
		if msg.err == nil {
			m.entryMetas.SetMetas(msg.metas)
		}

	case metasReloadMsg:
		if msg.err == nil {
			return m.loadMetas(msg.entryID)
		}

	case components.ToastShowMsg:
		return components.GlobalToast.Show(msg.Message, msg.Level, 4*time.Second)

	case components.EntrySelectedMsg:
		m.entryMetas.SetEntryID(msg.Entry.ID)
		return m.loadMetas(msg.Entry.ID)

	case components.MetaInputSubmitMsg:
		entry := m.entryList.SelectedEntry()
		if entry == nil {
			return nil
		}
		return m.saveMeta(entry.Path, msg.Name, msg.Value)

	case components.MetaOpenEditorMsg:
		return m.openEditor(msg.EntryID, msg.Name, msg.CurrentValue)

	case metaEditorDoneMsg:
		if msg.err != nil {
			return components.GlobalToast.Show(msg.err.Error(), components.ToastError, 4*time.Second)
		}
		entry := m.entryList.SelectedEntry()
		if entry == nil {
			return nil
		}
		return m.saveMeta(entry.Path, msg.name, msg.value)

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
		case key.Matches(msg, components.DefaultKeyMap.SyncEntry):
			entry := m.entryList.SelectedEntry()
			if entry != nil {
				return m.syncEntry(entry.Path)
			}
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

type metasReloadMsg struct {
	entryID int64
	err     error
}

// saveMeta writes via the metadata package, syncs, then reloads.
func (m *BrowserScreen) saveMeta(entryPath, name, value string) tea.Cmd {
	return func() tea.Msg {
		info, err := m.app.Drive().Get(entryPath)
		if err != nil {
			return components.ToastShowMsg{Message: fmt.Sprintf("drive error: %v", err), Level: components.ToastError}
		}

		handlers := metadata.GetHandlers(info)
		success, err := metadata.Set(info, name, value, handlers)
		if err != nil {
			return components.ToastShowMsg{Message: fmt.Sprintf("set error: %v", err), Level: components.ToastError}
		}
		if !success {
			return components.ToastShowMsg{Message: fmt.Sprintf("cannot set '%s' on this file type", name), Level: components.ToastError}
		}

		if err := m.app.Syncer().One(entryPath); err != nil {
			return components.ToastShowMsg{Message: fmt.Sprintf("sync error: %v", err), Level: components.ToastError}
		}

		entry, err := m.app.EntryRepo().GetByPath(entryPath)
		if err != nil {
			return components.ToastShowMsg{Message: fmt.Sprintf("reload error: %v", err), Level: components.ToastError}
		}
		return metasReloadMsg{entryID: entry.ID}
	}
}

// syncEntry runs sync.One then reloads metas for the entry.
func (m *BrowserScreen) syncEntry(entryPath string) tea.Cmd {
	return func() tea.Msg {
		if err := m.app.Syncer().One(entryPath); err != nil {
			return components.ToastShowMsg{Message: fmt.Sprintf("sync error: %v", err), Level: components.ToastError}
		}
		entry, err := m.app.EntryRepo().GetByPath(entryPath)
		if err != nil {
			return components.ToastShowMsg{Message: fmt.Sprintf("reload error: %v", err), Level: components.ToastError}
		}
		return components.BatchMsg{
			Msgs: []tea.Msg{
				components.ToastShowMsg{Message: "synced " + entryPath, Level: components.ToastSuccess},
				metasReloadMsg{entryID: entry.ID},
			},
		}
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
	shared := []key.Binding{km.Up, km.Down, km.FocusNext, km.SyncEntry, km.Help, km.Quit}
	if m.focus == focusMetas {
		metaBindings := []key.Binding{km.MetaAdd, km.MetaReplace, km.MetaUpdate, km.MetaEditor}
		m.footer.SetBindings(append(shared[:3], append(metaBindings, shared[3:]...)...)...)
	} else {
		pageBindings := []key.Binding{km.PagePrev, km.PageNext}
		m.footer.SetBindings(append(shared[:2], append(pageBindings, shared[2:]...)...)...)
	}
}

func (m *BrowserScreen) updateSizes() {
	const headerHeight = 1
	const footerHeight = 1
	contentHeight := m.height - headerHeight - footerHeight

	listWidth := m.width / 3
	metasWidth := m.width - listWidth

	m.entryList.SetSize(listWidth, contentHeight)
	m.entryMetas.SetSize(metasWidth, contentHeight)
	m.footer.SetWidth(m.width)
	m.help.SetSize(m.width, m.height)
}

var (
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("33")).
			Padding(0, 1)

	headerDimStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("244")).
			Background(lipgloss.Color("235")).
			Padding(0, 1)
)

// Render returns the screen content as a plain string (no tea.View wrapper).
// The root model uses this to composite overlays before creating the final View.
func (m *BrowserScreen) Render() string {
	if m.showHelp {
		return m.help.View()
	}

	const headerHeight = 1
	const footerHeight = 1
	contentHeight := m.height - headerHeight - footerHeight

	// Header: workspace path
	icon := "󰉋 "
	header := lipgloss.NewStyle().
		Width(m.width).
		Render(headerStyle.Render(icon+"Atlas") + m.workspacePath)

	mainArea := lipgloss.NewStyle().Width(m.width).Height(contentHeight).Render(
		lipgloss.JoinHorizontal(lipgloss.Top, m.entryList.View(), m.entryMetas.View()),
	)

	return lipgloss.JoinVertical(lipgloss.Left, header, mainArea, m.footer.View())
}

func (m *BrowserScreen) Title() string { return "󰉋 Browser" }
