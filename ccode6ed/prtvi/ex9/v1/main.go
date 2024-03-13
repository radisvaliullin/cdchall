package main

import (
	"fmt"
	"math/rand"
)

const (
	randRange = 100
)

func main() {
	fmt.Println("main")

	nd := genNode(2, nil)
	fmt.Println("nd: ", nd)
	printTree(0, nd)

	s := sum(nd)
	fmt.Printf("sum of tree: %v\n", s)
}

// O(N)
func sum(nd *Node) int {
	if nd == nil {
		return 0
	}
	return sum(nd.left) + nd.val + sum(nd.right)
}

// Node
type Node struct {
	left  *Node
	right *Node
	val   int
}

// genNode generates balanced binary search tree
func genNode(lv int, nd *Node) *Node {
	if nd == nil {
		nd = &Node{
			val: rand.Intn(randRange),
		}
	}
	if lv <= 0 {
		return nil
	}
	lv--
	nd.left = &Node{
		val: rand.Intn(randRange),
	}
	nd.right = &Node{
		val: rand.Intn(randRange),
	}
	genNode(lv, nd.left)
	genNode(lv, nd.right)
	return nd
}

func printTree(lv int, nd *Node) {
	if nd == nil {
		return
	}
	fmt.Printf("lv: %v; nd - %v; left - %v; right - %v\n", lv, nd, nd.left, nd.right)
	printTree(lv+1, nd.left)
	printTree(lv+1, nd.right)
}
