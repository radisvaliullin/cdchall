package main

import "fmt"

func main() {
	fmt.Println("main")

	arr := []int{43, 73, 1, 2, 3, 4}
	fmt.Println("arr: ", arr)
	reverse(arr)
	fmt.Println("reversed arr: ", arr)
}

// O(N)
func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		tailIdx := len(arr) - 1 - i
		arr[i], arr[tailIdx] = arr[tailIdx], arr[i]
	}
}
