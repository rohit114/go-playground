package fundamentals

import (
	"fmt"
	"strings"
)

func FindFrequencyOfChars() {
	myString := "hello geek rohit kumar"

	// Initialize a counter array of size 26 (for letters a-z)
	var counter [26]int
	frequency := make(map[int32]int)

	// Convert string to lowercase to handle both uppercase and lowercase letters
	myString = strings.ToLower(myString)

	// Iterate through each character in the string
	for _, char := range myString {
		if char >= 'a' && char <= 'z' { // Check if the character is a letter
			index := char - 'a' // Calculate the index for each letter
			counter[index]++    // Increment the counter at that index
			frequency[char]++
		}
	}

	// Print the letter counts
	fmt.Println("Letter frequencies:")
	for i, count := range counter {
		if count > 0 { // Only print letters that appear in the string
			fmt.Printf("%c: %d\n", 'a'+i, count)
		}
	}
	fmt.Println()

	for key, val := range frequency {
		fmt.Printf("%c -> %d", key, val)
	}

	fmt.Println(frequency[101])

}
