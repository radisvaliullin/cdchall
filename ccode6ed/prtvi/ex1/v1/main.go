package main

import "fmt"

func main() {
	fmt.Println("main")

	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sum, prod := foo(a)
	fmt.Printf("sum and prod: %v, %v\n", sum, prod)
}

func foo(arr []int) (sum, prod int) {
	sum = 0
	prod = 1
	for _, v := range arr {
		sum += v
	}
	for _, v := range arr {
		prod *= v
	}
	return sum, prod
}
