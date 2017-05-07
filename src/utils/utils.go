package utils

import (
	"color"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"server"
)

// CheckValue is an Object which handle every type which is used by the test (we use a pointer so we know when a type is empty)
type CheckValue struct {
	Hsl   *convertcolor.HslStruct
	Rgb   *convertcolor.RgbColor
	Hex   *convertcolor.Hex
	Hsv   *convertcolor.Hsv
	Cymk  *convertcolor.Cymk
	YCbCr *convertcolor.YCbCr
	Shade []convertcolor.RgbColor
	Tint  []convertcolor.RgbColor
}

// Handler handle the fake request
func Handler(r *http.Request, w http.ResponseWriter, typedata string) {
	var handler = serverManager.ColorHTTPHandler{}

	colorHTTP := handler.PrepareAPI(w, r, typedata)
	data, e := colorHTTP.HandleType()

	if e != nil {
		w.Write([]byte("error"))
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(data)
}

// ReadResponse read the response from the fake http server
func ReadResponse(recorder *httptest.ResponseRecorder) (*CheckValue, error) {
	result := recorder.Result()
	rawdata, e := ioutil.ReadAll(result.Body)
	var object *CheckValue

	if e != nil {
		return nil, e
	}

	err := json.Unmarshal(rawdata, &object)

	if err != nil {
		return nil, err
	}

	return object, nil
}
