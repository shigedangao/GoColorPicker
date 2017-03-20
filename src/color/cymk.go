package colorHelper

// Cymk Object
type Cymk struct {
	C float64
	Y float64
	M float64
	K float64
}

// RgbToCymk convert an RGB value to an Cymk
// Return Cymk
func (c RgbColor) RgbToCymk() Cymk {
	color := []float64{float64(c.Red) / 255, float64(c.Green) / 255, float64(c.Blue) / 255}

	_, max := getMinMax(color)

	k := 1 - max
	printColor := Cymk{
		C: float64((1 - color[0] - k) / (1 - k)),
		M: float64((1 - color[1] - k) / (1 - k)),
		Y: float64((1 - color[2] - k) / (1 - k)),
		K: float64(k),
	}

	return printColor
}

// ToRGB convert a Cymk to an RgbColor Object
func (y Cymk) ToRGB() RgbColor {
	color := RgbColor{
		Red:   uint8(255 * (1 - y.C) * (1 - y.K)),
		Green: uint8(255 * (1 - y.M) * (1 - y.K)),
		Blue:  uint8(255 * (1 - y.Y) * (1 - y.K)),
	}

	return color
}
