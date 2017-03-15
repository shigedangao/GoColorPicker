package color_test

import (
	"color"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test if the rgb is convert to hexa
func TestRgbToHexa(t *testing.T) {
	// Create a rgb clor
	firstSample := colorHelper.MakeColorFromInput(24, 98, 118)
	// now convert the rgb to hexa
	hexa := firstSample.ConvertRGBtoHexa()
	assert.Equal(t, hexa, "186276", "The test has fail")
}
