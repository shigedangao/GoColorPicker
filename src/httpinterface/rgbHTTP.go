package colorHTTPInterface

import (
	"bytes"
	"fmt"
	"net/http"
)

// RgbHandler is an http handler for the HTTP
type RgbHandler struct {
	R *http.Request
	W http.ResponseWriter
}

// Extract Map Data From URLs
func (r RgbHandler) extractMapDataFromURL() string {
	// Parse the URL based on the "/"
	// Get the url
	url := []byte(r.R.URL.Path)
	// Parse the url
	urlMap := bytes.Split(url, []byte("/"))
	buffer := bytes.NewBuffer(urlMap[2])

	return buffer.String()
}

// HandleReq Handle an HTTP request with the endpoint /rgb/{params}
func (r RgbHandler) HandleReq() ([]byte, error) {
	// now treat the operation ....
	// pare the url
	str := r.extractMapDataFromURL()
	fmt.Println(str)

	// call our getInterlRoute and data from a goroutine
	jsonData := make(chan []byte)
	go r.getInternalRoute(jsonData)

	data := <-jsonData
	// Get the params that we want to

	return data, nil
}

func (r RgbHandler) getInternalRoute(data chan []byte) {
	data <- []byte("hey cha")
}
