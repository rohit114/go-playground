package main

import (
	"fmt"
	"go-playground/src/shipment"
)

var CacheDB map[string]string = nil

func init() {

	//CacheDB := make(map[string]map[string]string)
	//set event_vs_state config
	// Initialize the global CacheDB variable
	CacheDB = make(map[string]string)
	event_vs_state_config := make(map[string]string)
	//config_key := "event_vs_state"

	event_vs_state_config["PACKED"] = "ORDER_PACKED"
	event_vs_state_config["PICKED_UP"] = "ORDER_PICKED_UP"
	event_vs_state_config["SHIPPED"] = "ORDER_SHIPPED"
	event_vs_state_config["IN_TRANSIT"] = "ORDER_IN_TRANSIT"
	event_vs_state_config["DELIVERED"] = "ORDER_DELIVERED"
	//CacheDB[config_key] = event_vs_state_config

	// for key, value := range event_vs_state_config {
	// 	fmt.Printf("init()")
	// 	fmt.Printf(key, " : ", value)
	// }

}

// Function to retrieve state by event name

func main() {

	fmt.Println("Hello World from Golang!")
	shipment.PlaceNewOrderAndShip()
}
