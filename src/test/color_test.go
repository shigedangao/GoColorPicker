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

func TestYcbCrToRgb(t *testing.T) {
	ycbcr := colorHelper.YCbCr{
		Y:  39.105,
		Cb: 117.2175,
		Cr: 117.9396,
	}

	rgb := ycbcr.ToRGB()
	fmt.Println(rgb)
	assert.EqualValues(t, uint8(25), rgb.Red, "red value is wrongly calculated")
	assert.EqualValues(t, uint8(50), rgb.Green, "green value is wrongly calculated")
	assert.EqualValues(t, uint8(19), rgb.Blue, "blue value is wrongly calculated")
}
