package main

import (
	"fmt"
)

func main() {
	fmt.Println("main")

	arr := []int{9, 9, 9}
	n := 9
	fmt.Printf("n - %v, arr - %v, idx - %v\n", n, arr, binarySearch(n, arr))

	arr = []int{9, 19, 29}
	n = 9
	fmt.Printf("n - %v, arr - %v, idx - %v\n", n, arr, binarySearch(n, arr))

	arr = []int{9, 19, 29}
	n = 29
	fmt.Printf("n - %v, arr - %v, idx - %v\n", n, arr, binarySearch(n, arr))

	arr = []int{}
	n = 9
	fmt.Printf("n - %v, arr - %v, idx - %v\n", n, arr, binarySearch(n, arr))

	arr = []int{1, 2, 3, 4}
	n = 9
	fmt.Printf("n - %v, arr - %v, idx - %v\n", n, arr, binarySearch(n, arr))
}

// cpu - O(logN);
// return index of search value n in array (sorted)
// if not found return len(arr)
// if found idx in range [0, len(arr))
func binarySearch(n int, arr []int) int {

	// range index in array
	i, j := 0, len(arr)
	for i < j {
		// get middle value by devide 2
		// overflow posible
		mid := int(uint(i+j) >> 1)
		if !(arr[mid] >= n) {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}
