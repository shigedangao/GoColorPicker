package convertcolor

import (
	"encoding/json"
)

// ToJSON from an Hsv value
func (h *Hsv) ToJSON() []byte {
	hsvData, e := json.Marshal(h)

	if e != nil {
		return []byte(e.Error())
	}

	return hsvData
}

// ToJSON from an Hsv value
func (c Cymk) ToJSON() []byte {
	cymkData, e := json.Marshal(c)

	if e != nil {
		return []byte(e.Error())
	}

	return cymkData
}

// ToJSON from an Ycbcr value
func (y YCbCr) ToJSON() []byte {
	ycbcrData, e := json.Marshal(y)

	if e != nil {
		return []byte(e.Error())
	}

	return ycbcrData
}

// ToJSON from an Hsl value
func (h HslStruct) ToJSON() []byte {
	hslData, e := json.Marshal(h)

	if e != nil {
		return []byte(e.Error())
	}

	return hslData
}

// ToJSON from a multiple context
func (jsonize colorJSON) ToJSON() []byte {
	jsonData, e := json.Marshal(jsonize)

	if e != nil {
		return []byte(e.Error())
	}

	return jsonData
}
