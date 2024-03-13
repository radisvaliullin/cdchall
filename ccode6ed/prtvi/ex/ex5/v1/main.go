package main

import "fmt"

func main() {
	fmt.Println("main")

	arrA := []int{43, 73, 1, 2, 3, 4}
	arrB := []int{44, 55, 66}
	printPair(arrA, arrB)
}

// O(a*b)
func printPair(arrA, arrB []int) {
	for i := 0; i < len(arrA); i++ {
		for j := 0; j < len(arrB); j++ {
			for k := 0; k < 10_000; k++ {
				fmt.Printf("pair: %v - %v - %v\n", arrA[i], arrB[j], k)
			}
		}
	}
}
