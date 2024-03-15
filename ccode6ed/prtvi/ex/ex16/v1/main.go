package main

import "fmt"

func main() {
	fmt.Println("main")

	printPowersOf2(15)
}

// cpu - O(logN)
// print powers of 2 for numbers in range [1, N]
func printPowersOf2(n int) int {
	if n < 1 {
		return 0
	} else if n == 1 {
		fmt.Println(1)
		return 1
	} else {
		prev := printPowersOf2(n / 2)
		curr := prev * 2
		fmt.Println(curr)
		return curr
	}
}
