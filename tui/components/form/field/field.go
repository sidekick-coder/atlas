package field

import tea "charm.land/bubbletea/v2"

type Field struct {
	Data Data 
	Definition Definition
}

func (f *Field) Render() string {
	if f.Definition == nil {
		return "Definition is nil"
	}

	return f.Definition.Render()
}

func (f *Field) Resize(width, height int) {
	if f.Definition == nil {
		return
	}

	f.Definition.Resize(width, height)
}

func (f *Field) Update(msg tea.Msg) tea.Cmd {
	if f.Definition == nil {
		return nil
	}

	return f.Definition.Update(msg)
}

func (f *Field) GetName() string {
	return f.Data.Name
}

func (f *Field) GetValue() string {
	if f.Definition == nil {
		return ""
	}

	return f.Definition.GetValue()
}
