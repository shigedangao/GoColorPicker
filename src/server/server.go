package serverManager

import (
	"bytes"
	"color"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// ColorHTTPHandler is an http handler for the HTTP
type ColorHTTPHandler struct {
	R *http.Request
	W http.ResponseWriter
	T string
}

type colorList struct {
	Rgb    convertcolor.RgbColor
	Hexa   convertcolor.Hex
	Hsv    *convertcolor.Hsv
	Hsl    *convertcolor.HslStruct
	Cymk   convertcolor.Cymk
	ycbcr  convertcolor.YCbCr
	Factor int
}

// JSONize interface
type JSONize interface {
	ToJSON() []byte
}

var (
	b       bytes.Buffer
	rgbHTTP ColorHTTPHandler
	hsvHTTP ColorHTTPHandler
	cmkHTTP ColorHTTPHandler
	ybrHTTP ColorHTTPHandler
	hslHTTP ColorHTTPHandler
)

// extract Post Data
func (h ColorHTTPHandler) extractPOSTData() (colorList, error) {

	if h.R == nil {
		return colorList{}, errors.New("request is empty")
	}

	var dataFromReq colorList

	if err := json.NewDecoder(h.R.Body).Decode(&dataFromReq); err != nil {
		fmt.Println(err)
	}

	return dataFromReq, nil
}

// extract Map Data From URL
func (h ColorHTTPHandler) extractMapDataFromURL() string {
	// Parse the URL based on the "/"
	// Get the url
	url := []byte(h.R.URL.Path)
	// Parse the url
	urlMap := bytes.Split(url, []byte("/"))
	buffer := bytes.NewBuffer(urlMap[2])

	return buffer.String()
}

func (h ColorHTTPHandler) prepareAPI(w http.ResponseWriter, r *http.Request, t string) ColorHTTPHandler {
	h.R = r
	h.W = w
	h.T = t

	return h
}

// @TODO refactor the type handler in order to not be anonymous (a part from the shade and tint)

// MakeServer - Create our MUX server
func MakeServer() {
	// Define our mux server
	mux := http.NewServeMux()

	// Handle the rgb request
	mux.HandleFunc("/rgb/", func(w http.ResponseWriter, r *http.Request) {
		rgbHTTP = rgbHTTP.prepareAPI(w, r, "rgb")
		data, e := rgbHTTP.HandleType("rgb")
		// Write the data
		handleResponse(w, data, e)
	})

	mux.HandleFunc("/hsv/", func(w http.ResponseWriter, r *http.Request) {
		hsvHTTP = hsvHTTP.prepareAPI(w, r, "hsv")
		data, e := hsvHTTP.HandleType("hsv")

		handleResponse(w, data, e)
	})

	mux.HandleFunc("/cymk/", func(w http.ResponseWriter, r *http.Request) {
		cmkHTTP = cmkHTTP.prepareAPI(w, r, "cymk")
		data, e := cmkHTTP.HandleType("cymk")

		handleResponse(w, data, e)
	})

	mux.HandleFunc("/ycbcr/", func(w http.ResponseWriter, r *http.Request) {
		ybrHTTP = ybrHTTP.prepareAPI(w, r, "ycbcr")
		data, e := ybrHTTP.HandleType("ycbcr")

		handleResponse(w, data, e)
	})

	mux.HandleFunc("/hsl/", func(w http.ResponseWriter, r *http.Request) {
		hslHTTP = hslHTTP.prepareAPI(w, r, "hsl")
		data, e := hslHTTP.HandleType("hsl")

		handleResponse(w, data, e)
	})

	// Listen our server
	fmt.Println("run server")
	log.Fatal(http.ListenAndServe(":1698", mux))
}

// sendData - sendData to the client
func handleResponse(w http.ResponseWriter, d []byte, err error) {
	if err != nil {
		log.Panic(err.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(d)
}
