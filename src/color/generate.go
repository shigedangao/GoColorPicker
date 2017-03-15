package colorHelper

// generateTint
//      * Generate a tint of color
// --> (c rgbColor)
// --> factor float32
// --> @TODO we might need to check if the struct is empty
// @
func (c rgbColor) GenerateTint(factor float32) []rgbColor {
	// Array of shade of color based on the factor
	var (
		i    int
		tint = make([]rgbColor, uint8(factor))
	)

	for i < int(factor) {
		red := ((255 - float32(c.red)) * (float32(i+1) / factor))
		green := ((255 - float32(c.green)) * (float32(i+1) / factor))
		blue := ((255 - float32(c.blue)) * (float32(i+1) / factor))

		tint[i] = rgbColor{
			red:   c.red + uint8(red),
			green: c.green + uint8(green),
			blue:  c.blue + uint8(blue),
		}

		i++
	}

	return tint
}

// generateShade
//      * Generate a shade of color
// --> (c rgbColor)
// --> factor float32
// --> @TODO we might need to check if the struct is empty
// @
func (c rgbColor) GenerateShade(factor float32) []rgbColor {
	// Array of shade of color based on the factor
	var (
		i     int
		shade = make([]rgbColor, uint8(factor))
	)

	for i < int(factor) {
		red := float32(c.red) * (float32(i+1) / factor)
		green := float32(c.green) * (float32(i+1) / factor)
		blue := float32(c.blue) * (float32(i+1) / factor)

		shade[i] = rgbColor{
			red:   uint8(red),
			green: uint8(green),
			blue:  uint8(blue),
		}

		i++
	}

	return shade
}
