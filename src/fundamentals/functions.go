package fundamentals

func MultiplyCallByValue(a, b int) int {
	a = a * 2 // modifying a inside the function
	return a * b
}

func MultiplyCallRefrence(a, b *int) int {
	*a = *a * 2 // modifying a's value at its memory address
	return (*a) * (*b)
}

func SwapInt(a, b *int) {
	var temp int = *a
	*a = *b
	*b = temp
}

func SwapChar(a, b *string) {
	var temp string = *a
	*a = *b
	*b = temp
}

func VeriadicSumFun(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}
