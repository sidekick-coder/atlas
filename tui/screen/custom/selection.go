package custom

import "github.com/sidekick-coder/atlas/tui/screen/custom/component"

func (s *Screen) GetCurrent() (*component.Component, bool) {
	if len(s.components) == 0 {
		return nil, false
	}

	index := s.selection.GetCursor()
	if index < 0 || index >= len(s.components) {
		return nil, false
	}

	return &s.components[index], true
}

func (s *Screen) Select(index int) {
	if index < 0 || index >= len(s.components) {
		return
	}

	if oc, ok := s.GetCurrent(); ok {
		oc.Definition.OnBlur()
	}

	s.selection.SetCursor(index)

	if cc, ok := s.GetCurrent(); ok {
		cc.Definition.OnFocus()
	}
}

func (s *Screen) Next() {
	s.Select(s.selection.GetNextIndex())
}

func (s *Screen) Prev() {
	s.Select(s.selection.GetPrevIndex())
}
