package color_test

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestHandleRgbRequest(t *testing.T) {
	data := []byte("{}")
	buf := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", "/rgb/hex", buf)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(req)
}
