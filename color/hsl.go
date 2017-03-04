package colorHelper

import (
	"errors"
	"fmt"
)

// Define the hue and it's function
type HslStruct struct {
	Angle      int
	Saturation float64
	Luminace   float64
}

// rgbToHue
// 		* Based function to convert a RGB to HUE
// --> (c rgbColor)
// @ *HueStruct (pointer)
func (c rgbColor) RgbToHsl() *HslStruct {
	var s float64
	// Calcul step from Niwa.nu

	// determinate the min max value
	rV := float64(c.red) / 255
	gV := float64(c.green) / 255
	bV := float64(c.blue) / 255

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
		s = (max - min) / (2 - (max - min))
	}

	fmt.Println("saturation equal to ", s)

	// Get the color which it's value is higher than the other
	maxColorName := getMaxColor(max, colorFloatArray)
	// Calculate the hue
	hue, _ := calcHue(maxColorName, colorFloatArray, max, min)

	// filling our struct with the rest of the parameters
	hue.Saturation = s
	hue.Luminace = luminace

	return hue
}

// GetHSL
//      * Get the HSL value from the HUE
// @ error
func (h *HslStruct) GetHSL() (error, []float64) {

	// if the luminace and the satuaration is equal to 0 there's an error somewhere...
	if h.Luminace == 0 || h.Saturation == 0 {
		// Dereference our pointer as the value is wrong
		*h = HslStruct{}

		return errors.New("the hue struct miss some datas deferencing the pointer"), nil
	}

	hsl := calcHSL(h.Luminace, h.Saturation, h.Angle)

	return nil, hsl
}

// calcHSL
//		* Calculate the HSL
//		* @TODO Create an own structure for the HSL
// --> l float64
// --> sat float64
// --> hueAngle int
func calcHSL(l float64, sat float64, hueAngle int) []float64 {
	var (
		tmpHSL     float64
		tempColors []float64
		hsl        []float64
	)

	// We convert the data to float32 in order to remove the memory footprint
	if l < 0.5 {
		tmpHSL = l * (1 + sat)
	} else {
		tmpHSL = l + sat - l*sat
	}

	tempHSLsecond := 2*l - tmpHSL
	hue := float64(hueAngle) / 360

	// create temporary value for each red blue and green value
	// Grow our slice
	tempColors = append(tempColors, hue+0.333, hue, hue-0.333)

	// now we need to check whenever the value are positive or negatvie or above 1

	for _, color := range tempColors {
		if color > 1 {
			color--
		} else {
			color++
		}
	}

	CalculateHSLFunc := getRightFormula(tempColors, tmpHSL, tempHSLsecond)
	hsl = CalculateHSLFunc()

	return hsl
}

// GetRightFormula
// 		* Calculate the HSl using the right formula
// --> tempColor []float64
// --> tempHSL []float64
// --> templHSLSecond []float64
// @ func, []float64
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

// Percent
// 		* Convert the raw float to a %
//		* As there's only these function as common i don't use interfaces.. should i ?
// --> (h *HueStruct)
// @ int, error
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
