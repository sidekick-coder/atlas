package layer

import (
	"log/slog"

	"charm.land/lipgloss/v2"
)

var layers []*Layer
var ScreenWidth int
var ScreenHeight int

func Add(l *Layer) {
	layers = append(layers, l)
	slog.Info("Adding layer", "id", l.ID, "x", l.X, "y", l.Y, "z", l.Z)
}

func Remove(layer *Layer) {
	for i, l := range layers {
		if l.ID == layer.ID {
			layers = append(layers[:i], layers[i+1:]...)
			slog.Info("Remove layer", "id", l.ID, "x", l.X, "y", l.Y, "z", l.Z)
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

		lipglossLayer := lipgloss.NewLayer(l.Render()).X(l.X).Y(l.Y).Z(l.Z)

		result = append(result, lipglossLayer)
	}

	return result
}

func SetScreenSize(width, height int) {
	ScreenWidth = width
	ScreenHeight = height
}
