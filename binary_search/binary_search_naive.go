package main

func binary_search(arr []int, find_item, low, high int) int {
	// Assume the array is already sorted
	mid := (high + low) / 2

	if low <= high {
		if arr[mid] == find_item {
			return mid
		} else if find_item < arr[mid] {
			// Lesser than the mid, so value exists to the left of the middle point
			return binary_search(arr, find_item, mid-1, high)
		} else if find_item > arr[mid] {
			// Item must be larger than the mid, so need to find it from the right side of the arr
			return binary_search(arr, find_item, mid+1, high)
		}
	}

	return -1 // Item is not in array
}

func main() {
	// Array/Slice must be sorted
	// Sorting the array should be simple though - with the slices package
	// slices.Sort()
	arr := []int{1, 2, 3, 4, 5}

	high := len(arr) - 1
	low := 0

	// You can log this here to be sure it works as expected
	binary_search(arr, 4, low, high)
}
