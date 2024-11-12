package shipment

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

const EVENT_STATE_KEY string = "event_vs_state"

var mu sync.Mutex // Declare the mutex
func init() {

	// Initialize the global CacheDB variable
	InitCache()
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

	SetCache(EVENT_STATE_KEY, string(json))

}

func CreateOrderOrder(product *Product) Order {

	order, order_obj := CreateNewOrder(product.id*10, *product)
	fmt.Printf("order_obj: %+v\n", order_obj)
	savToDB(order_obj)

	return *order
}

func UpdateOrderStateOnEvent(event Event, order *Order) {

	switch event {

	case PACKED:
		order.UpdateStateOfOrder(ORDER_PACKED)
		order_obj := GetOrderObject(order)
		savToDB(order_obj)

	case PICKED_UP:
		order.UpdateStateOfOrder(ORDER_PICKED_UP)
		order_obj := GetOrderObject(order)
		savToDB(order_obj)

	case SHIPPED:
		order.UpdateStateOfOrder(ORDER_SHIPPED)
		order_obj := GetOrderObject(order)
		savToDB(order_obj)

	case IN_TRANSIT:
		order.UpdateStateOfOrder(ORDER_IN_TRANSIT)
		order_obj := GetOrderObject(order)
		savToDB(order_obj)

	case DELIVERED:
		order.UpdateStateOfOrder(ORDER_DELIVERED)
		order_obj := GetOrderObject(order)
		savToDB(order_obj)
	default:
		fmt.Println("Invalid Event!")

	}

	validateState(event, order.state) //if incorrect state throw error and notify admin
	fmt.Println("Size of CacheDB =======>", len(GetCache()))
	defer func() {

		if r := recover(); r != nil {
			fmt.Println("Caught panic:", r)
		} else {
			fmt.Println("No panic occurred.")
		}

	}()
}

func PlaceNewOrderAndShip(id int, name string, price int) {
	product := CreateNewProduct(id, name, price)
	order := CreateOrderOrder(product)

	defer UpdateOrderStateOnEvent(DELIVERED, &order)
	defer fmt.Println(DELIVERED, " ")
	defer time.Sleep(1000 * time.Millisecond)

	defer UpdateOrderStateOnEvent(IN_TRANSIT, &order)
	defer fmt.Println(IN_TRANSIT)
	defer time.Sleep(1000 * time.Millisecond)

	defer UpdateOrderStateOnEvent(SHIPPED, &order)
	defer fmt.Println(SHIPPED)
	defer time.Sleep(1000 * time.Millisecond)

	defer UpdateOrderStateOnEvent(PICKED_UP, &order)
	defer fmt.Println(PICKED_UP)
	defer time.Sleep(1000 * time.Millisecond)

	defer UpdateOrderStateOnEvent(PACKED, &order)
	defer fmt.Print(PACKED)
	defer time.Sleep(1000 * time.Millisecond)
}

func validateState(event Event, incomingState State) {

	//EventVsCorrectState from the DB
	event_vs_state := CacheDB[EVENT_STATE_KEY]
	event_vs_state_parsed := parseData(event_vs_state)
	correctState := event_vs_state_parsed[string(event)]

	if string(incomingState) != correctState {
		panic("Invalid state for the given event ")
	}

}

func parseData(jsonString string) map[string]string {
	var parsedData map[string]string
	// Unmarshal the JSON string back to the map
	err := json.Unmarshal([]byte(jsonString), &parsedData)
	if err != nil {
		log.Fatal(err)
	}
	return parsedData
}

func savToDB(order_obj OrderObject) {
	// order_key := "order_" + strconv.Itoa(order_obj.Id)
	// fmt.Println("order_key: ", order_key)

	// orderJSON, err := json.Marshal(order_obj)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// jsonString := string(orderJSON)
	// CacheDB[order_key] = jsonString
	// fmt.Println("SAVED: ", jsonString)
	// fmt.Println(order_key, ": ", CacheDB[order_key])
	// Lock the mutex to ensure exclusive access to CacheDB
	mu.Lock()
	defer mu.Unlock()

	order_key := "order_" + strconv.Itoa(order_obj.Id)
	fmt.Println("order_key: ", order_key)

	// Convert the order object to JSON
	orderJSON, err := json.Marshal(order_obj)
	if err != nil {
		log.Fatal(err)
	}

	// Convert JSON byte slice to string
	jsonString := string(orderJSON)

	// Store the JSON string in CacheDB
	CacheDB[order_key] = jsonString

	// Print the saved entry and the CacheDB value
	fmt.Println("SAVED: ", jsonString)
	fmt.Println(order_key, ": ", CacheDB[order_key])
}
