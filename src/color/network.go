package convertcolor

import (
	"encoding/json"
)

// RgbResponse makeJSONData convert an RgbStruct and it's error into a slice of a byte
func (r RgbResponse) MakeJSONData() []byte {
	data, e := json.Marshal(r)

	if e != nil {
		return []byte("e: " + e.Error())
	}

	return data
}

// HsvResponse makeJSONData convert an HsvResponse struct into a slice of a byte
func (h HsvResponse) MakeJSONData() []byte {
	data, e := json.Marshal(h)

	if e != nil {
		return []byte("e: " + e.Error())
	}

	return data
}

// HslResponse makeJSONData convert an HslResponse struct into a slice of a byte
func (h HslResponse) MakeJSONData() []byte {
	data, e := json.Marshal(h)

	if e != nil {
		return []byte("e:" + h.E.Error())
	}

	return data
}

func (st GenerateResponse) MakeJSONData() []byte {
	data, e := json.Marshal(st)

	if e != nil {
		return []byte("e:" + e.Error())
	}

	return data
}

func (h HexResponse) MakeJSONData() []byte {
	if len(h.H) == 0 {
		return []byte("e: " + h.E.Error())
	}

	data, _ := json.Marshal(h)

	return data
}

// Cymk
func (y CymkResponse) makeJSONData() []byte {
	data, e := json.Marshal(y)

	if e != nil {
		return []byte("e: " + e.Error())
	}

	return data
}
