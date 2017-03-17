package colorHelper

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

// Outercolor define a random property of a color to be based on ...
// We don't use the pointer as we does not play with the type for the moment...
type rgbColor struct {
	red   uint8
	green uint8
	blue  uint8
}

type ConvertToRGB interface {
	HexaToRgb() rgbColor
	CymkToRgb() rgbColor
	HsvToRgb() (*rgbColor, error)
	ConvertToRGB() rgbColor
}

// Create a nil hue store, if the user does not it, it's still nil and not allocated in the memory
// If we use make then it will have a place in hte memory
var (
	store   []rgbColor
	hueSqrt = math.Sqrt(3)
)

// define the constant that we need threw the package

// MakeColorFromInput ...
func MakeColorFromInput(r uint8, g uint8, b uint8) rgbColor {
	sample := rgbColor{
		red:   r,
		green: g,
		blue:  b,
	}

	return sample
}

// UpdateCurrentColor
//      * Create and update a new color
//      * Should we compare using the Type interface instead of convert the type into a String ?
// --> nr Uint8
// --> ng Uint8
// --> nb Uint8
// @ Error & rgbColor
func (c rgbColor) UpdateCurrentColor(nr uint8, ng uint8, nb uint8) (error, rgbColor) {

	if reflect.TypeOf(nr).String() != "uint8" {
		return errors.New("red value is not in the valid format between 0 - 255"), rgbColor{}
	}

	c.red = nr
	c.green = ng
	c.blue = nb

	return nil, c
}

// ConvertRGBtoHexa
//		* Convert an RGB Color into HEXA
// --> (c rgbColor)
// @ String
func (c rgbColor) ConvertRGBtoHexa() string {
	var hexValue string
	// Before converting our value to String
	// We need to cast our Uint8 -> Int64
	rInt := int64(c.red)
	gInt := int64(c.green)
	bInt := int64(c.blue)

	// We directly create an array as we're not going to manipulate or extending the array (so no slice)
	RGBArray := []int64{rInt, gInt, bInt}

	for i := 0; i < len(RGBArray); i++ {
		// convert the Int64 value into the hexa
		// Since hex is an integer literal we can use strconv to convert it... (thanks stackoverflow)
		hexValue += strconv.FormatInt(RGBArray[i], 16)
	}

	return hexValue
}

// Hexa To Rgb
// 		* Convert an hexa value to an Rgb value
// --> hex string
// @ rgbColor{}
func HexaToRgb(hex string) rgbColor {
	// Split the hexa
	var hexaMap = make(map[string]string)
	hexaMap["red"] = hex[:2]
	hexaMap["green"] = hex[2:4]
	hexaMap["blue"] = hex[4:6]

	// convert the value into a uint 16
	red, _ := strconv.ParseInt(hexaMap["red"], 16, 32)
	green, _ := strconv.ParseInt(hexaMap["green"], 16, 32)
	blue, _ := strconv.ParseInt(hexaMap["blue"], 16, 32)

	return rgbColor{
		red:   uint8(red),
		green: uint8(green),
		blue:  uint8(blue),
	}
}

// SaveColor
//		* Save the color into an array of store color
// --> (c rgbColor)
func (c rgbColor) SaveColor() error {
	// populate the slice
	// we allow the user to create a panel of color ranging from 0 color to 6
	if len(store) < 6 {
		store = append(store, c)

		return nil
	}

	return errors.New("the store is already full")
}
