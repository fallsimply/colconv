package colconv

import (
	"fmt"
	"math"
)

// RgbToHsl takes 3 rgb digits and turn them into a hsl string
//
// Based on https://github.com/gerow/go-color
func RgbToHsl(red, green, blue uint8) string {
	var h, s float64

	//region TYPE CONVERSIONS
	// RGB INT TO FLOAT
	r := float64(red) / 255
	g := float64(green) / 255
	b := float64(blue) / 255
	//endregion

	max := max(r, g, b)
	min := min(r, g, b)
	delta := max - min
	deltaH := delta / 2
	total := max + min

	//region LUMINOSITY
	l := total / 2
	//endregion

	//region SATURATION
	switch {
	// GRAY
	case delta == 0:
		return fmt.Sprint("hsl(0, 0%, ", math.Round(l*100), "%)")
	// NOT GRAY
	case l < 0.5:
		s = delta / total
	default:
		s = delta / (2 - total)
	}
	//endregion

	//region HUE
	r2 := (((max - r) / 6) + deltaH) / delta
	g2 := (((max - g) / 6) + deltaH) / delta
	b2 := (((max - b) / 6) + deltaH) / delta
	switch {
	case r == max:
		h = b2 - g2
	case g == max:
		h = (1.0 / 3.0) + r2 - b2
	case b == max:
		h = (2.0 / 3.0) + g2 - r2
	}
	//endregion

	//region EDGE CASES
	//Wraparounds
	switch {
	case h < 0:
		h++
	case h > 1:
		h--
	}
	//endregion

	h *= 360
	s *= 100
	l *= 100

	return fmt.Sprintf("hsl(%.0f, %.0f%%, %.0f%%)", math.Round(h), math.Round(s), math.Round(l))
}
