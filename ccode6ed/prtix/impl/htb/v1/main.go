package main

import (
	"fmt"

	"github.com/radisvaliullin/cdchall/ccode6ed/prtix/impl/htb/v1/htb"
)

// Implement simple hash table.
func main() {

	ht := htb.NewHTB(0)
	ht.Add("one", "1")
	ht.Add("three", "3")
	fmt.Printf("hash table: %+v\n", ht)
	k := "one"
	v, ok := ht.Get(k)
	fmt.Printf("hash table: k - %v, v - %v, %v\n", k, v, ok)
	k = "two"
	v, ok = ht.Get(k)
	fmt.Printf("hash table: k - %v, v - %v, %v\n", k, v, ok)
	//
	ht.Del("three")
	fmt.Printf("hash table: %+v\n", ht)
}
