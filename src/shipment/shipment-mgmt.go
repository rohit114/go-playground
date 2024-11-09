package shipment

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

var CacheDB map[string]string = nil

const event_vs_state_key string = "event_vs_state"

func init() {

	// Initialize the global CacheDB variable
	CacheDB = make(map[string]string)
	event_vs_state_config := make(map[string]string)

	event_vs_state_config["PACKED"] = "ORDER_PACKED"
	event_vs_state_config["PICKED_UP"] = "ORDER_PICKED_UP"
	event_vs_state_config["SHIPPED"] = "ORDER_SHIPPED"
	event_vs_state_config["IN_TRANSIT"] = "ORDER_IN_TRANSIT"
	event_vs_state_config["DELIVERED"] = "ORDER_DELIVERED"

	json, err := json.Marshal(event_vs_state_config)
	if err != nil {
		log.Fatal(err)
	}
	CacheDB[event_vs_state_key] = string(json)

}

func GetStateByEvent(key string) string {
	eventVsStateConfig := CacheDB[key]
	if len(eventVsStateConfig) > 0 {
		return eventVsStateConfig
	}
	return ""
}
func UpdateOrderStateOnEvent(event Event, order *Order) {

	switch event {

	case PACKED:

		order.UpdateStateOfOrder(ORDER_PACKED)

	case PICKED_UP:
		order.UpdateStateOfOrder(ORDER_PICKED_UP)
	case SHIPPED:
		order.UpdateStateOfOrder(ORDER_SHIPPED)

	case IN_TRANSIT:
		order.UpdateStateOfOrder(ORDER_IN_TRANSIT)

	case DELIVERED:
		order.UpdateStateOfOrder(ORDER_DELIVERED)
	default:
		fmt.Println("Invalid Event!")

	}

	validateState(event, order.state) //if incorrect state throw error and notify admin

	defer func() {

		if r := recover(); r != nil {
			fmt.Println("Caught panic:", r)
		} else {
			fmt.Println("No panic occurred.")
		}

	}()
}

func validateState(event Event, incomingState State) {

	//EventVsCorrectState from the DB
	event_vs_state := CacheDB[event_vs_state_key]
	event_vs_state_parsed := parseData(event_vs_state)
	//fmt.Print("event_vs_state_parsed:====>", event_vs_state_parsed)
	correctState := event_vs_state_parsed[string(event)]
	if string(incomingState) != correctState {
		panic("Invalid state for the given event ")
	}

}

func parseData(jsonString string) map[string]string {
	// Create a new map to hold the parsed data
	var parsedData map[string]string

	// Unmarshal the JSON string back to the map
	//err :=
	err := json.Unmarshal([]byte(jsonString), &parsedData)
	if err != nil {
		log.Fatal(err)
	}

	// Output the parsed data
	//fmt.Println("Parsed Data:", parsedData)
	return parsedData
}

func CreateOrderOrder(product *Product) Order {

	order, order_obj := CreateNewOrder(1, *product)
	// Convert the object to JSON string

	fmt.Printf("order_obj: %+v\n", order_obj)
	order_key := "order_" + strconv.Itoa(order_obj.Id)
	fmt.Println("order_key: ", order_key)

	orderJSON, err := json.Marshal(order_obj)
	if err != nil {
		log.Fatal(err)
	}

	jsonString := string(orderJSON)
	CacheDB[order_key] = jsonString
	fmt.Println("SAVED: ", jsonString)
	fmt.Println(order_key, ": ", CacheDB[order_key])
	return *order
}

func PlaceNewOrderAndShip() {
	product := CreateNewProduct(1, "iPhone 16 Pro", 98000)
	order := CreateOrderOrder(product)

	defer UpdateOrderStateOnEvent(DELIVERED, &order)
	defer fmt.Println(DELIVERED, " ")
	defer time.Sleep(2000 * time.Millisecond)

	defer UpdateOrderStateOnEvent(IN_TRANSIT, &order)
	defer fmt.Println(IN_TRANSIT)
	defer time.Sleep(2000 * time.Millisecond)

	defer UpdateOrderStateOnEvent(SHIPPED, &order)
	defer fmt.Println(SHIPPED)
	defer time.Sleep(1000 * time.Millisecond)

	defer UpdateOrderStateOnEvent(PICKED_UP, &order)
	defer fmt.Println(PICKED_UP)
	defer time.Sleep(2000 * time.Millisecond)

	defer UpdateOrderStateOnEvent(PACKED, &order)
	defer fmt.Print(PACKED)
	defer time.Sleep(1000 * time.Millisecond)
}
