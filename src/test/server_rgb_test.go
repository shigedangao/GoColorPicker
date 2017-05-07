package color_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"utils"

	"github.com/stretchr/testify/assert"
)

var hsldata = json.RawMessage(`{
	"Rgb" : {
		"Red": 100,
		"Green": 100,
		"Blue": 5
	}, 
		"Factor": 4
	}`)

// Test Handle Rgb To HSL Request
func TestHandleRgbToHSLRequest(t *testing.T) {

	req := httptest.NewRequest("POST", "http://localhost:1698/rgb/hsl", bytes.NewBuffer(hsvdata))
	w := httptest.NewRecorder()

	utils.Handler(req, w, "rgb")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, "an error happened while trying to get the payload")
	}

	assert.NotEmpty(t, data.Hsl, "rgb to hsl is empty")
}

// TestHandleRgbToHexRequest test an rgb to an hex
func TestHandleRgbToHexRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1698/rgb/hex", bytes.NewBuffer(hsvdata))

	w := httptest.NewRecorder()
	utils.Handler(req, w, "rgb")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, "error happened")
	}

	assert.NotEmpty(t, data.Hex, "rgb to hex is empty")
}

// TestHandleRgbToHsvRequest test an rgb to an hsv request
func TestHandleRgbToHsvRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1698/rgb/hsv", bytes.NewBuffer(hsvdata))

	w := httptest.NewRecorder()
	utils.Handler(req, w, "rgb")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, "error happened")
	}

	assert.NotEmpty(t, data.Hsv, "rgb to hsv is empty")
}

// TestHandleRgbToCymkRequest test rgb to cymk
func TestHandleRgbToCymkRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1698/rgb/cymk", bytes.NewBuffer(hsvdata))

	w := httptest.NewRecorder()
	utils.Handler(req, w, "rgb")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, "error happened")
	}

	assert.NotEmpty(t, data.Cymk, "rgb to cymk is empty")
}

// TestHandleRgbToYcbcrRequest test an rgb request to get an ycbcr value
func TestHandleRgbToYcbcrRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1798/rgb/ycbcr", bytes.NewBuffer(hsvdata))

	w := httptest.NewRecorder()
	utils.Handler(req, w, "rgb")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, "error happened")
	}

	assert.NotEmpty(t, data.YCbCr, "rgb to ycbcr is empty")
}

// TestHandleRgbToShadeRequest test an rgb request to get a shade value
func TestHandleRgbToShadeRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1798/rgb/shade", bytes.NewBuffer(hsvdata))

	w := httptest.NewRecorder()
	utils.Handler(req, w, "rgb")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, "error happened")
	}

	assert.NotEmpty(t, data.Shade, "rgb to shade is empty")
}

// TestHandleRgbToShadeRequest test an rgb request to get a tint value
func TestHandleRgbToTintRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1798/rgb/tint", bytes.NewBuffer(hsvdata))

	w := httptest.NewRecorder()
	utils.Handler(req, w, "rgb")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, "error happened")
	}

	assert.NotEmpty(t, data.Tint, "rgb to tint is empty")
}
