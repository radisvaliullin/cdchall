package main

import "fmt"

func main() {
	fmt.Println("main")

	n := 5
	fmt.Printf("sum of range from 0 to %v: %v\n", n, sum(n))
}

// t O(N), mem O(N)
func sum(n int) int {
	if n <= 0 {
		return 0
	}
	return n + sum(n-1)
}
