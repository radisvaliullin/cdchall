package game

import "fmt"

// for testing
func printGrid(grid [][]int) {
	for _, r := range grid {
		fmt.Printf("%+v\n", r)
	}
}
