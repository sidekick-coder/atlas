package layer

import "github.com/sidekick-coder/atlas/internal/utils"

type Layer struct {
	ID string
	Render func() string
	X int
	Y int
}

func Create() *Layer {
	id, err := utils.CreateID()

	if err != nil {
		panic(err)
	}

	return &Layer{
		ID: id,
	}
}

func (l *Layer) SetID(id string) {
	l.ID = id
}

func (l *Layer) SetRender(renderFunc func() string) {
	l.Render = renderFunc
}

func (l *Layer) SetPosition(x, y int) {
	l.Y = y
	l.X = x
}

