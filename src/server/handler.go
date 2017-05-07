package serverManager

import (
	"color"
	"errors"
)

// This file handle every request for converting a given HSV type to an other type

// HandleType handle the request for hsv
func (h ColorHTTPHandler) HandleType() ([]byte, error) {
	var (
		dataToWrite []byte
	)
	_color, e := h.extractPOSTData()

	if e != nil {
		return nil, e
	}

	subroute := h.extractMapDataFromURL()
	hsvChan := make(chan []byte)

	// In the underlying process we first convert the data into an RGB value in order to manipulate easily the data

	rgb, _ := _color.getRGBColorFromType(h.T)

	switch subroute {
	case "hex":
		go rgb.ToHex(hsvChan)
		break
	case "hsl":
		go rgb.ToHsl(hsvChan)
	case "rgb":
		go rgb.ToRGBJson(hsvChan, rgb)
		break
	case "cymk":
		go rgb.ToCymk(hsvChan)
		break
	case "ycbcr":
		go rgb.ToYcbCr(hsvChan)
		break
	case "tint":
		go rgb.ToTint(hsvChan, _color.Factor)
		break
	case "shade":
		go rgb.ToShade(hsvChan, _color.Factor)
		break
	case "hsv":
		go rgb.ToHsv(hsvChan)
		break
	default:
		hsvChan <- []byte("type is not supported")
	}

	dataToWrite = <-hsvChan
	return dataToWrite, nil
}

func (c colorList) getRGBColorFromType(typeColor string) (convertcolor.RgbColor, error) {

	var (
		rgb convertcolor.RgbColor
		e   error
	)

	switch typeColor {
	case "hsv":
		rgb, e = c.Hsv.ToRGB()
		break
	case "hsl":
		rgb, e = c.Hsl.ToRGB()
		break
	case "hex":
		rgb, e = c.Hexa.ToRGB()
		break
	case "cymk":
		rgb = c.Cymk.ToRGB()
		break
	case "ycbcr":
		rgb = c.ycbcr.ToRGB()
		break
	case "rgb":
		rgb = c.Rgb
	default:
		return rgb, errors.New("type is not supported")
	}

	return rgb, e
}
