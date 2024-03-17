package main

import (
	"fmt"
	"slices"
)

// check that two strings is permutuation
func main() {

	str1 := "sunrise日の出"
	str2 := "日の出esirnus"
	fmt.Println(isPermut(str1, str2))
	fmt.Println(isPermutV2(str1, str2))
	str1 = "sunrise__日の出"
	str2 = "日の出  esirnus"
	fmt.Println(isPermut(str1, str2))
	fmt.Println(isPermutV2(str1, str2))
}

// spaces matter and case sensetive
// cpu - O(N*logN)
func isPermut(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	chars1 := []rune(s1)
	chars2 := []rune(s2)
	slices.Sort(chars1)
	slices.Sort(chars2)
	return slices.Equal(chars1, chars2)
}

// cpu - O(N)
// mem - ~O(N)
func isPermutV2(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	// count chars
	// preallocate proper size
	charsMap := make(map[rune]int, len(s1))
	for _, ch := range s1 {
		charsMap[ch]++
	}
	for _, ch := range s2 {
		charsMap[ch]--
		if charsMap[ch] < 0 {
			return false
		}
	}
	return true
}
