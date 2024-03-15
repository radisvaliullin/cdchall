package main

import "fmt"

func main() {
	fmt.Println("main")

	printAllFib(7)
}

// cpu - O(N) (instead of 2^N because we use memorization cache)
func printAllFib(n int) {
	// preallocate slice with size N
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("%v fib: %v\n", i, fib(i, memo))
	}
}

func fib(n int, memo []int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if memo[n] > 0 {
		return memo[n]
	}
	memo[n] = fib(n-1, memo) + fib(n-2, memo)
	return memo[n]
}
