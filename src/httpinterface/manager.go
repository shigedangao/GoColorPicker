package serverManager

import "net/http"

// ColorHttpHandler is an http handler for the HTTP
type ColorHttpHandler struct {
	R *http.Request
	W http.ResponseWriter
	T string
}

// HandleReq handle the request for every type
func (c ColorHttpHandler) HandleReq() {
	// Based on the T we shall execute their own parser
	switch c.T {
	case "rgb":
		// Call specific thing for handling rgb
		break
	case "hsv":
		break
	case "cymk":
		break
	case "ycbcr":
		break
	case "hsl":
		break
	}
}
