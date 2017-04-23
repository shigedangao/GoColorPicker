package convertcolor

// colorJSON - This struct is only use when doing BULK request the data will be trimmed when creating the JSON

// Note : We don't create a single reference and cleanup each time after sending back to the main routine. Testing this method increase the response of about += 4ms

type colorJSON struct {
	Rgb   RgbColor   `json:"rgb,omitempty"`
	Hexa  Hex        `json:"hex,omitempty"`
	Hsv   *Hsv       `json:"hsv,omitempty"`
	Ck    *Cymk      `json:"cymk,omitempty"`
	Yb    *YCbCr     `json:"ycbcr,omitempty"`
	Hsl   *HslStruct `json:"hsl,omitempty"`
	Shade []RgbColor `json:"shade,omitempty"`
	Tint  []RgbColor `json:"tint,omitempty"`
	E     error      `json:"error,omitempty"`
}

// ToHex embeded the ToHex method to be called by the server conccurently
func (r RgbColor) ToHex(c chan []byte) {

	hex, e := r.ConvertRGBtoHexa()

	colorize := colorJSON{
		Hexa: hex,
		E:    e,
	}

	c <- colorize.ToJSON()
}

// ToHsv embeded the HSV value into a chan of byte
func (r RgbColor) ToHsv(c chan []byte) {
	hsv, e := r.RgbToHsv()

	colorize := colorJSON{
		Hsv: hsv,
		E:   e,
	}

	c <- colorize.ToJSON()
}

// ToCymk convert a Cymk to JSON
func (r RgbColor) ToCymk(c chan []byte) {
	cymk := r.RgbToCymk()

	colorize := &colorJSON{
		Ck: cymk,
	}

	c <- colorize.ToJSON()
}

// ToYcbCr convert a Ycbcr to JSON
func (r RgbColor) ToYcbCr(c chan []byte) {
	ycbcr := r.ConvertYCbCr()

	colorize := colorJSON{
		Yb: ycbcr,
	}

	c <- colorize.ToJSON()
}

// ToHsl convert an RGB based value to an Hsl value
func (r RgbColor) ToHsl(c chan []byte) {
	hsl := r.RgbToHsl()

	if hsl == nil {
		c <- []byte("hsl is empty")
	}

	colorize := colorJSON{
		Hsl: hsl,
	}

	c <- colorize.ToJSON()
}

// ToShade RGB Convert an RGB value to a shade of RGB Color
func (r RgbColor) ToShade(c chan []byte, factor int) {
	shade, e := r.GenerateShadeTint(factor, "shade")

	colorize := colorJSON{
		Shade: shade,
		E:     e,
	}

	c <- colorize.ToJSON()
}

// ToTint RGB Convert an RGB value to a tint of RGB Color
func (r RgbColor) ToTint(c chan []byte, factor int) {
	tint, e := r.GenerateShadeTint(factor, "tint")

	colorize := colorJSON{
		Tint: tint,
		E:    e,
	}

	c <- colorize.ToJSON()
}

// ToRGBJson convert an rgb value to a json
func (r RgbColor) ToRGBJson(c chan []byte, rgb RgbColor) {
	colorize := colorJSON{
		Rgb: rgb,
	}

	c <- colorize.ToJSON()
}
