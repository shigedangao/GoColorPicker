package colorHelper

import (
	"errors"
	"fmt"
	"math"
)

// HslStruct Object
type HslStruct struct {
	Angle      int
	Saturation float64
	Luminace   float64
}

// RgbToHsl convert an RGB to an HSL value
// Return *HslStruct (pointer)
// Those calculation are based on an article published on niwa.nu
func (c RgbColor) RgbToHsl() *HslStruct {
	var s float64

	// determinate the min max value
	rV := float64(c.Red) / 255
	gV := float64(c.Green) / 255
	bV := float64(c.Blue) / 255

	colorFloatArray := []float64{rV, gV, bV}
	min, max := getMinMax(colorFloatArray)

	// If min and max is equal it mean that there's no saturation
	// Therefore no hue
	if min == max {
		return nil
	}

	luminace := (min + max) / 2
	if luminace < 0.5 {
		s = (max - min) / (max + min)
	} else {
		s = (max - min) / (2 - max - min)
	}

	// Get the color which it's value is higher than the other
	maxColorName := getMaxColor(max, colorFloatArray)
	// Calculate the hue
	hue, _ := calcHue(maxColorName, colorFloatArray, max, min)

	// filling our struct with the rest of the parameters
	hue.Saturation = s
	hue.Luminace = luminace

	return hue
}

// getRightFormula choose the right formula and calculate it's value
// Params tempColor []float64, tempHSL float64, tempHSLSecond float64
// Return []float64
func getRightFormula(tempColor []float64, tempHSL float64, tempHSLSecond float64) func() []float64 {

	colorHSLValue := make([]float64, 3)

	for i := 0; i < 3; i++ {

		if tempColor[i]*6 < 1 {
			colorHSLValue[i] = tempHSLSecond + (tempHSL-tempHSLSecond)*6*tempColor[i]
		} else if tempColor[i]*2 < 1 {
			colorHSLValue[i] = tempHSL
		} else if tempColor[i]*3 < 2 {
			colorHSLValue[i] = tempHSLSecond + (tempHSL-tempHSLSecond)*(0.666-tempColor[i])*6
		} else {
			colorHSLValue[i] = tempHSLSecond
		}
	}

	return func() []float64 {
		for idx, hsl := range colorHSLValue {
			colorHSLValue[idx] = hsl * 255
		}

		return colorHSLValue
	}
}

// Percent convert the HSL struct value to a Percent value
func (h *HslStruct) Percent(valueWanted string) (int, error) {
	if h == nil {
		return 0, errors.New("HSL is empty")
	}

	switch valueWanted {
	case "Luminace":
		return int(h.Luminace * 100), nil
	case "Saturation":
		return int(h.Saturation * 100), nil
	}

	return 0, errors.New("The value is not present withing the struct")
}

// ToRGB convert an HslStruct into an RgbColor Object
// Return RgbColor Object || error
func (h *HslStruct) ToRGB() (RgbColor, error) {
	if h == nil {
		return RgbColor{}, errors.New("Hsl struct is empty")
	}

	c := (1 - math.Abs((2*h.Luminace)-1)) * h.Saturation
	x := c * (1 - math.Abs(math.Mod(float64(h.Angle)/60, 2)-1))
	m := h.Luminace - c/2

	// get the right r'g'b'
	r, g, b := h.getRightFormula(c, x)

	fmt.Println("green", g)
	fmt.Println("m ", m)

	rgbMap := make(map[string]float64)
	rgbMap["red"] = (r + m) * 255
	rgbMap["green"] = (g + m) * 255
	rgbMap["blue"] = (b + m) * 255

	return RgbColor{
		Red:   uint8(rgbMap["red"]),
		Green: uint8(rgbMap["green"]),
		Blue:  uint8(rgbMap["blue"]),
	}, nil
}

// getRightFormula return the formula to calculate the RGB color
// Params chrome float64 || interm float64
// Return red float64, green float64, blue float64
func (h *HslStruct) getRightFormula(chrome float64, interm float64) (float64, float64, float64) {

	var (
		red   float64
		green float64
		blue  float64
	)

	switch {
	case h.Angle >= 0 && h.Angle < 60:
		red = chrome
		green = interm
		blue = 0
		break
	case h.Angle >= 60 && h.Angle < 120:
		red = interm
		green = chrome
		blue = 0
		break
	case h.Angle >= 120 && h.Angle < 180:
		red = 0
		green = chrome
		blue = interm
		break
	case h.Angle >= 180 && h.Angle < 240:
		red = 0
		green = interm
		blue = chrome
		break
	case h.Angle >= 240 && h.Angle < 300:
		red = interm
		green = 0
		blue = chrome
		break
	case h.Angle >= 300 && h.Angle < 360:
		red = chrome
		green = 0
		blue = interm
	default:
		red = 0
		green = 0
		blue = 0
	}

	return red, green, blue
}
