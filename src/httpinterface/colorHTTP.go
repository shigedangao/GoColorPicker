package colorHTTPInterface

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
	isBulk(false)
	reqData, _ := extractPOSTData(r.R)

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
		data <- []byte("route " + urlMap + " is not supported")
	}

	// as each type has it's own makeJSONData function we can call it threw the generation
	data <- processData(container)
}
