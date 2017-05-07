package color_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"utils"

	"github.com/stretchr/testify/assert"
)

var hsvdata = json.RawMessage(`{
	"Hsv": {
		"H": 100,
		"S": 20.0,
		"V": 10.5
	}, 
	"Factor": 4
}`)

// TestHandleHsvToRgbRequest
func TestHandleHsvToRgbRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1698/hsv/rgb", bytes.NewBuffer(hsvdata))

	w := httptest.NewRecorder()

	utils.Handler(req, w, "hsv")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, e.Error())
	}

	assert.NotEmpty(t, data.Rgb, "hsv to rgb is empty")
}

// TestHandleHsvToHexRequest
func TestHandleHsvToHexRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1698/hsv/hex", bytes.NewBuffer(hsvdata))

	w := httptest.NewRecorder()

	utils.Handler(req, w, "hsv")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, e.Error())
	}

	assert.NotEmpty(t, data.Hex, "hsv to hex is empty")
}

// TestHandleHsvToHslRequest
func TestHandleHsvToHslRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1698/hsv/hsl", bytes.NewBuffer(hsvdata))

	w := httptest.NewRecorder()
	utils.Handler(req, w, "hsv")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, e.Error())
	}

	assert.NotEmpty(t, data.Hsl, "hsv to hsl is empty")
}

// TestHandleHsvToCymkRequest
func TestHandleHsvToCymkRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1698/hsv/cymk", bytes.NewBuffer(hsvdata))

	w := httptest.NewRecorder()
	utils.Handler(req, w, "hsv")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, e.Error())
	}

	assert.NotEmpty(t, data.Cymk, "hsv to cymk is empty")
}

// TestHandleHsvToYcbcrRequest
func TestHandleHsvToYcbcrRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1698/hsv/ycbcr", bytes.NewBuffer(hsvdata))

	w := httptest.NewRecorder()
	utils.Handler(req, w, "hsv")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, e.Error())
	}

	assert.NotEmpty(t, data.YCbCr, "hsv to ycbcr is empty")
}

// TestHandleHsvToShadeRequest
func TestHandleHsvToShadeRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1698/hsv/shade", bytes.NewBuffer(hsvdata))

	w := httptest.NewRecorder()
	utils.Handler(req, w, "hsv")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, e.Error())
	}

	assert.NotEmpty(t, data.Shade, "hsv to shade is empty")
}

// TestHandleHsvToTintRequest
func TestHandleHsvToTintRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1698/hsv/tint", bytes.NewBuffer(hsvdata))

	w := httptest.NewRecorder()
	utils.Handler(req, w, "hsv")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, e.Error())
	}

	assert.NotEmpty(t, data.Tint, "hsv to tint is empty")
}
