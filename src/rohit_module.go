package rohitmodule

import "fmt"

func Hello() {
	fmt.Println("Hello from rohitmodule")

}

func FindAllDuplicates(nums []int) []int {
	result := []int{}

	my_map := make(map[int]int)
	fmt.Print("Hello1")
	for _, value := range nums {
		my_map[value]++
	}
	fmt.Print("Hello2")
	for key, _ := range my_map {
		if my_map[key] > 1 {
			fmt.Print("Hello3")
			result = append(result, key)
		}
		my_map[key]++
		//fmt.Println(key, " :", val)

	}
	fmt.Print("Hello4", len(result))

	return result

}

func IsDuplicateExists(nums []int) bool {

	my_set := make(map[int]bool)
	for _, val := range nums {
		if my_set[val] {
			return true
		}
		my_set[val] = true
	}
	fmt.Println("my_set size", len(my_set))

	return false
}
