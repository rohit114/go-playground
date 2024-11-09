package fundamentals

import (
	"fmt"
	"sort"
)

func ListPractice() {

	var numbers [100]int //Array Delacration 1

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			numbers[i] = i * 2
		} else {
			numbers[i] = i
		}

	}

	// for i := 0; i < 20; i++ {
	// 	print(i, " = ", numbers[i], "\n")
	// }

	nums2 := [10]int{66, 77, 88, 99} //Array Delacration 2 with initilization
	for _, val := range nums2 {
		print(val, " ")
	}
	println()

	nums3 := [...]int{9, 7, 6, 4, 5, 3} //Array Delacration 3 with initilization
	nums3Copy := &nums3
	nums3Copy[0] = 999
	fmt.Println("arr2 size", len(nums3))
	fmt.Println("nums3:")
	for _, val := range nums3 {
		print(val, " ")
	}
	println()

	fmt.Println("nums3Copy:")
	for _, val := range nums3Copy {
		print(val, " ")
	}
	println()

}

func PrintMatrix(matrix [][]int) {
	//matrix := [][]int{{1, 2, 4}, {4, 5, 6}, {7, 8, 9}}

	for _, rows := range matrix {
		for _, val := range rows {

			fmt.Print(val, " ")
		}
		fmt.Println()
	}

}

func ReverseArray(arr *[]int, size int) {
	i := 0
	j := size - 1
	for i <= j {
		temp := (*arr)[i]
		(*arr)[i] = (*arr)[j]
		(*arr)[j] = temp
		i++
		j--

	}
}

func DeepCopyShallowCopyDemo() {
	original := []int{1, 2, 3, 4, 5}
	shallowCopy := original
	shallowCopy[0] = 999

	fmt.Println("shallowCopy:", shallowCopy)
	fmt.Println("original:", original)

	deepCopy := make([]int, len(original))
	n := copy(deepCopy, original)

	deepCopy[0] = 88888
	fmt.Println("deepCopy:", deepCopy)
	fmt.Println("original:", original)
	fmt.Println("n:", n)

}

func SortingDemo() {
	my_arrr := []int{2200, 44, 66, 77, 220, 99, 110}
	my_slice := my_arrr[:3]
	my_slice[0] = 9999
	print("---->", my_arrr[0])
	for _, val := range my_slice {
		print(val, " ")
	}
	println()
	//fundamentals.DeepCopyShallowCopyDemo()

	sort.Ints(my_arrr)
	println("my_arrr sorted ascending:")
	for _, val := range my_arrr {
		print(val, " ")
	}
	println("\n")

	sort.Sort(sort.Reverse(sort.IntSlice(my_arrr))) //decending
	println("my_arrr sorted decending:")
	for _, val := range my_arrr {
		print(val, " ")
	}
	println("\n")
	str_list := []string{"apple", "coconut", "orange", "banana"}
	sort.Strings(str_list)
	println("str_list sorted ASC")
	for _, val := range str_list {
		print(val, " ")
	}
	println("\n")
	sort.Sort(sort.Reverse(sort.StringSlice(str_list)))
	println("str_list sorted DESC")
	for _, val := range str_list {
		print(val, " ")
	}
	println("\n")
}

func SortingDemo2() {
	my_arrr := []int{2200, 44, 66, 77, 220, 99, 110}

	for _, val := range my_arrr {
		print(val, " ")
	}
	println()

	sort.Ints(my_arrr)
	println("my_arrr sorted ascending:")
	for _, val := range my_arrr {
		print(val, " ")
	}
	println("\n")

	//sort.Sort(sort.Reverse(sort.IntSlice(my_arrr))) //decending
	sort.Slice(my_arrr, func(i, j int) bool {
		return my_arrr[i] > my_arrr[j]
	})
	println("my_arrr sorted decending:")
	for _, val := range my_arrr {
		print(val, " ")
	}
	println("\n")

	str_list := []string{"apple", "coconut", "orange", "banana"}
	sort.Strings(str_list)
	println("str_list sorted ASC")
	for _, val := range str_list {
		print(val, " ")
	}
	println("\n")
	//sort.Sort(sort.Reverse(sort.StringSlice(str_list)))

	sort.Slice(str_list, func(i, j int) bool {
		return str_list[i] > str_list[j]
	})
	println("str_list sorted DESC")
	for _, val := range str_list {
		print(val, " ")
	}
	println("\n")
	defer DeepCopyShallowCopyDemo()
}

func topKFrequent(nums []int, k int) []int {
	// Input: nums = [1,1,1,2,2,3,4,4,4,4,4,4,4], k = 2
	// Output: [1,2]
	mymap := make(map[int]int)
	for _, val := range nums {
		mymap[val]++
	}

	var sortedList []int
	for num := range mymap {
		sortedList = append(sortedList, num)
	}

	sort.Slice(sortedList, func(i, j int) bool {
		return mymap[sortedList[i]] > mymap[sortedList[j]]
	})

	return sortedList[:k]

}
