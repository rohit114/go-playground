package fundamentals

import (
	"fmt"
)

func riskyFunction() {
	// Simulate an error
	panic("Something went wrong!")
}

func TestExceptionHandle() {
	// Using defer for 'finally'
	defer func() {
		if r := recover(); r != nil {
			// Catch the panic and handle it
			fmt.Println("Caught panic:", r)
		} else {
			// If no panic occurred, this part will execute
			fmt.Println("No panic occurred.")
		}
	}()

	// Simulating a try-catch-finally behavior
	fmt.Println("Starting risky operation...")

	// Try block equivalent
	riskyFunction()

	// This won't be reached if panic is triggered
	fmt.Println("Operation completed successfully.")
}

func riskyFun2() {
	panic("Bad Request")
}
func TestExceptionHandle2() {

	defer func() {

		if r := recover(); r != nil {
			fmt.Println("Caught panic:", r)
		} else {
			fmt.Println("No panic occurred.")
			defer func() {

				fmt.Println("Releasing lock...")

			}()
		}

	}()

	fmt.Println("Start")
	riskyFun2()
	fmt.Println("Success")
}
