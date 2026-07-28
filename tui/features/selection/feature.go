package selection

import "github.com/sidekick-coder/atlas/tui/features/event"

type ChangeEvent struct {
	Old int
	New int
}

type Feature struct {
	cursor  int
	total   int
	enabled bool
	loop    bool

	Change *event.Event[ChangeEvent]
}

func Create() *Feature {
	return &Feature{
		cursor:  -1,
		total:   0,
		enabled: true,
		loop:    false,

		Change: event.Create[ChangeEvent](),
	}
}

func (f *Feature) GetTotal() int {
	return f.total
}

func (f *Feature) SetTotal(total int) {
	f.total = total

	if f.cursor >= total {
		f.cursor = total - 1
	}

	if f.cursor < 0 && total > 0 {
		f.cursor = -1
	}
}

func (f *Feature) GetCursor() int {
	return f.cursor
}

func (f *Feature) SetCursor(cursor int) {
	oldCursor := f.cursor 

	f.cursor = cursor

	f.Change.Emit(ChangeEvent{
		Old: oldCursor,
		New: cursor,
	})
}

func (f *Feature) IsSelected(index int) bool {
	return f.cursor == index
}

func (f *Feature) GetNextIndex() int {
	isLast := f.cursor == f.total-1

	if isLast && !f.loop {
		return f.cursor
	}

	if f.cursor < 0 {
		return 0
	}

	if isLast {
		return 0
	}

	return f.cursor + 1
}

func (f *Feature) Next() {
	f.SetCursor(f.GetNextIndex())
}

func (f *Feature) GetPrevIndex() int {
	isFirst := f.cursor == 0

	if isFirst && !f.loop {
		return f.cursor
	}

	if f.cursor < 0 {
		return 0
	}

	if isFirst {
		return f.total - 1
	}

	return f.cursor - 1
}

func (f *Feature) Prev() {
	f.SetCursor(f.GetPrevIndex())
}

func (f *Feature) Clear() {
	f.SetCursor(-1)
}
