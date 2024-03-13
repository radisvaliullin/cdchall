package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("main")

	narr0 := []int{rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100)}
	narr1 := []int{rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100)}

	printOneByOne(narr0, narr1)
	printMixed(narr0, narr1)
}

// time - O(N0+N1)
func printOneByOne(narr0, narr1 []int) {
	for _, n := range narr0 {
		fmt.Println(n)
	}
	for _, n := range narr1 {
		fmt.Println(n)
	}
}

// time - O(N0*N1)
func printMixed(narr0, narr1 []int) {
	for _, n0 := range narr0 {
		for _, n1 := range narr1 {
			fmt.Println(n0, n1)
		}
	}
}
