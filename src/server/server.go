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
	b         bytes.Buffer
	colorHTTP ColorHTTPHandler
	mux       *http.ServeMux
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

// TypeHandler handle the request for every datas
func TypeHandler(typedata string) {
	// Create our own custom handler
	mux.HandleFunc(string([]byte("/"+typedata+"/")), func(w http.ResponseWriter, r *http.Request) {
		colorHTTP := colorHTTP.prepareAPI(w, r, typedata)
		data, e := colorHTTP.HandleType(typedata)

		handleResponse(w, data, e)
	})
}

// @TODO refactor the type handler in order to not be anonymous (a part from the shade and tint)

// MakeServer - Create our MUX server
func MakeServer() {
	mux = http.NewServeMux()
	route := readRoute()

	for _, data := range route {
		routeData := data.(map[string]interface{})

		resource := routeData["datatype"].(string)
		TypeHandler(resource)
	}

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
