package convertcolor

import (
	"errors"
	"strconv"
)

// YCbCr struct Object
type YCbCr struct {
	Y  float64
	Cb float64
	Cr float64
}

// ConvertYCbCr convert an RGB color to an YCbCr
// Return YCbCr
func (c RgbColor) ConvertYCbCr() *YCbCr {
	castValue := map[string]float64{"red": float64(c.Red), "green": float64(c.Green), "blue": float64(c.Blue)}

	// assign the map value
	// we make the digital ycbcr conversion not the analog one
	color := &YCbCr{
		Y:  0.299*castValue["red"] + 0.587*castValue["green"] + 0.114*castValue["blue"],
		Cb: -0.1687*castValue["red"] - 0.3313*castValue["green"] + 0.5*castValue["blue"] + 128,
		Cr: 0.5*castValue["red"] - 0.418688*castValue["green"] - 0.0813*castValue["blue"] + 128,
	}

	return color
}

// ToRGB convert an YCbCr to an RgbColor Object
// Return RgbColor
func (y *YCbCr) ToRGB() RgbColor {
	color := RgbColor{
		Red:   uint8(y.Y + 1.402*(y.Cr-128)),
		Green: uint8(y.Y - 0.34414*(y.Cb-128) - 0.71414*(y.Cr-128)),
		Blue:  uint8(y.Y + 1.772*(y.Cb-128)),
	}

	return color
}

// FormatFloatToInt - Format a float ycbcr to an int ycbcr
func (y *YCbCr) FormatFloatToInt() (map[string]int, error) {
	yF, eF := strconv.Atoi(strconv.FormatFloat(y.Y, 'f', 0, 64))
	yCb, eCb := strconv.Atoi(strconv.FormatFloat(y.Cb, 'f', 0, 64))
	yCr, eCr := strconv.Atoi(strconv.FormatFloat(y.Cr, 'f', 0, 64))

	if eF != nil || eCb != nil || eCr != nil {
		return nil, errors.New("unable to convert ycbcr float value to an int")
	}

	ycbcrInt := map[string]int{
		"Y":  yF,
		"Cb": yCb,
		"Cr": yCr,
	}

	return ycbcrInt, nil
}
