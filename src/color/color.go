package convertcolor

import (
	"errors"
	"math"
	"reflect"
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
	FormatFloatToInt() map[string]int
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

// SaveColor save a color
// Return error
func (c RgbColor) SaveColor() error {
	if len(store) < 6 {
		store = append(store, c)

		return nil
	}

	return errors.New("the store is already full")
}
