package colorHTTPInterface

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
)

// global variable which define the upgrader
var upgrade = websocket.Upgrader {
	HandshakeTimeout : 1024,
	ReadBufferSize : 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}


// Handle the websocket connection
func HandleSocket(w http.ResponseWriter, r *http.Request) (*websocket.Conn){
	// Once the socket_con is etablisehd we can use this socket to read and write the data
	socket_con, err := upgrade.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}

	return socket_con
}




