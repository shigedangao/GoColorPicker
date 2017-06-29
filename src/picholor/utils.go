package picholor

import "errors"

// Get Max Uint8
func getMax(colors map[string]uint8) (uint8, error) {

	if len(colors) == 0 {
		return 0, errors.New("Colors array cannot be empty for GetMax")
	}

	// Set max to index 0
	max := colors["red"]

	for _, d := range colors {
		if d > max {
			max = d
		}
	}

	return max, nil
}

// Get Min Uint8
func getMin(colors map[string]uint8) (uint8, error) {
	if len(colors) == 0 {
		return 0, errors.New("Colors array cannot be empty for GetMin")
	}

	min := colors["red"]

	for _, d := range colors {
		if d < min {
			min = d
		}
	}

	return min, nil
}
