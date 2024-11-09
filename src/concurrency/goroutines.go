package myconcurrencymodule

import (
	"fmt"
	"time"
)

func greet(message string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1000 * time.Millisecond)
		println(message)
	}
}

func GoRoutinesDemo1() {
	go greet("Active Go Routine")

	greet("MAIN")

}

// For goroutine 1
func Aname() {

	arr1 := [4]string{"Rohit", "Suman", "Aman", "Ria"}

	for t1 := 0; t1 <= 3; t1++ {

		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%s\n", arr1[t1])
	}
}

// For goroutine 2
func Aid() {

	arr2 := [4]int{300, 301, 302, 303}

	for t2 := 0; t2 <= 3; t2++ {

		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d\n", arr2[t2])
	}
}

func MultiGoroutinesDemo() {
	fmt.Println("!...Main Go-routine Start...!")

	// calling Goroutine 1
	go Aname()

	// calling Goroutine 2
	go Aid()

	time.Sleep(4000 * time.Millisecond)

	defer fmt.Println("\n!...Main Go-routine End...!")

}
