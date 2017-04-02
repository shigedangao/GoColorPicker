package colorHTTPInterface

import (
	"color"
	"encoding/json"
	"net/http"
)

type reqHandler struct {
	Rgb    convertcolor.RgbColor
	Hexa   convertcolor.Hex
	Hsv    *convertcolor.Hsv
	Hsl    *convertcolor.HslStruct
	Factor int
}

// extractPOSTData extract the data from a post
func extractPOSTData(params []string, r *http.Request) (*reqHandler, error) {

	dataFromReq := &reqHandler{}
	err := json.NewDecoder(r.Body).Decode(dataFromReq)

	if err != nil {
		return nil, err
	}

	return dataFromReq, nil
}

// Convert the data into a JSON depending of the type...
