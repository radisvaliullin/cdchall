package main

import "fmt"

func main() {
	fmt.Println("main")

	n := 4
	fmt.Printf("factorial of %v is %v\n", n, factorial(n))
}

// O(N)
func factorial(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
