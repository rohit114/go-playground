package main

import (
	"fmt"
	"go-playground/src/shipment"
	"time"
)

func main() {

	fmt.Println("Hello World from Golang!")
	go shipment.PlaceNewOrderAndShip(4, "iPhone 16 Pro", 98000)
	go shipment.PlaceNewOrderAndShip(3, "PlayStation", 48000)
	go shipment.PlaceNewOrderAndShip(2, "Samsung", 66000)
	go shipment.PlaceNewOrderAndShip(1, "Nokia", 45000)

	time.Sleep(8000 * time.Millisecond)
	defer func() {
		cache := shipment.GetCache()
		for key, val := range cache {
			fmt.Println(key, " ====> ", val)
		}
	}()
}
