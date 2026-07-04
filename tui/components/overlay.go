package components

import (
	"strings"

	lipgloss "charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
)

// PlaceOverlay centers box over bg using line-by-line ANSI-aware compositing.
// bg should be the fully-rendered background string.
// box is the raw dialog box (no padding around it).
func PlaceOverlay(box, bg string, width, height int) string {
	bgLines := strings.Split(bg, "\n")
	boxLines := strings.Split(box, "\n")

	boxH := len(boxLines)
	boxW := lipgloss.Width(boxLines[0])

	startY := (height - boxH) / 2
	startX := (width - boxW) / 2
	if startY < 0 {
		startY = 0
	}
	if startX < 0 {
		startX = 0
	}

	// Ensure bg has enough lines.
	for len(bgLines) < height {
		bgLines = append(bgLines, strings.Repeat(" ", width))
	}

	for i, boxLine := range boxLines {
		y := startY + i
		if y >= len(bgLines) {
			break
		}
		bgLines[y] = overlayLine(bgLines[y], boxLine, startX, width)
	}

	return strings.Join(bgLines, "\n")
}

// overlayLine replaces the visual segment [x, x+fgWidth) in bgLine with fgLine.
func overlayLine(bgLine, fgLine string, x, totalW int) string {
	fgW := ansi.StringWidth(fgLine)

	// Pad short bg lines so we always have enough room.
	bgW := ansi.StringWidth(bgLine)
	if bgW < totalW {
		bgLine += strings.Repeat(" ", totalW-bgW)
	}

	before := ansi.Truncate(bgLine, x, "")
	after := ansi.TruncateLeft(bgLine, x+fgW, "")

	// If before is shorter than x (e.g. double-width chars), pad it.
	beforeW := ansi.StringWidth(before)
	if beforeW < x {
		before += strings.Repeat(" ", x-beforeW)
	}

	return before + fgLine + after
}
