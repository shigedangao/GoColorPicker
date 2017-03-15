package serverManager

import (
	"log"
	"net/http"
)

// Handling the conversion
// I use closure to make it much more cleaner to user
func conversionHandler(typeName string) http.Handler {

	// we can make some call here...
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello everyone !"))
	}

	return http.HandlerFunc(fn)
}

func MakeServer() {
	mux := http.NewServeMux()

	// call our closure here by passing the first params
	// handle rgb to hue route
	rgbToHue := conversionHandler("rgbToHue")
	mux.Handle("/convert/rgbToHue", rgbToHue)

	// handle rgb to hexa

	mux.HandleFunc("/convert/rgbToHex", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Rgb to hexa !"))
	})
	// handle the template

	log.Fatal(http.ListenAndServe(":1698", mux))
}
