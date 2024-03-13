package main

import "fmt"

func main() {
	fmt.Println("main")

	s := "qwerty"
	permut(s, "")
}

// O(N^2*N!), where N - number of chars
// O(N!) is number of permutations.
// If all permutations is leaves of tree, when we have N paths of call "permut" before generate each leaf.
// Also we have O(N) operations to print each chars of each permuts.
// As result we get O(N! * N * N) = O(N^2*N!)
func permut(s, pref string) {
	if len(s) == 0 {
		fmt.Printf("next permutuation is %v\n", pref)
	} else {
		for i := 0; i < len(s); i++ {
			rem := s[0:i] + s[i+1:]
			permut(rem, pref+string(s[i]))
		}
	}
}
