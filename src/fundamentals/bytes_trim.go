package fundamentals

import (
	"bytes"
	"fmt"
	"strings"
)

func TrimByteDemo() {

	byteSlice := []byte("  !!Hello, World!  ")
	trimmed_space := bytes.TrimSpace(byteSlice)
	trimmed_symbol := bytes.TrimLeft(trimmed_space, "!")
	fmt.Println(string(trimmed_space))  // "Hello, World!"
	fmt.Println(string(trimmed_symbol)) // "Hello, World!"

	// Creating and trimming
	// the slice of bytes
	// Using Trim function
	res1 := bytes.Trim([]byte("****Welcome to GeeksforGeeks****"), "*")
	res2 := bytes.Trim([]byte("!!!!Learning how to trim a slice of bytes@@@@"), "!@")
	res3 := bytes.Trim([]byte("^^Geek&&"), "$")

	// Display the results
	fmt.Printf("Final Slice:\n")
	fmt.Printf("\nSlice 1: %s", res1)
	fmt.Printf("\nSlice 2: %s", res2)
	fmt.Printf("\nSlice 3: %s", res3)

}

func TrimSringeDemo() {

	var stringSlice string = "   Rohit Kumar!!!    "
	trimmed_space := strings.TrimSpace(stringSlice)
	trimmed_symbol := strings.TrimLeft(trimmed_space, "!")
	fmt.Println(string(trimmed_space))  // "Hello, World!"
	fmt.Println(string(trimmed_symbol)) // "Hello, World!"

	// Creating and trimming
	// the slice of bytes
	// Using Trim function
	res1 := strings.Trim("****Welcome to GeeksforGeeks****", "*")
	res2 := strings.Trim("!!!!Learning how to trim a slice of bytes@@@@", "!@$")
	res3 := strings.Trim("^^Geek&&", "$")

	// Display the results
	fmt.Printf("Final Slice:\n")
	fmt.Printf("\nSlice 1: %s", res1)
	fmt.Printf("\nSlice 2: %s", res2)
	fmt.Printf("\nSlice 3: %s", res3)
	fmt.Println()

}

func SplitSringeDemo() {
	var my_str string = "Audi,Honda,Tata,BYD"
	// list := strings.Split(my_str, ",")
	// for _, val := range list {
	// 	fmt.Println(val)
	// }
	// fmt.Println(list[2])

	list2 := strings.SplitN(my_str, ",", -1)

	for _, val := range list2 {
		fmt.Println(val)
	}
	//fmt.Println(list2[2])

	my_slice := []int{}
	for i := 0; i <= 10; i++ {
		my_slice = append(my_slice, i*i)
	}

	for _, val := range my_slice {
		fmt.Println(val)
	}

}
