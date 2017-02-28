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

	//fmt.Println(c)
	//fmt.Println(firstSample)

	// try to convert an rgb color to an hexa
	hexa := c.ConvertRGBtoHexa()
	fmt.Println(hexa)

	// save a color in the store
	c.SaveColor()
	firstSample.SaveColor()

	hueValue := c.GenerateOtherFormat("hue")
	fmt.Println("HUE is equal to ", hueValue.Saturation)

	// now that we have a hue we can get the HSL

	// Get the HSL
	_, hsl := hueValue.GetHSL()
	fmt.Println(hsl)
}
