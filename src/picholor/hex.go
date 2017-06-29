package picholor

import (
	"encoding/json"
	"strconv"
)

type Hex string

func (h Hex) buildFromJSON(data []byte) (Hex, []byte) {
	var hex Hex
	err := json.Unmarshal(data, hex)

	exception := hex.handleError(err)

	return hex, exception
}

func (h Hex) handleError(e error) []byte {
	if e != nil {
		return []byte(e.Error())
	}

	return nil
}

func (r Rgb) buildHex() Hex {

	var hex Hex
	colors := map[string]uint8{
		"red":   r.R,
		"green": r.G,
		"blue":  r.B,
	}

	for _, d := range colors {
		hex += Hex(strconv.FormatUint(uint64(d), 10))
	}

	return hex
}

func (h Hex) toRGB() {

}
