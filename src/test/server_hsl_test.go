package color_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"utils"

	"github.com/stretchr/testify/assert"
)

var hsl = json.RawMessage(`{
	"Hsl": {
		"Angle": 340,
		"Saturation": 10.51,
		"Luminace": 0.55
	}, 
		"Factor": 4
	}`)

// TestHandleHsvToRgbRequest
func TestHandleHslToRgbRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1698/hsl/rgb", bytes.NewBuffer(hsl))

	w := httptest.NewRecorder()
	utils.Handler(req, w, "hsl")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, "an error happened while trying to get the payload")
	}

	assert.NotEmpty(t, data.Rgb, "not empty")
}

// TestHandleHsvToHsvRequest
func TestHandleHslToHsvRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1698/hsl/hsv", bytes.NewBuffer(hsl))

	w := httptest.NewRecorder()
	utils.Handler(req, w, "hsl")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, "an error happened while trying to get the payload")
	}

	assert.NotEmpty(t, data.Hsv, "not empty")
}

// TestHandleHsvToHsvRequest
func TestHandleHslToHexRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1698/hsl/hex", bytes.NewBuffer(hsl))

	w := httptest.NewRecorder()
	utils.Handler(req, w, "hsl")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, "an error happened while trying to get the payload")
	}

	assert.NotEmpty(t, data.Hex, "not empty")
}

// TestHandleHsvToHsvRequest
func TestHandleHslToCymkRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1698/hsl/cymk", bytes.NewBuffer(hsl))

	w := httptest.NewRecorder()
	utils.Handler(req, w, "hsl")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, "an error happened while trying to get the payload")
	}

	assert.NotEmpty(t, data.Cymk, "not empty")
}

// TestHandleHsvToYcbcrRequest
func TestHandleHslToYcbcrRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1698/hsl/ycbcr", bytes.NewBuffer(hsl))

	w := httptest.NewRecorder()
	utils.Handler(req, w, "hsl")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, "an error happened while trying to get the payload")
	}

	assert.NotEmpty(t, data.YCbCr, "not empty")
}

// TestHandleHsvToShadeRequest
func TestHandleHslToShadeRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1698/hsl/shade", bytes.NewBuffer(hsl))

	w := httptest.NewRecorder()
	utils.Handler(req, w, "hsl")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, "an error happened while trying to get the payload")
	}

	assert.NotEmpty(t, data.Shade, "not empty")
}

// TestHandleHsvToTintRequest
func TestHandleHslToTintRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:1698/hsl/tint", bytes.NewBuffer(hsl))

	w := httptest.NewRecorder()
	utils.Handler(req, w, "hsl")
	data, e := utils.ReadResponse(w)

	if e != nil {
		assert.FailNow(t, "an error happened while trying to get the payload")
	}

	assert.NotEmpty(t, data.Tint, "not empty")
}
