package picholor

import (
	"encoding/json"
)

type Cmyk struct {
	C uint8 `json:"C"`
	M uint8 `json:"M"`
	Y uint8 `json:"Y"`
	K uint8 `json:"K"`
}

func (c Cmyk) buildFromJSON(data []byte) (Cmyk, []byte) {
	var cmyk Cmyk
	err := json.Unmarshal(data, cmyk)

	exception := cmyk.handleError(err)
	return cmyk, exception
}

func (c Rgb) buildCMYKFromRgb() Cmyk {
	colors := map[string]uint8{
		"red":   c.R,
		"green": c.G,
		"blue":  c.B,
	}

	// Get the max color value
	max, err := getMax(colors)

	if err != nil {

	}

	k := 1 - float64(max)

	cmyk := Cmyk{
		C: calculateTypeValue(colors["red"], k),
		M: calculateTypeValue(colors["green"], k),
		Y: calculateTypeValue(colors["blue"], k),
		K: uint8(k),
	}

	return cmyk
}

func calculateTypeValue(v uint8, k float64) uint8 {
	nV := (1 - float64(v) - k) / (1 - k)

	return uint8(nV * 100)
}

func (c Cmyk) toRGB() (Rgb, error) {

	return Rgb{
		R: uint8(255 * (1 - c.C) * (1 - c.K)),
		G: uint8(255 * (1 - c.M) * (1 - c.K)),
		B: uint8(255 * (1 - c.Y) * (1 - c.K)),
	}, nil
}

func (c Cmyk) handleError(e error) []byte {
	if e == nil {
		return nil
	}

	return []byte("picholor error type `CMYK`: " + e.Error())
}
