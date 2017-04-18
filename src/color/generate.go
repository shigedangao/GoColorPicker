package convertcolor

import (
	"errors"
)

// GenerateShadeTint generate a shade or a tint
// Params factor : int && genType : String
// Return []RgbColor || error
func (c RgbColor) GenerateShadeTint(factor int, genType string) ([]RgbColor, error) {

	colors := c.calculateColors(factor, genType)

	if colors == nil {
		return nil, errors.New("An error happened while generating the colors")
	}

	return colors, nil
}

// calculateColors calculate the shade of colors
// (!) This is an unexported function link to the RgbColor Object
// Params factor : Int && genType string
// Return []RgbColor
func (c RgbColor) calculateColors(factor int, genType string) []RgbColor {
	var (
		colors = make([]RgbColor, uint8(factor))
	)

	for i := 0; i < factor; i++ {
		if genType == "tint" {
			colors[i] = RgbColor{
				Red:   c.Red + uint8(getTint(c.Red, i, factor)),
				Green: c.Green + uint8(getTint(c.Green, i, factor)),
				Blue:  c.Blue + uint8(getTint(c.Blue, i, factor)),
			}
		} else if genType == "shade" {
			colors[i] = RgbColor{
				Red:   c.Red + uint8(getShade(c.Red, i, factor)),
				Green: c.Green + uint8(getShade(c.Green, i, factor)),
				Blue:  c.Blue + uint8(getShade(c.Blue, i, factor)),
			}
		} else {
			return nil
		}
	}

	return colors
}

// getTint return a tint of float32
// Return float32
func getTint(color uint8, i int, factor int) float32 {

	floatI := float32(i)
	return (255 - float32(color)) * (float32(floatI+1) / float32(factor))
}

// getShade return a shade value of float32
// Return float32
func getShade(color uint8, i int, factor int) float32 {
	floatI := float32(i)
	floatFactor := float32(factor)

	return float32(color) * (float32(floatI+1) / floatFactor)
}
