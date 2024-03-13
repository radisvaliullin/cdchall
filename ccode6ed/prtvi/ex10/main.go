package main

import "fmt"

func main() {

	fmt.Println("main")

	n := 73
	fmt.Printf("is %v prime: %v\n", n, isPrime(n))
}

// O(sqrt(N))
func isPrime(n int) bool {
	for x := 2; x*x <= n; x++ {
		if n%2 == 0 {
			return false
		}
	}
	return true
}
