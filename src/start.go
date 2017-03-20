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

func colorSample() {
	// make some random color
	firstSample := colorHelper.MakeColorFromInput(255, 255, 255)

	// try to update the color and see the effect
	c, err := firstSample.UpdateCurrentColor(190, 120, 200)

	if err != nil {
		fmt.Println("an error happened while converting the color")
	}

	// try to convert an rgb color to an hexa
	hexa := c.ConvertRGBtoHexa()
	fmt.Println(hexa)

	// hexa to rgb
	rgbFromHexa, _ := colorHelper.ToRGB(hexa)
	fmt.Println("rgb from hexa", rgbFromHexa)

	// save a color in the store
	c.SaveColor()
	firstSample.SaveColor()

	hslValue := c.RgbToHsl()
	fmt.Println("HSL is equal to", hslValue)

	// hsl to rgb
	rgbHSL, _ := hslValue.ToRGB()
	fmt.Println("HSL to rgb equal", rgbHSL)

	// now that we have a hue we can get the HSL

	percentHsl, _ := hslValue.Percent("Luminace")
	fmt.Println("percent hsl luminace", percentHsl)

	// now let just create a reference of our interface and test if it works...
	hextest := &colorHelper.HslStruct{
		Angle:      292,
		Saturation: 0.421,
		Luminace:   0.627,
	}
	var i colorHelper.Color
	i = hextest
	rgbInterfaceTest, _ := i.ToRGB()

	fmt.Println(rgbInterfaceTest)

	// // Generate shade
	shade, e := c.GenerateShadeTint(4, "shade")

	if e != nil {
		fmt.Println("error")
	}

	// Generate tint
	tint, _ := c.GenerateShadeTint(4, "tint")
	// // print the shade and the tint
	fmt.Println("shade", shade)
	fmt.Println("tint", tint)

	// // create a ycbcr
	// ycbcr := c.ConvertYCbCr()
	// fmt.Println("ycbcr", ycbcr)

	// // convert it back to rgb
	// rgb := ycbcr.ConvertToRGB()
	// fmt.Println("rgb", rgb)

	// // get the cymk color
	// cymk := rgb.RgbToCymk()
	// fmt.Println("cymk ", cymk)

	// // cymk to rgb
	// rgbCy := cymk.CymkToRgb()
	// fmt.Println("cymk to rgb", rgbCy)

	// // RGB to hsv
	hsv, _ := c.RgbToHsv()
	fmt.Println("rgb to hsv", hsv)
	// // get the percent of the hsv
	// percentHSV, _ := hsv.Percent("Value")
	// fmt.Println("percent hsv value", percentHSV)

	// // HSV to Rgb

	// e, rgbHsv := hsv.HsvToRgb()
	// fmt.Println("error : ", e)
	// fmt.Println("hsv to rgb", rgbHsv)
}

func main() {
	colorSample()
	//serverManager.MakeServer()
}
