package colorHTTPInterface

import (
	"color"
	"encoding/json"
	"net/http"
)

type handler struct {
	Rgb  colorHelper.RgbColor
	Hexa string
	Hsv  *colorHelper.Hsv
	Hsl  *colorHelper.HslStruct
}

// extractPOSTData extract the data from a post
func extractPOSTData(params []string, r *http.Request) (*handler, error){

	dataFromReq := &handler{}
	err := json.NewDecoder(r.Body).Decode(dataFromReq)

	if err != nil {
		return nil, err
	}

	return dataFromReq, nil
}

// makeJSONData transform our struct into a json
func makeJSONData(h handler) ([]byte, error){
	data, err := json.Marshal(h)

	if err != nil {
		return nil, err
	}

	return data, nil
}


