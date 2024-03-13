package main

import "fmt"

func main() {
	fmt.Println("main")

	printAllFib(7)
}

// O(2^N) (2^1+2^2+2^3+...+2^N => 2^(N+1))
func printAllFib(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%v fib: %v\n", i, fib(i))
	}
}

func fib(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
