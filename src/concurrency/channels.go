package myconcurrencymodule

import (
	"fmt"
	"time"
)

func ChannelDemo() {
	// Creating a channel Using var keyword
	var myChannel chan int
	fmt.Println("Value of the channel: ", myChannel)
	fmt.Printf("Type of the channel: %T ", myChannel)

	// Creating a channel using make() function
	mychannel1 := make(chan int)
	fmt.Println("\nValue of the channel1: ", mychannel1)
	fmt.Printf("Type of the channel1: %T ", mychannel1)
}

func sendData(ch chan int) {
	fmt.Println("Sending data...")
	ch <- 42 // This will block until the data is received
	fmt.Println("Data sent!")
}

func BlockingSendReceiveDemo() {
	ch := make(chan int)

	// Start a goroutine to send data to the channel
	go sendData(ch)

	// Simulate some delay
	time.Sleep(2 * time.Second)

	fmt.Println("Receiving data...")
	data := <-ch // This will unblock the send operation in sendData

	fmt.Println("Received data:", data)
	time.Sleep(2 * time.Second)
}

func ChannelLength() {
	mychnl := make(chan string, 4)
	mychnl <- "GFG"
	mychnl <- "gfg"
	mychnl <- "Geeks"
	mychnl <- "GeeksforGeeks"

	// Finding the length of the channel
	// Using len() function
	fmt.Println("Length of the channel is: ", len(mychnl))
}
