package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("main")

	narr := []int{rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100)}
	fmt.Printf("list of random int: %v\n", narr)

	min, max := minMaxV1(narr)
	fmt.Printf("minMaxV1: min - %v; max - %v\n", min, max)

	min, max = minMaxV2(narr)
	fmt.Printf("minMaxV2: min - %v; max - %v\n", min, max)
}

// time - O(N)
func minMaxV1(narr []int) (int, int) {
	if len(narr) == 0 {
		return 0, 0
	}
	min := narr[0]
	max := narr[0]
	for _, n := range narr[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

// cpu - O(N)
func minMaxV2(narr []int) (int, int) {
	if len(narr) == 0 {
		return 0, 0
	}
	min := narr[0]
	max := narr[0]
	for _, n := range narr[1:] {
		if n < min {
			min = n
		}
	}
	for _, n := range narr[1:] {
		if n > max {
			max = n
		}
	}
	return min, max
}
