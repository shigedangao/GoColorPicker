package picholor

import "encoding/json"

type Rgb struct {
	R uint8 `json:"red"`
	G uint8 `json:"green"`
	B uint8 `json:"blue"`
}

// Build From JSON
func (r Rgb) buildFromJSON(data []byte) (Rgb, []byte) {
	var color Rgb

	// Unmarshal the rgb color
	err := json.Unmarshal(data, color)
	exception := color.handleError(err)

	return color, exception
}

// Handle RGB Error
func (r Rgb) handleError(e error) []byte {
	if e == nil {
		return nil
	}

	return []byte("picholor error type `RGB` :" + e.Error())
}

// Create RGB From Other existing type
func makeRgbFromOtherType(t common) (Rgb, []byte) {
	color, e := t.toRGB()
	exception := color.handleError(e)

	if exception != nil {
		return Rgb{}, exception
	}

	return color, nil
}
