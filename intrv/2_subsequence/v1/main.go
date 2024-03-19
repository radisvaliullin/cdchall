package main

import (
	"fmt"
	"slices"
)

func main() {

	arr := []int{1, 5, 2, 3, 5, 4, 6}
	// exp subseq [1,2,3,5,6] or [1,2,3,4,6]
	expLn := 5
	fmt.Println("max sub seq len:", findMaxSubSeqLen(arr), "exp len:", expLn)

	arr = []int{1, 5, 2, 3, 5, 4, 6, 0, 0, 7, 12, 24, 8, 9, 10}
	// exp subseq [1,2,3,5,6,7,8,9,10] or [1,2,3,4,6,7,8,9,10]
	expLn = 9
	fmt.Println("max sub seq len:", findMaxSubSeqLen(arr), "exp len:", expLn)

	arr = []int{1, 5, 2, 3, 5, 4, 6, 0, 0, 7, 12, 24, 25, 26, 27, 8, 9, 10}
	// exp subseq [1,2,3,5,6,7,12,24,25,26,27]
	expLn = 11
	fmt.Println("max sub seq len:", findMaxSubSeqLen(arr), "exp len:", expLn)
	fmt.Println(findMaxSubSeq(arr))
}

func findMaxSubSeqLen(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	min := slices.Min(arr)
	out := findMaxSubSeqRecursion(min, arr)
	return len(out)
}

func findMaxSubSeq(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	return findMaxSubSeqRecursion(slices.Min(arr), arr)
}

func findMaxSubSeqRecursion(prev int, arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	out := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] > prev {
			tailOut := findMaxSubSeqRecursion(arr[i], arr[i+1:])
			if len(tailOut)+1 > len(out) {
				out = append([]int{arr[i]}, tailOut...)
			}
		}
	}
	return out
}
