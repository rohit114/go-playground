package fundamentals

func BinarySearch(arr []int, l, h, target int) int {

	for l <= h {
		mid := l + (h-l)/2

		if arr[mid] == target {
			return mid
		}

		if target < arr[mid] {
			//go left
			h = mid - 1
		} else {
			//go right
			l = mid + 1
		}
	}
	return -1

}
