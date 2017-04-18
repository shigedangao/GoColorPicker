package serverManager

import (
	"color"
)

type colorList struct {
	Rgb    convertcolor.RgbColor
	Hexa   convertcolor.Hex
	Hsv    *convertcolor.Hsv
	Hsl    *convertcolor.HslStruct
	Factor int
}

// SendDataIface interface
type JSONize interface {
	ToJSON() []byte
}
