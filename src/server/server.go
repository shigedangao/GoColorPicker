package serverManager

import (
	"bytes"
	"fmt"
	"httpinterface"
	"log"
	"net/http"
)

var (
	b       bytes.Buffer
	rgbHTTP serverManager.ColorHttpHandler
	hsvHTTP serverManager.ColorHttpHandler
	cmkHTTP serverManager.ColorHttpHandler
	ybrHTTP serverManager.ColorHttpHandler
	hslHTTP serverManager.ColorHttpHandler
)

func (chandler serverManager.ColorHttpHandler) prepareAPI(w http.ResponseWriter, r *http.Request, t string) {

	chandler.R = r
	chandler.W = w
	chandler.T = t
}

// MakeServer - Create our MUX server
func MakeServer() {
	// Define our mux server
	mux := http.NewServeMux()

	// Handle the rgb request
	mux.HandleFunc("/rgb/", func(w http.ResponseWriter, r *http.Request) {
		// Call our manager here...
		rgbHTTP.prepareAPI(w, r, "rgb")

		data, e := rgbHTTP.HandleReq()
		handleResponse(w, data, e)
	})

	mux.HandleFunc("/hsv/", func(w http.ResponseWriter, r *http.Request) {
		hsvHTTP.prepareAPI(w, r, "hsv")

		data, e := hsvHTTP.HandleHsvReq()
		handleResponse(w, data, e)
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
