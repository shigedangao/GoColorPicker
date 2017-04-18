package convertcolor

// colorJSON - This struct is only use when doing BULK request the data will be trimmed when creating the JSON
type colorJSON struct {
	Rgb  RgbColor   `json:"rgb"`
	Hexa Hex        `json:"hex"`
	Hsv  *Hsv       `json:"hsv"`
	Ck   Cymk       `json:"cymk"`
	Yb   YCbCr      `json:"ycbcr"`
	Hsl  *HslStruct `json:"hsl"`
	E    error      `json:"error"`
}

// ToHex embeded the ToHex method to be called by the server conccurently
func (r RgbColor) ToHex(c chan []byte) {

	hex, e := r.ConvertRGBtoHexa()

	colorize := &colorJSON{
		Hexa: hex,
		E:    e,
	}

	c <- colorize.ToJSON()
}

// ToHsv embeded the HSV value into a chan of byte
func (r RgbColor) ToHsv(c chan []byte) {
	hsv, e := r.RgbToHsv()

	if e != nil {
		c <- []byte(e.Error())
	}

	c <- hsv.ToJSON()
}

// ToCymk convert a Cymk to JSON
func (r RgbColor) ToCymk(c chan []byte) {
	cymk := r.RgbToCymk()

	// Check if the cymk object is empty
	if cymk == (Cymk{}) {
		c <- []byte("cymk is empty")
	}

	c <- cymk.ToJSON()
}

// ToYcbCr convert a Ycbcr to JSON
func (r RgbColor) ToYcbCr(c chan []byte) {
	ycbcr := r.ConvertYCbCr()

	if ycbcr == (YCbCr{}) {
		c <- []byte("ycbcr is empty")
	}

	c <- ycbcr.ToJSON()
}

// ToHsl convert an RGB based value to an Hsl value
func (r RgbColor) ToHsl(c chan []byte) {
	hsl := r.RgbToHsl()

	if hsl == nil {
		c <- []byte("hsl is empty")
	}

	c <- hsl.ToJSON()
}
