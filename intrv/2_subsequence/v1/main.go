package main

import "fmt"

func main() {

	arr := []int{1, 5, 2, 3, 5, 4, 6}
	// exp subseq [1,2,3,5,6] or [1,2,3,4,6]
	expLn := 5

	fmt.Println("max sub seq len:", findMaxSubSeqLen(arr), "exp len:", expLn)
}

func findMaxSubSeqLen(arr []int) int {
	return 0
}
