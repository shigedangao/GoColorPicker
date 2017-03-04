package colorHelper

import (
	"errors"
	"fmt"
	"math"
)

type hsv struct {
	h int
	s float64
	v float64
}

// RgbToHsv
//	* Convert an RGB Color to HSV
//	* HSV represent Hue, Saturation, Value
// --> (c rgbColor)
func (c rgbColor) RgbToHsv() (error, hsv) {
	// first get the min max value
	color := []float64{float64(c.red), float64(c.green), float64(c.blue)}
	min, max := getMinMax(color)

	// creating the HSV value
	maxColor := getMaxColor(max, color)
	fmt.Println(maxColor)
	hueStruct, err := calcHue(maxColor, color, max, min)

	fmt.Println("hue str", hueStruct)

	if err != nil {
		return errors.New("An error happened while converting the RGB Color to HSV"), hsv{}
	}

	hsvStruct := hsv{
		h: hueStruct.Angle,
		s: 1 - min/max,
		v: max / 255,
	}

	return nil, hsvStruct
}

func (h hsv) HsvToRgb() (error, *rgbColor) {
	max := 255 * h.v
	min := max * (1 - h.s)

	// We need to calculate a Z value
	// Source from : http://www.had2know.com/technology/hsv-rgb-conversion-formula-calculator.html
	z := (max - min) * (1 - (math.Mod((float64(h.h)/60), 2) - 1))
	rgb := h.calcRgbFromHsv(max, min, z)

	if rgb == nil {
		return errors.New("An error happened while converting hsv to rgb"), nil
	}

	return nil, rgb
}

func (h hsv) calcRgbFromHsv(max float64, min float64, z float64) *rgbColor {

	var (
		red   float64
		green float64
		blue  float64
	)

	switch {
	case h.h >= 0 && h.h < 60:
		red = max
		green = z + min
		blue = min
	case h.h >= 60 && h.h < 120:
		red = z + min
		green = max
		blue = min
	case h.h >= 120 && h.h < 180:
		red = min
		green = max
		blue = z + min
	case h.h >= 180 && h.h < 240:
		red = min
		green = z + min
		blue = max
	case h.h >= 240 && h.h < 300:
		red = z + min
		green = min
		blue = max
	case h.h >= 300 && h.h < 360:
		red = max
		green = min
		blue = z + min
	}

	rgb := &rgbColor{
		red:   uint8(red),
		green: uint8(green),
		blue:  uint8(blue),
	}

	return rgb
}
