package fundamentals

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func StringPractice1() {
	fmt.Println("hello from string test")

	fmt.Println("Hello from StringPractice2")
	var name string = "Rohit"
	for index, value := range name {
		fmt.Printf("%d : %c\n", index, value)
	}

	myList := []string{"apple", "mango", "orange", "banana"}
	for index, value := range myList {
		fmt.Printf("%d : %s\n", index, value)
	}

	var myList2 [10]string

	myList2[0] = "Alice"
	myList2[2] = "Bob"
	fmt.Print(myList2)
	fmt.Println()

}

func CapitalizeStringDemo() {
	name := "  rohit kumar  "
	name2 := "rohIT KumaR"
	email := "rohit@gmail.com"
	name = strings.TrimSpace(name)
	name = string(unicode.ToUpper(rune(name[0]))) + name[1:]
	fmt.Println("name:", name)

	email = strings.ToUpper(email)
	name2 = strings.ToLower(name2)
	fmt.Println("email:", email)
	fmt.Println("name2:", name2)

}

func FindFrequencyOfCharsInTheString(str string) {

	frequency := make(map[int32]int)

	for _, char := range str {

		frequency[char]++
	}

	for key, val := range frequency {
		fmt.Printf("%c -> %d \n", key, val)
	}

}

func IsAnagram(s string, t string) bool {

	sBytes := []rune(s)
	tBytes := []rune(t)

	sort.Slice(sBytes, func(i, j int) bool {
		return sBytes[i] > sBytes[j]
	})
	sort.Slice(tBytes, func(i, j int) bool {
		return tBytes[i] > tBytes[j]
	})

	return string(tBytes) == string(sBytes)

}

func IsAnagram2(s string, t string) bool {

	mymap := make(map[rune]int)

	for _, val := range s {
		mymap[val]++
	}

	for key, val := range mymap {
		fmt.Printf("%c -> %d \n", key, val)
	}

	fmt.Println()
	for _, val := range t {
		mymap[val]--
	}

	for key, val := range mymap {
		fmt.Printf("%c -> %d \n", key, val)
	}

	for _, val := range mymap {
		if val != 0 {
			return false
		}
	}
	return true

}
