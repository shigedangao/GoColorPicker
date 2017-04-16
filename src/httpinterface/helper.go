package serverManager

import (
	"bytes"
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

// SendDataIface interface
type SendDataIface interface {
	MakeJSONData() []byte
}

var (
	rgbData   convertcolor.RgbResponse
	shade     convertcolor.GenerateResponse
	tint      convertcolor.GenerateResponse
	hslData   convertcolor.HslResponse
	hsvData   convertcolor.HsvResponse
	hexData   convertcolor.HexResponse
	container = make([]SendDataIface, 6)
)

// Extract Map Data From URLs
func (r ColorHttpHandler) extractMapDataFromURL() string {
	// Parse the URL based on the "/"
	// Get the url
	url := []byte(r.R.URL.Path)
	// Parse the url
	urlMap := bytes.Split(url, []byte("/"))
	buffer := bytes.NewBuffer(urlMap[2])

	return buffer.String()
}

// // HandleReq Handle an HTTP request with the endpoint /rgb/{params}
// func (r ColorHttpHandler) HandleReq() ([]byte, error) {
// 	// now treat the operation ....
// 	// pare the url
// 	str := r.extractMapDataFromURL()

// 	// call our getInterlRoute and data from a goroutine
// 	jsonData := make(chan []byte)
// 	go r.getInternalRoute(jsonData, str)

// 	data := <-jsonData
// 	// Get the params that we want to

// 	return data, nil
// }

// // extractPOSTData extract the data from a post
func extractPOSTData(r *http.Request) (*reqHandler, error) {

	dataFromReq := &reqHandler{}
	err := json.NewDecoder(r.Body).Decode(dataFromReq)

	if err != nil {
		return nil, err
	}

	return dataFromReq, nil
}

// // getRgb return an rgb value of a type..
// func (r reqHandler) getRgb(typestruct string) convertcolor.RgbColor {

// 	var rgb convertcolor.RgbColor

// 	switch typestruct {
// 	case "hsv":
// 		rgbvalue, e := r.Hsv.ToRGB()

// 		if e != nil {
// 			log.Panicf(e.Error())
// 		} else {
// 			rgb = rgbvalue
// 		}
// 		break
// 	}

// 	return rgb
// }

// process Data process the empty interface array of mixed object and get their json
func processData(colors []SendDataIface) []byte {
	var (
		joiner     [][]byte
		datatosend []byte
	)

	for _, color := range colors {
		// as every object have a makeJSONData we execute it
		if color != nil {
			data := color.MakeJSONData()
			joiner = append(joiner, data)
		}
	}

	datatosend = bytes.Join(joiner, []byte(", "))

	return datatosend
}

// if we do not make a bulk request we need to flush our interface
func isBulk(bulk bool) {
	if !bulk {
		container = make([]SendDataIface, 6)
	}
}
