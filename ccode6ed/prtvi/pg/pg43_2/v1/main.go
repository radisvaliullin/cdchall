package main

import (
	"fmt"
)

func main() {
	fmt.Println("main")

	narr := []int{}

	narr = insert(narr, 0)
	narr = insert(narr, 1)
	narr = insert(narr, 2)
	narr = insert(narr, 3)
	fmt.Println("insert result:", narr)
}

// go slice append item using capacity
// if capacity end, runtime reallocate new space for capacity
// insertion of item out of capacity has O(2N) time
// but amortized time is O(1)
func insert(narr []int, n int) []int {
	return append(narr, n)
}
