/*ColorPicker is a GUI project where you can pick and customize a color that you wnat to use
  This might be useful when you want to deal with some color
  First we're going to test it on a website.. using some test
  Then move to the GUI

  A color picker is first compose of a a suite of color
  Define in multiple format. for the simplification we're going to admit
  that there're only 2 format, the HEX and the RGB color
*/
package main

import (
	"color"
	"fmt"
)

func main() {
	// make some random color
	firstSample := colorHelper.MakeColorFromInput(255, 255, 255)

	// try to update the color and see the effect
	err, c := firstSample.UpdateCurrentColor(24, 98, 118)

	if err != nil {
		fmt.Println("an error happened while converting the color")
	}

	// try to convert an rgb color to an hexa
	hexa := c.ConvertRGBtoHexa()
	fmt.Println(hexa)

	// save a color in the store
	c.SaveColor()
	firstSample.SaveColor()

	hslValue := c.RgbToHsl()
	fmt.Println("HUE is equal to ", hslValue)

	// now that we have a hue we can get the HSL

	// Get the HSL
	_, hue := hslValue.GetHSL()
	fmt.Println("hsl value", hue)
	percentHsl, _ := hslValue.Percent("Luminace")
	fmt.Println("percent hsl luminace", percentHsl)

	// Generate shade
	shade := c.GenerateShade(4)
	// Generate tint
	tint := c.GenerateTint(4)
	// print the shade and the tint
	fmt.Println("shade", shade)
	fmt.Println("tint", tint)

	// create a ycbcr
	ycbcr := c.ConvertYCbCr()
	fmt.Println("ycbcr", ycbcr)

	// convert it back to rgb
	rgb := ycbcr.ConvertToRGB()
	fmt.Println("rgb", rgb)

	// get the cymk color
	cymk := rgb.RgbToCymk()
	fmt.Println("cymk ", cymk)

	// cymk to rgb
	rgbCy := cymk.CymkToRgb()
	fmt.Println("cymk to rgb", rgbCy)

	// RGB to hsv
	_, hsv := c.RgbToHsv()
	fmt.Println("rgb to hsv", hsv)
	// get the percent of the hsv
	percentHSV, _ := hsv.Percent("Value")
	fmt.Println("percent hsv value", percentHSV)

	// HSV to Rgb

	e, rgbHsv := hsv.HsvToRgb()
	fmt.Println("error : ", e)
	fmt.Println("hsv to rgb", rgbHsv)
}
