package convertcolor

import (
	"errors"
	"strconv"
)

// Hex is a struct referencing an hexa decimal color value
type Hex string

// ConvertRGBtoHexa Convert an RGB to an Hexa value
// Return string
func (c RgbColor) ConvertRGBtoHexa() (Hex, error) {
	var (
		hexValue Hex
	)

	rInt := int64(c.Red)
	gInt := int64(c.Green)
	bInt := int64(c.Blue)

	RGBArray := []int64{rInt, gInt, bInt}

	for i := 0; i < len(RGBArray); i++ {
		// just casting our type hex string from a string..
		hexValue += Hex(strconv.FormatInt(RGBArray[i], 16))
	}

	if len(hexValue) > 6 {
		return hexValue, errors.New("something bad happened while generating the Hex")
	}

	return hexValue, nil
}

// ToRGB convert an hexa to RGB, we use a special type in order to have the same structure for using the package
// Return RgbColor || error
func (h Hex) ToRGB() (RgbColor, error) {
	var (
		red   uint8
		green uint8
		blue  uint8
	)
	// Split the hexa
	var hexaMap = make(map[string]string)
	hexaMap["red"] = string(h)[:2]
	hexaMap["green"] = string(h)[2:4]
	hexaMap["blue"] = string(h)[4:6]

	for i, value := range hexaMap {
		tmpVal, err := strconv.ParseInt(value, 16, 32)

		if err != nil {
			return RgbColor{}, err
		}

		switch i {
		case "red":
			red = uint8(tmpVal)
			break
		case "green":
			green = uint8(tmpVal)
			break
		case "blue":
			blue = uint8(tmpVal)
			break
		default:
			return RgbColor{}, errors.New("Missing hex value")
		}
	}

	return RgbColor{
		Red:   red,
		Green: green,
		Blue:  blue,
	}, nil
}
