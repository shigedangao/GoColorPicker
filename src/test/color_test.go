package color_test

import (
	"color"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test RGB Creation
func TestRgbCreation(t *testing.T) {
	sample := colorHelper.RgbColor{
		Red:   255,
		Green: 255,
		Blue:  255,
	}

	assert.ObjectsAreEqualValues(sample, colorHelper.MakeColorFromInput(255, 255, 255))
}

// TestBadRgbStruct creating a bad RgbColor struct
func TestBadRgbStructLiteral(t *testing.T) {
	assert.NotPanics(t, func() {
		sample := colorHelper.RgbColor{
			Red:   0,
			Green: 255,
			Blue:  254,
		}

		fmt.Println("no panics ", sample)
	}, "Making a bad struct should not make them call Panics")
}

// TestBadRgbStructInterface creating a bad RgbColor using the interface
func TestBadRgbStructInterface(t *testing.T) {
	assert.NotPanics(t, func() {
		sample := colorHelper.MakeColorFromInput(255, 0, 254)
		fmt.Println("no panics ", sample)
	}, "should not panics when creating a struct using an interface..")
}

// Creating an RGB using a constructor and test the conversion toward Hexa
func TestRgbToHexa(t *testing.T) {
	firstSample := colorHelper.MakeColorFromInput(24, 98, 118)
	hexa := firstSample.ConvertRGBtoHexa()
	assert.Equal(t, hexa, "186276", "The test has fail")
}

// Hexa to an RGB value
func TestHexaToRGB(t *testing.T) {
	rgb, e := colorHelper.ToRGB("186276")

	if e != nil {
		assert.Fail(t, "error converting the rgb to an hexa")
	}
	assert.Equal(t, uint8(24), rgb.Red, "red is wrong")
	assert.Equal(t, uint8(98), rgb.Green, "green is wrong")
	assert.Equal(t, uint8(118), rgb.Blue, "blue is wrong")
}

// TestRgbToCymk test the conversion of an RGB value to Cymk
func TestRgbToCymk(t *testing.T) {
	rgb := colorHelper.RgbColor{
		Red:   25,
		Green: 50,
		Blue:  20,
	}

	cymk := rgb.RgbToCymk()

	assert.EqualValues(t, "0.5", strconv.FormatFloat(cymk.C, 'f', 1, 64), "c is not equal")
	assert.EqualValues(t, "0.6", strconv.FormatFloat(cymk.Y, 'f', 1, 64), "y is not equal")
	assert.EqualValues(t, "0", strconv.FormatFloat(cymk.M, 'f', 0, 64), "m is not equal")
	assert.EqualValues(t, "0.804", strconv.FormatFloat(cymk.K, 'f', 3, 64), "k is not equal")
}

// TestCymkToRgb testing cymk to rgb
func TestCymkToRgb(t *testing.T) {
	cymk := colorHelper.Cymk{
		C: 0.5,
		Y: 0.6,
		M: 0,
		K: 0.804,
	}

	rgb := cymk.ToRGB()
	assert.EqualValues(t, uint8(24), rgb.Red, "red value is wrongly calculated")
	assert.EqualValues(t, uint8(49), rgb.Green, "green value is wrongly calculated")
	assert.EqualValues(t, uint8(19), rgb.Blue, "blue value is wrongly calculated")
}

// TestRgbYcBcR testing rgb to ycbcr
func TestRgbYcBcR(t *testing.T) {
	rgb := colorHelper.RgbColor{
		Red:   25,
		Green: 50,
		Blue:  20,
	}

	ycbcr := rgb.ConvertYCbCr()
	// assert
	assert.Equal(t, "39", strconv.FormatFloat(ycbcr.Y, 'f', 0, 64), "y value is wrong")
	assert.Equal(t, "117", strconv.FormatFloat(ycbcr.Cb, 'f', 0, 64), "Cb value is wrong")
	assert.Equal(t, "118", strconv.FormatFloat(ycbcr.Cr, 'f', 0, 32), "Cr value is wrong")
}

// Test YcbCr to Rgb
func TestYcbCrToRgb(t *testing.T) {
	ycbcr := colorHelper.YCbCr{
		Y:  39.105,
		Cb: 117.2175,
		Cr: 117.9396,
	}

	rgb := ycbcr.ToRGB()
	assert.EqualValues(t, uint8(25), rgb.Red, "red value is wrongly calculated")
	assert.EqualValues(t, uint8(50), rgb.Green, "green value is wrongly calculated")
	assert.EqualValues(t, uint8(19), rgb.Blue, "blue value is wrongly calculated")
}

// Create a normal Rgb value and test to convert it into HSL
func TestRgbToHsl(t *testing.T) {
	sample := colorHelper.RgbColor{
		Red:   3,
		Green: 34,
		Blue:  76,
	}

	// convert the sample to an Hsl
	hsl := sample.RgbToHsl()
	// get the int representation of the struct
	hslPercent, err := hsl.FormatFloatToInt()

	if err != nil {
		assert.FailNow(t, "an error happened while converting the HSL into a MAP of INT percent")
	}
	assert.EqualValues(t, 214, hsl.Angle, "The Hue value is wrong")
	assert.EqualValues(t, 92, hslPercent["saturation"], "The saturation value is wrong")
	assert.EqualValues(t, 15, hslPercent["luminace"], "The luminace value is wrong")
}

// TestRgbToHslUniform Create a "uniform" constant RGB value and convert it into an HSL
func TestRgbToHslUniform(t *testing.T) {
	sample := colorHelper.RgbColor{
		Red:   250,
		Green: 250,
		Blue:  250,
	}

	// convert the sample to an Hsl
	hsl := sample.RgbToHsl()
	// get the int representation of the struct
	hslPercent, err := hsl.FormatFloatToInt()

	if err != nil {
		assert.FailNow(t, "an error happened while converting the HSL into a MAP of INT percent")
	}
	assert.EqualValues(t, 0, hsl.Angle, "The Hue value is wrong")
	assert.EqualValues(t, 0, hslPercent["saturation"], "The saturation value is wrong")
	assert.EqualValues(t, 98, hslPercent["luminace"], "The luminace value is wrong")
}

// TestHslToRGB test a conversion from an HSL struct to an RGB struct
func TestHslToRGB(t *testing.T) {

	// Precision might play a bit on the value
	// It's recommend to use the RAW value than the Round int value to convert the data

	hslSample := &colorHelper.HslStruct{
		Angle:      203,
		Saturation: 0.9797979797979799,
		Luminace:   0.19411764705882353,
	}

	rgb, e := hslSample.ToRGB()

	if e != nil {
		assert.FailNow(t, "an error happened while converting the hsl to an RGB")
	}

	assert.EqualValues(t, uint8(1), rgb.Red, "Red is wrongly calculated")
	assert.EqualValues(t, uint8(60), rgb.Green, "Green is wrongly calculated")
	assert.EqualValues(t, uint8(97), rgb.Blue, "Blue is wrongly calculated")
}

// testRgbToHsv test a conversion from an RGB value to an HSV
func TestRgbToHsv(t *testing.T) {
	rgb := colorHelper.RgbColor{
		Red:   1,
		Green: 60,
		Blue:  98,
	}

	// convert the rgb to the hsv
	hsv, e := rgb.RgbToHsv()

	if e != nil {
		assert.FailNow(t, "converting an rgb to an hsv has fail")
	}

	// i should use a parsefloat or convert it into a percent...
	assert.EqualValues(t, 203, hsv.H, "hsv's hue value is wrong")
	assert.EqualValues(t, 0.9897959183673469, hsv.S, "hsv's S value is wrong")
	assert.EqualValues(t, 0.3843137254901961, hsv.V, "hsv's V value is wrong")
}

// TestHsvToRgb - convert an hsv to an rgb
func TestHsvToRgb(t *testing.T) {
	sample := colorHelper.Hsv{
		H: 203,
		S: 0.98,
		V: 0.38,
	}

	// now converting an hsv to an rgb
	rgb, e := sample.ToRGB()

	if e != nil {
		assert.EqualValues(t, 1, rgb.Red, "red valus is wrong")
		assert.EqualValues(t, 60, rgb.Green, "green value is wrong")
		assert.EqualValues(t, 98, rgb.Blue, "blue value is wrong")
	}
}

// TestGeneerateTint - Testing the creation of a tint with normal values
func TestGenerateTint(t *testing.T) {
	// generating multipe tint value
	// make as simple rgb value
	rgb := colorHelper.RgbColor{
		Red:   1,
		Green: 60,
		Blue:  98,
	}

	assert.NotPanics(t, func() {
		// generate the tint
		tint, _ := rgb.GenerateShadeTint(3, "tint")
		assert.NotNil(t, tint, "tint should not be nil")
	}, "Creating tint should not panic")
}

// TestGeneerateShade - Testing the creation of a tint with normal values
func TestGenerateShade(t *testing.T) {
	// generating multipe tint value
	// make as simple rgb value
	rgb := colorHelper.RgbColor{
		Red:   1,
		Green: 60,
		Blue:  98,
	}

	assert.NotPanics(t, func() {
		// generate the tint
		shade, _ := rgb.GenerateShadeTint(3, "shade")
		assert.NotNil(t, shade, "shade should not be nil")
	}, "Creating shade should not panic")
}

// TestGeneerateBadTint - Trying to generate a bad tint using 0 as a factor
func TestGenerateBadTint(t *testing.T) {
	// generating multipe tint value
	// make as simple rgb value
	rgb := colorHelper.RgbColor{
		Red:   1,
		Green: 60,
		Blue:  98,
	}

	assert.NotPanics(t, func() {
		// generate the tint
		tint, _ := rgb.GenerateShadeTint(0, "tint")
		assert.Empty(t, tint, "it is not empty")
	}, "Creating bad tint should not panic")
}

// TestGeneerateBadShade - Testing the creation a bad shade using 0 as a factor
func TestGenerateBadShade(t *testing.T) {
	// generating multipe tint value
	// make as simple rgb value
	rgb := colorHelper.RgbColor{
		Red:   1,
		Green: 60,
		Blue:  98,
	}

	assert.NotPanics(t, func() {
		// generate the tint
		shade, _ := rgb.GenerateShadeTint(0, "shade")
		assert.Empty(t, shade, "it is not empty")
	}, "Creating bad shade should not panic")
}

// Testing generate wrong type of shade
func TestGenerateWrongType(t *testing.T) {
	rgb := colorHelper.RgbColor{
		Red:   1,
		Green: 60,
		Blue:  98,
	}

	assert.NotPanics(t, func() {
		shade, e := rgb.GenerateShadeTint(4, "")
		assert.Nil(t, shade, "shade is not nil")
		assert.Equal(t, "An error happened while generating the colors", e.Error(), "not equal")
	}, "should not panics when passing a wrong genType")
}
