package main

import "fmt"

func main() {
	fmt.Println("main")

	arr := []int{43, 73, 1, 2, 3, 4}
	printPair(arr)
}

func printPair(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Printf("pair: %v - %v\n", arr[i], arr[j])
		}
	}
}
