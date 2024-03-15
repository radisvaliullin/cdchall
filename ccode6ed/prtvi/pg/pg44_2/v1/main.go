package main

import "fmt"

func main() {
	fmt.Println("main")

	n := 7
	fmt.Printf("n - %v, f(n) - %v\n", n, f(n))
}

// cpu - O(2^N)
// mem - O(N)
func f(n int) int {
	if n <= 1 {
		return 1
	}
	return f(n-1) + f(n-1)
}
