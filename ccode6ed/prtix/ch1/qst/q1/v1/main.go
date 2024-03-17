package main

import (
	"fmt"
	"slices"
)

// check that string has only uniq chars
// what if not use additional struct
func main() {

	s := "qwertyqwerty"
	fmt.Println("string -", s, "has duplicate chars -", !isUniqCharStrV1(s))
	fmt.Println("string -", s, "has duplicate chars -", !isUniqCharStrV2(s))
	fmt.Println("string -", s, "has duplicate chars -", !isUniqCharStrV3(s))
	fmt.Println("string -", s, "has duplicate chars -", !isUniqCharStrV4(s))
	s = "asdf"
	fmt.Println("string -", s, "has duplicate chars -", !isUniqCharStrV1(s))
	fmt.Println("string -", s, "has duplicate chars -", !isUniqCharStrV2(s))
	fmt.Println("string -", s, "has duplicate chars -", !isUniqCharStrV3(s))
	fmt.Println("string -", s, "has duplicate chars -", !isUniqCharStrV4(s))
}

// using help data struct
// cpu - O(N) where N len of s, or O(1) since alphabe limited const value 128, or O(c) where c size of alphabet
// mem - O(1) or O(c)
func isUniqCharStrV1(s string) bool {
	// if only ASCII chars
	// if len more than chars in alphabet it meas we have duplicates
	if len(s) > 128 {
		return false
	}
	used_chars := make([]bool, 128)
	for _, ch := range s {
		if used_chars[ch] {
			return false
		}
		used_chars[ch] = true
	}
	return true
}

// using help data struct
// mem - O(1)
func isUniqCharStrV2(s string) bool {
	// if only ASCII chars subset
	// if len more than chars in alphabet it meas we have duplicates
	if len(s) > 128 {
		return false
	}
	mask0 := uint64(0)
	mask1 := uint64(0)
	for _, ch := range s {
		pos := ch
		mask := mask0
		if ch > 63 {
			pos = pos - 64
			mask = mask1
		}
		if (mask & (1 << pos)) > 0 {
			return false
		}
		if ch > 63 {
			mask1 |= (1 << pos)
		} else {
			mask0 |= (1 << pos)
		}
	}
	return true
}

// no help data struct
// cpu - O(N^2)
func isUniqCharStrV3(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

// no help data struct
// Go string immutable so actually I use help struct
// cpu - O(NlogN)
func isUniqCharStrV4(s string) bool {
	chrs := make([]rune, len(s))
	for i, r := range s {
		chrs[i] = r
	}
	slices.Sort(chrs)
	for i := 0; i < len(chrs)-1; i++ {
		j := i + 1
		if chrs[i] == chrs[j] {
			return false
		}
	}
	return true
}
