package serverManager

import (
	"bytes"
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

var (
	b       bytes.Buffer
	rgbHTTP ColorHTTPHandler
	hsvHTTP ColorHTTPHandler
	cmkHTTP ColorHTTPHandler
	ybrHTTP ColorHTTPHandler
	hslHTTP ColorHTTPHandler
)

// extract Post Data
func (h ColorHTTPHandler) extractPOSTData() (*colorList, error) {

	if h.R == nil {
		return nil, errors.New("request is empty")
	}

	var dataFromReq *colorList

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

// MakeServer - Create our MUX server
func MakeServer() {
	// Define our mux server
	mux := http.NewServeMux()

	// Handle the rgb request
	mux.HandleFunc("/rgb/", func(w http.ResponseWriter, r *http.Request) {
		rgbHTTP = rgbHTTP.prepareAPI(w, r, "rgb")
		data, e := rgbHTTP.HandleRGBRequest()
		// Write the data
		handleResponse(w, data, e)
	})

	mux.HandleFunc("/hsv/", func(w http.ResponseWriter, r *http.Request) {
		hsvHTTP.prepareAPI(w, r, "hsv")
	})

	mux.HandleFunc("/cymk/", func(w http.ResponseWriter, r *http.Request) {
		cmkHTTP.prepareAPI(w, r, "cymk")
	})

	mux.HandleFunc("/ycbcr/", func(w http.ResponseWriter, r *http.Request) {
		ybrHTTP.prepareAPI(w, r, "ycbcr")
	})

	mux.HandleFunc("/hsl/", func(w http.ResponseWriter, r *http.Request) {
		hslHTTP.prepareAPI(w, r, "hsl")
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
