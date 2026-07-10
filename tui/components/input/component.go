package input

type Component struct {
	buf     []rune
	cursor  int

	width   int
	height  int

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

func (i *Input) Enable() *Input {
	i.enabled = true
	return i
}

func (i *Input) Disable() *Input {
	i.enabled = false
	return i
}

func (i *Input) SetInitialValue(initialValue string) {
	i.buf = []rune(initialValue)
	i.cursor = len(i.buf)
}

func (i *Input) GetValue() string {
	return string(i.buf)
}
