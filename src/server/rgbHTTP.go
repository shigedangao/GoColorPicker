package serverManager

// This file handle every request for converting a given RGB type to an other Type

// HandleRGBRequest handle everyr request for the RGB type
func (h ColorHTTPHandler) HandleRGBRequest() ([]byte, error) {
	var dataToWrite []byte
	_color, e := h.extractPOSTData()

	if e != nil {
		return nil, e
	}

	subroute := h.extractMapDataFromURL()
	rgbChan := make(chan []byte)

	switch subroute {
	case "hex":
		go _color.Rgb.ToHex(rgbChan)
		break
	case "hsv":
		go _color.Rgb.ToHsv(rgbChan)
		break
	case "cymk":
		go _color.Rgb.ToCymk(rgbChan)
		break
	case "ycbcr":
		go _color.Rgb.ToYcbCr(rgbChan)
		break
	case "hsl":
		go _color.Rgb.ToHsl(rgbChan)
		break
	case "shade":
		go _color.Rgb.ToShade(rgbChan, _color.Factor)
		break
	case "tint":
		go _color.Rgb.ToTint(rgbChan, _color.Factor)
		break
	default:
		return nil, nil
	}

	dataToWrite = <-rgbChan

	return dataToWrite, nil
}
