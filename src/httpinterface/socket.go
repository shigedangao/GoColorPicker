package colorHTTPInterface

import (
	"fmt"
	"log"
	"net/http"

	"bytes"

	"github.com/gorilla/websocket"
)

// global variable which define the upgrader
var upgrade = websocket.Upgrader{
	HandshakeTimeout: 1024,
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handle the websocket connection
func HandleSocket(w http.ResponseWriter, r *http.Request) *websocket.Conn {

	// Once the socket_con is etablisehd we can use this socket to read and write the data
	socket_con, err := upgrade.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}

	for {
		messageType, p, err := socket_con.ReadMessage()
		var buffer = bytes.NewBuffer(p)

		if err != nil {
			return nil
		}

		fmt.Println("message type ", messageType)
		fmt.Println("message data ", buffer.String())
	}

	return socket_con
}
