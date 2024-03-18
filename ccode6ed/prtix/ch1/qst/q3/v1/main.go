package main

import (
	"fmt"
)

// replace space with %20
// Input: "qwerty asdf zxcv    ", Output: "qwerty%20asdf%20zxcv"
func main() {

	str := "日の出qwerty adsf zcxv    "
	fmt.Printf("str <%v> with replaces spaces <%v>\n", str, replaceSpaces(str))
}

// cpu - O(N)
// mem - O(N) (we use temp rune slice, we can not convert slice (rune or byte) to string without allocation new space in Golang)
func replaceSpaces(s string) string {
	charCntr := 0
	spaceCntr := 0
	spacesTailSize := 0
	for _, ch := range s {
		charCntr++
		if ch == ' ' {
			spaceCntr++
			spacesTailSize++
		} else {
			// if not tail need set 0
			spacesTailSize = 0
		}
	}
	// each space (one char) we need repace with %20 (three chars)
	// we already have rune for one char of each space so we need just add 2 more for each space
	newSchars := make([]rune, len(s)+spaceCntr*2-spacesTailSize)
	// we need index of char, we can not use i becuase it is index of bytes
	charIdx := 0
	// we need delta var to calculate index in newSchars because for each space we add two more chars
	// for each space char we increase delta for two
	delta := 0
	// iterate from string but insert to rune slice
	for i, ch := range s {
		// if tail spaces break
		if i > len(s)-spacesTailSize-1 {
			break
		}
		charIdx++
		if s[i] == ' ' {
			newSchars[charIdx+delta] = '%'
			delta++
			newSchars[charIdx+delta] = '2'
			delta++
			newSchars[charIdx+delta] = '0'
		} else {
			newSchars[charIdx+delta] = ch
		}
	}
	return string(newSchars)
}
