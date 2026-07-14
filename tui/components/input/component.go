package input

type Component struct {
	buf     []rune
	cursor  int
	width   int
	height  int
	offset  int
	enabled bool
}

type Input = Component



func Create() *Input {
	return &Input{
		width:   80,
		height:  3,
		enabled: false,
	}
}

func New() *Input {
	return Create()
}

func (i *Input) SetSize(width, height int) *Input {
	i.width = width
	i.height = height

	return i
}

func (i *Input) SetWidth(width int) *Input {
	i.width = width
	return i
}

func (i *Input) Enable() *Input {
	i.enabled = true
	i.LoadBindings()
	return i
}

func (i *Input) Disable() *Input {
	i.enabled = false
	i.UnloadBindings()
	return i
}

func (i *Input) SetValue(v string) {
	i.buf = []rune(v)
	i.cursor = len(i.buf)
}

func (i *Input) SetInitialValue(initialValue string) {
	i.buf = []rune(initialValue)
	i.cursor = len(i.buf)
}

func (i *Input) GetValue() string {
	return string(i.buf)
}
