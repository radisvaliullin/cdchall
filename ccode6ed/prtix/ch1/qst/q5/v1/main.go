package main

import (
	"fmt"
	"unicode/utf8"
)

// check that string has one step changes
// or one insert or one remove or one replace
// utf-8 strings
func main() {

	// one replace
	str1 := "qwerty日の出asdf"
	str2 := "qwerty日のgasdf"
	fmt.Println(isOneEditAway(str1, str2))

	// one insert
	str1 = "qwerty日の出asdf"
	str2 = "qwerty日の出asdf1"
	fmt.Println(isOneEditAway(str1, str2))

	// two insert
	str1 = "qwerty日の出asdf"
	str2 = "qwerty日の出asdf12"
	fmt.Println(isOneEditAway(str1, str2))

	// one remove
	str1 = "qwerty日の出asdf"
	str2 = "qwerty日の出asd"
	fmt.Println(isOneEditAway(str1, str2))

	// two remove
	str1 = "qwerty日の出asdf"
	str2 = "qwerty日の出asd"
	fmt.Println(isOneEditAway(str1, str2))
}

// cpu - O(N), where N - len(s1)+len(s2), we find len by iterating each string
// for ASCII we can get cpu - O(N), where N - shortest of string, if get length as length of bytes (not rune)
// memp - O(1), we do not allocate new spaces
//
// replace do not change lenght of strigns
// one insert or one remove change length for one char
func isOneEditAway(s1, s2 string) bool {
	if utf8.RuneCountInString(s1) == utf8.RuneCountInString(s2) {
		return isOneEditReplace(s1, s2)
	} else if utf8.RuneCountInString(s1)+1 == utf8.RuneCountInString(s2) {
		return isOneEditInsertFirstStr(s1, s2)
	} else if utf8.RuneCountInString(s1) == utf8.RuneCountInString(s2)+1 {
		return isOneEditInsertFirstStr(s2, s1)
	}
	return false
}

// we not compare byte by byte
// we compare rune by rune
func isOneEditReplace(s1, s2 string) bool {
	// flag for detect first diff
	isDiff := false

	idx1 := 0
	idx2 := 0
	for idx1 < len(s1) && idx2 < len(s2) {
		r, len := utf8.DecodeRuneInString(s1[idx1:])
		// if not utf8 return false
		if r == utf8.RuneError {
			return false
		}
		r2, len2 := utf8.DecodeRuneInString(s2[idx2:])
		// if not utf8 return false
		if r2 == utf8.RuneError {
			return false
		}
		if len != len2 || r != r2 {
			// we get second diff char
			if isDiff {
				return false
			}
			isDiff = true
		}
		idx1 += len
		idx2 += len2
	}

	return true
}

func isOneEditInsertFirstStr(s1, s2 string) bool {
	idx1 := 0
	idx2 := 0
	for idx1 < len(s1) && idx2 < len(s2) {
		r, len := utf8.DecodeRuneInString(s1[idx1:])
		// if not utf8 return false
		if r == utf8.RuneError {
			return false
		}
		r2, len2 := utf8.DecodeRuneInString(s2[idx2:])
		// if not utf8 return false
		if r2 == utf8.RuneError {
			return false
		}
		if len != len2 || r != r2 {
			if idx1 != idx2 {
				return false
			}
			idx2++
		} else {
			idx1 += len
			idx2 += len2
		}
	}

	return true
}
