package convertcolor

import (
	"errors"
	"math"
)

// getMinMax return the min and the max value between a floating point array
// it's use to calculate the Hue
// Params colorValue []float64
// Return float64 || float64
func getMinMax(colorValue []float64) (float64, float64) {
	var (
		min float64
		max float64
	)

	// Get the min and max value from an array of float64 value
	for i := 0; i < len(colorValue); i++ {
		for j := 0; j < len(colorValue); j++ {
			tempMin := math.Min(colorValue[i], colorValue[j])
			tempMax := math.Max(colorValue[i], colorValue[j])

			if tempMin < min || min == 0 {
				min = tempMin
			}

			if tempMax > max {
				max = tempMax
			}
		}
	}

	return min, max
}

// getMaxColor return the color which have the maximum value
// Params maxValue float64, colorValue float64
// Return string (red || green || blue)
func getMaxColor(maxValue float64, colorValue []float64) string {

	for i := 0; i < len(colorValue); i++ {
		if colorValue[i] == maxValue {
			if i == 0 {
				return "red"
			} else if i == 1 {
				return "green"
			} else {
				return "blue"
			}
		}
	}

	return ""
}

// calcHue calculate the HUE
// Param colorName string, colorValue []float64, max float64, min float64
// Return *HslStruct (pointer) || error
func calcHue(colorName string, colorValue []float64, max float64, min float64) (*HslStruct, error) {
	var hue float64
	switch colorName {
	case "red":
		hue = (colorValue[1] - colorValue[2]) / (max - min)
	case "green":
		hue = 2 + ((colorValue[2] - colorValue[0]) / (max - min))
	case "blue":
		hue = 4 + ((colorValue[0] - colorValue[1]) / (max - min))
	default:
		return nil, errors.New("no colorname provide")
	}

	hue = hue * 60

	if hue < 0 {
		hue += 360
	}

	newHue := &HslStruct{
		Angle: int(hue),
	}

	return newHue, nil
}
