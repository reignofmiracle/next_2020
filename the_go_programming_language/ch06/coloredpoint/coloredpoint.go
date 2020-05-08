package coloredpoint

import (
	"image/color"

	"../geometry"
)

// ColoredPoint ...
type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}
