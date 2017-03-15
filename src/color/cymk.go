package colorHelper

type cymk struct {
	c uint8
	y uint8
	m uint8
	k uint8
}

// RgbToCymk
//      * Convert a color from RGB to CYMK
// --> (c rgbColor)
// @ cymk
func (c rgbColor) RgbToCymk() cymk {
	color := []float64{float64(c.red) / 255, float64(c.green) / 255, float64(c.blue) / 255}

	_, max := getMinMax(color)

	k := 1 - max
	printColor := cymk{
		c: uint8((1 - color[0] - k) / (1 - k)),
		m: uint8((1 - color[1] - k) / (1 - k)),
		y: uint8((1 - color[2] - k) / (1 - k)),
		k: uint8(k),
	}

	return printColor
}

// CymkToRGB
//      * Convert a CYMK color to RGB
// --> (y cymk)
// @ rgbColor
func (y cymk) CymkToRgb() rgbColor {
	color := rgbColor{
		red:   uint8(255 * (1 - y.c) * (1 - y.k)),
		green: uint8(255 * (1 - y.m) * (1 - y.k)),
		blue:  uint8(255 * (1 - y.y) * (1 - y.k)),
	}

	return color
}
