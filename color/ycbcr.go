package colorHelper

type YCbCr struct {
	Y  float32
	Cb float32
	Cr float32
}

// convertYCbCr
//      * Convert an RGB value to YCbCr
// --> (c rgbColor)
// @ YCbCr
func (c rgbColor) ConvertYCbCr() YCbCr {
	castValue := []float32{float32(c.red), float32(c.green), float32(c.blue)}
	color := YCbCr{
		Y:  0.299*castValue[0] + 0.587*castValue[1] + 0.114*castValue[2],
		Cb: -0.1687*castValue[0] - 0.3313*castValue[1] + 0.5*castValue[2] + 128,
		Cr: 0.5*castValue[0] - 0.4187*castValue[1] - 0.0813*castValue[2] + 128,
	}

	return color
}

// convertToRGB
//      * Conver the YCbCr to RGB
// --> (y YCbCr)
// @ rgbColor
func (y YCbCr) ConvertToRGB() rgbColor {
	color := rgbColor{
		red:   uint8(y.Y + 1.402*(y.Cr-128)),
		green: uint8(y.Y - 0.34414*(y.Cb-128) - 0.71414*(y.Cr-128)),
		blue:  uint8(y.Y + 1.772*(y.Cb-128)),
	}

	return color
}
