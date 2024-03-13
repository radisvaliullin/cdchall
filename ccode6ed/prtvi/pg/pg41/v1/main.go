package main

import "fmt"

func main() {
	fmt.Println("main")

	n := 5
	fmt.Printf("pair sum of range from 0 to %v: %v\n", n, pairSumOfRange(n))
}

// t O(N) (n calls of pairSum), mem O(1) (each pairSum use same stack)
func pairSumOfRange(n int) int {
	s := 0
	for i := 0; i < n; i++ {
		s += pairSum(i, i+1)
	}
	return s
}

func pairSum(a, b int) int {
	return a + b
}
