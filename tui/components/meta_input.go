package components

// MetaInputSubmitMsg is emitted when a meta value edit is confirmed.
type MetaInputSubmitMsg struct {
	EntryID int64
	Name    string
	Value   string
}

// MetaOpenEditorMsg is emitted when the user requests to open $EDITOR.
type MetaOpenEditorMsg struct {
	EntryID      int64
	Name         string
	CurrentValue string
}
