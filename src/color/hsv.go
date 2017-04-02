package convertcolor

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// Hsv Object
type Hsv struct {
	H int
	S float64
	V float64
}

// HsvResponse Object
type HsvResponse struct {
	V *Hsv  `json:"res"`
	E error `json:"err"`
}

// RgbToHsv convert an RgbColor Object to an Hsv object pointer
// Return *Hsv || error
func (c RgbColor) RgbToHsv() (*Hsv, error) {
	// first get the min max value
	color := []float64{float64(c.Red), float64(c.Green), float64(c.Blue)}
	min, max := getMinMax(color)

	// creating the HSV value
	maxColor := getMaxColor(max, color)
	fmt.Println(maxColor)
	hueStruct, err := calcHue(maxColor, color, max, min)

	fmt.Println("hue str", hueStruct)

	if err != nil {
		return nil, errors.New("An error happened while converting the RGB Color to HSV")
	}

	hsvStruct := &Hsv{
		H: hueStruct.Angle,
		S: 1 - min/max,
		V: max / 255,
	}

	return hsvStruct, nil
}

// ToRGB Convert an Hsv Object to an Rgb one
// Return *RgbColor || error
func (h *Hsv) ToRGB() (RgbColor, error) {
	max := 255 * h.V
	min := max * (1 - h.S)

	// We need to calculate a Z value
	// Source from : http://www.had2know.com/technology/hsv-rgb-conversion-formula-calculator.html
	z := (max - min) * (1 - (math.Mod((float64(h.H)/60), 2) - 1))
	rgb := h.calcRgbFromHsv(max, min, z)

	if rgb == (RgbColor{}) {
		return RgbColor{}, errors.New("An error happened while converting hsv to rgb")
	}

	return rgb, nil
}

// calcRgbFromHsv calc Hsv value from RgbColor Object
func (h *Hsv) calcRgbFromHsv(max float64, min float64, z float64) RgbColor {

	var (
		red   float64
		green float64
		blue  float64
	)

	switch {
	case h.H >= 0 && h.H < 60:
		red = max
		green = z + min
		blue = min
	case h.H >= 60 && h.H < 120:
		red = z + min
		green = max
		blue = min
	case h.H >= 120 && h.H < 180:
		red = min
		green = max
		blue = z + min
	case h.H >= 180 && h.H < 240:
		red = min
		green = z + min
		blue = max
	case h.H >= 240 && h.H < 300:
		red = z + min
		green = min
		blue = max
	case h.H >= 300 && h.H < 360:
		red = max
		green = min
		blue = z + min
	}

	rgb := RgbColor{
		Red:   uint8(red),
		Green: uint8(green),
		Blue:  uint8(blue),
	}

	return rgb
}

// Percent convert an hsv value to a percent value given the params that we want
// Params valueWanted
// Return int || error
func (h *Hsv) Percent(valueWanted string) (int, error) {

	switch valueWanted {
	case "Saturation":
		return int(h.S * 100), nil
	case "Value":
		return int(h.V * 100), nil
	}

	return 0, errors.New("The value is not present withing the struct")
}

// FormatFloatToInt format an Hsv pointer strut data float to an Int
func (h *Hsv) FormatFloatToInt() (map[string]int, error) {
	// convert the value to an int
	s, eS := strconv.Atoi(strconv.FormatFloat(h.S*100, 'f', 0, 64))
	v, eV := strconv.Atoi(strconv.FormatFloat(h.V*100, 'f', 0, 64))

	if eS != nil || eV != nil {
		return nil, errors.New("an error happened while converting the hsv values into a map")
	}

	hsvMap := map[string]int{
		"H": h.H,
		"S": s,
		"V": v,
	}

	return hsvMap, nil
}
