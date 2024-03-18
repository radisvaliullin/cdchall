package main

import (
	"fmt"
	"unicode"
)

// check that string is permutuation of palindrom
func main() {

	str := "qwert y1234Trewq"
	fmt.Printf("%v\n", isPermutOfPalin(str))
	fmt.Printf("%v\n", isPermutOfPalinV2(str))
	fmt.Printf("%v\n", isPermutOfPalinV3onlyLatinLetter(str))
	str = "qwer y1234Trewq"
	fmt.Printf("%v\n", isPermutOfPalin(str))
	fmt.Printf("%v\n", isPermutOfPalinV2(str))
	fmt.Printf("%v\n", isPermutOfPalinV3onlyLatinLetter(str))
}

// cpu - O(N)
// mem - O(N)
func isPermutOfPalin(s string) bool {
	charCntr := make(map[rune]int, len(s))
	// count each letter chars
	for _, ch := range s {
		if unicode.IsLetter(ch) {
			lowerCh := unicode.ToLower(ch)
			charCntr[lowerCh]++
		}
	}
	// check is palindrom
	foundOdd := false
	for _, cntr := range charCntr {
		if cntr%2 == 1 {
			// if we already have odd char it is not palindrom
			// only one char can be odd number if we have text with odd number of lenght of letter
			if foundOdd {
				return false
			}
			foundOdd = true
		}
	}
	return true
}

// cpu - O(N) (same complexity but one loop less)
// mem - O(N)
func isPermutOfPalinV2(s string) bool {
	oddCntr := 0
	charCntr := make(map[rune]int, len(s))
	// count each letter chars
	for _, ch := range s {
		if unicode.IsLetter(ch) {
			lowerCh := unicode.ToLower(ch)
			charCntr[lowerCh]++
			// if we have char with odd number after end of loop we will have incremented counter
			if charCntr[lowerCh]%2 == 1 {
				oddCntr++
			} else {
				oddCntr--
			}
		}
	}
	// check is palindrom
	return oddCntr <= 1
}

// O(N)
// mem - O(1)
func isPermutOfPalinV3onlyLatinLetter(s string) bool {
	// use bit vector for check odd or even number of char repeat in string
	bitVector := uint64(0)
	for _, ch := range s {
		// only ASCII Latin
		if unicode.IsLetter(ch) && ch < unicode.MaxASCII {
			lowerCh := unicode.ToLower(ch)
			idx := lowerCh - 'a'
			// XOR (1 if only one of two is 1)
			// always flip 0 to 1 and 1 to 0
			// if char number even we will have 0 if odd then 1 (end of loop)
			bitVector ^= 1 << idx
		}
	}
	// check that only one bit is set
	isOnlyOneBit := bitVector&(bitVector-1) == 0
	return bitVector == 0 || isOnlyOneBit
}
