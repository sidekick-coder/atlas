package selection

type Feature struct {
	cursor  int
	total   int
	enabled bool
}

func Create() *Feature {
	return &Feature{
		cursor:  -1,
		total:   0,
		enabled: true,
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
}

func (f *Feature) GetCursor() int {
	return f.cursor
}

func (f *Feature) SetCursor(cursor int) {
	f.cursor = cursor
}

func (f *Feature) IsSelected(index int) bool {
	return f.cursor == index
}

func (f *Feature) Next() {
	isLast := f.cursor == f.total-1

	if isLast {
		f.cursor = 0
		return
	}

	f.cursor++
}

func (f *Feature) Prev() {
	isFirst := f.cursor == 0

	if isFirst {
		f.cursor = f.total- 1
		return
	}

	f.cursor--
}

func (f *Feature) Clear() {
	f.cursor = 0
}
