package convertcolor

import (
	"encoding/json"
)

// ToJSON from a multiple context
func (jsonize colorJSON) ToJSON() []byte {
	jsonData, e := json.Marshal(jsonize)

	if e != nil {
		return []byte(e.Error())
	}

	return jsonData
}
