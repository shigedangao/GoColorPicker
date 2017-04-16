package serverManager

import (
	"encoding/json"
)

// HandleHsvReq handle request for the hsv route
func (r ColorHttpHandler) HandleHsvReq() ([]byte, error) {
	route := r.extractMapDataFromURL()

	jsonData := make(chan []byte)
	go r.HandleHsvRoute(jsonData, route)
	data := <-jsonData

	return data, nil
}

// HandleHsvRoute handle route which use hsv as their basis object to contact the package
func (r ColorHttpHandler) HandleHsvRoute(data chan []byte, urlMap string) {
	isBulk(false)
	reqdata, _ := extractPOSTData(r.R)

	if reqdata == nil {
		data <- []byte(json.RawMessage("{err : no data sent}"))
		return
	}

	// if it's not a bulk request we need to flush our interface slice

	rgb := reqdata.getRgb("hsv")

	switch urlMap {
	case "rgb":
		rgbData.R, rgbData.E = reqdata.Hsv.ToRGB()
		container[0] = rgbData
		break
	case "hexa":
		// First we need to convert the value to an rgb
		hexData.H, hexData.E = rgb.ConvertRGBtoHexa()
		container[1] = hexData
		break
	case "hsl":
		hslData.H = rgb.RgbToHsl()
		container[2] = hslData
		break
	case "shade":
		shade.R, shade.E = rgb.GenerateShadeTint(reqdata.Factor, "shade")
		container[3] = shade
		break
	case "tint":
		tint.R, tint.E = rgb.GenerateShadeTint(reqdata.Factor, "tint")
		container[4] = tint
		break
	default:
		data <- []byte("route " + urlMap + " is not supported")
	}

	data <- processData(container)
}
