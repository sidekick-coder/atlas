package screen

type SizeMsg struct {
	Width  int
	Height int
}

type AddMsg struct {
	Name string 
	Options map[string]any
}

type RemoveMsg struct {
	Index int
}

type ReplaceMsg struct {
	Index int
	Name  string
	Options map[string]any
}

type ReplaceCurrentMsg struct {
	Name  string
	Options map[string]any
}
