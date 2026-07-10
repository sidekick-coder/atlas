package layer

import "charm.land/lipgloss/v2"

var layers []*Layer
var ScreenWidth int 
var ScreenHeight int

func Add(layer *Layer) {
	layers = append(layers, layer)
}

func Remove(layer *Layer) {
	for i, l := range layers {
		if l.ID == layer.ID {
			layers = append(layers[:i], layers[i+1:]...)
			break
		}
	}
}

func GetLipglossLayers() []*lipgloss.Layer {
	var result []*lipgloss.Layer

	for _, l := range layers {
		if l.Render == nil {
			continue
		}

		lipglossLayer := lipgloss.NewLayer(l.Render()).X(l.X).Y(l.Y)

		result = append(result, lipglossLayer)
	}

	return result
}

func SetScreenSize(width, height int) {
	ScreenWidth = width
	ScreenHeight = height
}


