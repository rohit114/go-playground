package fundamentals

import "fmt"

func PonitersDemo() {

	// storing the hexadecimal
	// values in iables
	x := 0xFF
	y := 0x9C

	// Displaying the values
	fmt.Printf("Type of variable x is %T\n", x)
	fmt.Printf("Value of x in hexadecimal is %X\n", x)
	fmt.Printf("Value of x in decimal is %v\n", x)

	fmt.Printf("Type of variable y is %T\n", y)
	fmt.Printf("Value of y in hexadecimal is %X\n", y)
	fmt.Printf("Value of y in decimal is %v\n", y)

	// foo := "foo"
	// bar := 555
	my_str := "foo"
	my_int := 555

	var foo_address *string = &my_str
	var bar_address *int = &my_int

	fmt.Println("foo_address memory", foo_address)
	fmt.Println("foo_address val", *foo_address)

	fmt.Println("bar_address memory", &bar_address)
	fmt.Println("bar_address val", *bar_address)

}

func PonitersDemoUsingShortHnad() {

	// foo := "foo"
	// bar := 555
	my_str := "foo"
	my_int := 555

	foo_address := &my_str
	bar_address := &my_int

	fmt.Println("foo_address memory", foo_address)
	fmt.Println("foo_address val", *foo_address)

	fmt.Println("bar_address memory", &bar_address)
	fmt.Println("bar_address val", *bar_address)

}

func change2XValue(num *int) {
	*num = (*num) * (2)
}

func PassingPointerToFunctionDemo() {
	num := 123
	fmt.Println("Before:", num)
	change2XValue(&num)
	fmt.Println("After:", num)

}

type Person struct {
	name string
	age  int
}

func PointerToStructDemo() {

	p1 := Person{name: "Rohit Kkummarr", age: 29}

	fmt.Println(p1.name, " , ", p1.age)
	//pointer using &
	p1_address := &p1

	fmt.Println("p1 address", &p1_address)
	p1_address.name = "ROHIT KUMAR"
	p1_address.age = 27
	fmt.Println(p1.name, " , ", p1.age)

}

func DoublePoninterDemo() {

	var num int = 114

	ptr1 := &num //address of num

	var ptr2 **int = &ptr1
	//ptr2 := &ptr1 //using shortehand

	fmt.Println(num, " is stored at", ptr1)
	fmt.Println(ptr1, "is stored at", ptr2)

	fmt.Println("value stored at address", ptr1, " is ", *ptr1)
	fmt.Println("value stored at address", ptr2, " is ", *ptr2)
	fmt.Println("value pointing with double ptr: ", ptr2, " is ", **ptr2)
}

func ComparePonintersDemo() {
	n1 := 111
	n2 := 555

	ptr1 := &n1
	ptr2 := &n2
	ptr3 := &n1

	fmt.Println("ptr1==ptr2: ", ptr1 == ptr2)
	fmt.Println("ptr1==ptr3: ", ptr1 == ptr3)
	fmt.Println("ptr2==ptr3: ", ptr2 == ptr3)

	// fmt.Println("address ptr1==ptr1: ", &ptr1 == &ptr1)
	// fmt.Println("address ptr1==ptr3: ", &ptr1 == &ptr3)
	// fmt.Println("address ptr2==ptr3: ", &ptr2 == &ptr3)

}
