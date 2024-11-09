package fundamentals

import "fmt"

func MapsPractice() {
	fmt.Println("Hello from maps")

	var mymap = make(map[int]string)
	mymap[1] = "Audi"
	mymap[2] = "BMW"
	mymap[3] = "TATA"

	for key, val := range mymap {
		print(key, ": ", val, "\n")
	}

	println("_____________________________")
	// for items := range mymap {
	// 	print(items, "\n")
	// }

	mymap2 := map[string]int{
		"default": 0,
	}
	mymap2["foo"] = 111
	mymap2["bar"] = 222
	mymap2["car"] = 333

	for key, val := range mymap2 {
		print(key, ": ", val, "\n")
	}

}

func IsDuplicateExists(arr []int, n int) bool {
	set := make(map[int]bool)
	for _, val := range arr {
		if set[val] {
			return true
		}
		set[val] = true

	}
	return false

}
