package main

import "fmt"

func main() {
	fmt.Println("main")

	n := 5
	fmt.Printf("fib of %v is %v\n", n, fib(n))
}

// O(2^N) (actually 1.6^2, because bottom of call stack sometime has one call).
func fib(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
