package serverManager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type route struct {
	method   string
	datatype string
}

func readRoute() map[string]interface{} {
	routesUnmarshal, e := ioutil.ReadFile("config/route.json")

	if e != nil {
		fmt.Println(e)
		log.Fatal("wasn't able to read the route file")
	}

	// Register a route interface using the map
	var maproute map[string]interface{}
	json.Unmarshal(routesUnmarshal, &maproute)
	return maproute
}
