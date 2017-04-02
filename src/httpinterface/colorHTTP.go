package colorHTTPInterface

import (
	"bytes"
	"color"
	"net/http"
)

// RgbHandler is an http handler for the HTTP
type ColorHttpHandler struct {
	R *http.Request
	W http.ResponseWriter
}

type SendDataIface interface {
	MakeJSONData() []byte
}

// we might need to overload the error ?

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

	var (
		rgbData   convertcolor.RgbResponse
		shade     convertcolor.GenerateResponse
		tint      convertcolor.GenerateResponse
		hslData   convertcolor.HslResponse
		hsvData   convertcolor.HsvResponse
		hexData   convertcolor.HexResponse
		container = make([]SendDataIface, 6)
	)

	reqData, _ := extractPOSTData(nil, r.R)

	switch urlMap {
	case "makergb":
		// Convert an HEX to an RGB
		rgbData.R, rgbData.E = reqData.Hexa.ToRGB()
		container[0] = rgbData
		break
	case "hexa":
		hexData.H, hexData.E = reqData.Rgb.ConvertRGBtoHexa()
		container[1] = hexData
		break
	case "hsv":
		hsvData.V, hsvData.E = reqData.Rgb.RgbToHsv()
		container[2] = hsvData
		break
	case "hsl":
		hslData.H = reqData.Rgb.RgbToHsl()
		container[3] = hslData
		break
	case "shade":
		shade.R, shade.E = reqData.Rgb.GenerateShadeTint(reqData.Factor, "shade")
		container[4] = shade
		break
	case "tint":
		tint.R, tint.E = reqData.Rgb.GenerateShadeTint(reqData.Factor, "tint")
		container[5] = tint
		break
	default:
		data <- []byte("route " + urlMap + " not supported")
	}

	// as each type has it's own makeJSONData function we can call it threw the generation
	data <- processData(container)
}

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
