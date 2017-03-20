package colorHelper

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

// RgbColor is an RGB Color Object
type RgbColor struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

// Color interface register a common Object that a Color can used to convert to an RGB
type Color interface {
	ToRGB() (RgbColor, error)
}

var (
	store   []RgbColor
	hueSqrt = math.Sqrt(3)
)

// MakeColorFromInput return a new Color
// Return RgbColor
func MakeColorFromInput(r uint8, g uint8, b uint8) RgbColor {
	sample := RgbColor{
		Red:   r,
		Green: g,
		Blue:  b,
	}

	return sample
}

// UpdateCurrentColor update a color and changing it's value
// Return RgbColor || error
func (c RgbColor) UpdateCurrentColor(nr uint8, ng uint8, nb uint8) (RgbColor, error) {

	if reflect.TypeOf(nr).String() != "uint8" {
		return RgbColor{}, errors.New("red value is not in the valid format between 0 - 255")
	}

	c.Red = nr
	c.Green = ng
	c.Blue = nb

	return c, nil
}

// ConvertRGBtoHexa Convert an RGB to an Hexa value
// Return string
func (c RgbColor) ConvertRGBtoHexa() string {
	var hexValue string

	rInt := int64(c.Red)
	gInt := int64(c.Green)
	bInt := int64(c.Blue)

	RGBArray := []int64{rInt, gInt, bInt}

	for i := 0; i < len(RGBArray); i++ {
		hexValue += strconv.FormatInt(RGBArray[i], 16)
	}

	return hexValue
}

// ToRGB convert an hexa to RGB
// Return RgbColor || error
func ToRGB(hex string) (RgbColor, error) {
	var (
		red   uint8
		green uint8
		blue  uint8
	)
	// Split the hexa
	var hexaMap = make(map[string]string)
	hexaMap["red"] = hex[:2]
	hexaMap["green"] = hex[2:4]
	hexaMap["blue"] = hex[4:6]

	for i, value := range hexaMap {
		tmpVal, err := strconv.ParseInt(value, 16, 32)

		if err != nil {
			return RgbColor{}, errors.New("An error happened while converting the Hex to RGB")
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

// SaveColor save a color
// Return error
func (c RgbColor) SaveColor() error {
	if len(store) < 6 {
		store = append(store, c)

		return nil
	}

	return errors.New("the store is already full")
}
