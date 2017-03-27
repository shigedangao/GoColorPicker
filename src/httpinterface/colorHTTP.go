package colorHTTPInterface

import (
	"bytes"
	"net/http"
	"color"
)

// RgbHandler is an http handler for the HTTP
type ColorHttpHandler struct {
	R *http.Request
	W http.ResponseWriter
}

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

// HandleReq Handle an HTTP request with the endpoint /rgb/{params}
func (r ColorHttpHandler) HandleReq() ([]byte, error) {
	// now treat the operation ....
	// pare the url
	str := r.extractMapDataFromURL()

	// call our getInterlRoute and data from a goroutine
	jsonData := make(chan []byte)
	go r.getInternalRoute(jsonData, str)

	data := <-jsonData
	// Get the params that we want to

	return data, nil
}

// GetInternalRoute execute a set of function based on the route wanted
// Allowed route
// ---> /rgb/{params}
// ---> /rgb/makergb
// ---> /rgb/hexa
// ---> /rgb/hsv
// ---> /rgb/hsl
// ---> /rgb/shade
// ---> /rgb/tint
func (r ColorHttpHandler) getInternalRoute(data chan []byte, urlMap string) {

	var (jsonData []byte
	     h handler)

	reqData, _ := extractPOSTData(nil, r.R)
	switch urlMap {
	case "makergb":
		// Convert an HEX to an RGB
		color, _ := colorHelper.ToRGB(reqData.Hexa)
		h.Rgb = color
		break
	case "hexa":
		color := reqData.Rgb.ConvertRGBtoHexa()
		h.Hexa = color
		break
	case "hsv":
		color, _ := reqData.Rgb.RgbToHsv()
		h.Hsv = color
		break
	case "hsl":
		color := reqData.Rgb.RgbToHsl()
		h.Hsl = color
		break
	case "shade":

		break
	case "tint":
		break
	default:
		data <- []byte("route " + urlMap + " not supported")
	}

	// now we can make the json
	jsonData, e := makeJSONData(h)

	if e != nil {
		data <- []byte(e.Error())
	}

	data <- []byte(jsonData)
}
