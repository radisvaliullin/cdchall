package main

import (
	"fmt"
)

/*
given a 2d grid
[

	[1, 1, 1, 0, 0, 0],
	[1, 1, 1, 0, 0, 0],
	[1, 1, 1, 0, 0, 0],
	[0, 0, 0, 1, 0, 0],
	[0, 0, 0, 0, 1, 1],
	[0, 0, 0, 0, 1, 1],

]
where 1 is land, 0 sea.
Write function accepting grid and returning number of independent lands.
If two land cells bordered vertically and horizontally then they dependent.
If not bordered or bordered diagonally then independent.
On example above output value should be 3.
*/
func main() {

	grid := [][]int{
		{1, 1, 1, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 1, 1},
	}
	fmt.Printf("grid has %v independent lands\n", findLands(grid))

	grid2 := [][]int{
		{1, 1, 1, 1},
		{0, 0, 0, 1},
		{0, 0, 0, 1},
		{1, 1, 1, 1},
	}
	fmt.Printf("grid has %v independent lands\n", findLands(grid2))

	grid3 := [][]int{
		{1, 1, 1, 1},
		{1, 0, 0, 1},
		{1, 0, 0, 1},
		{1, 1, 1, 1},
	}
	fmt.Printf("grid has %v independent lands\n", findLands(grid3))

	grid4 := [][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 1, 0, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 1, 0, 1},
		{1, 1, 1, 1, 1},
	}
	fmt.Printf("grid has %v independent lands\n", findLands(grid4))

	grid5 := [][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 1, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}
	fmt.Printf("grid has %v independent lands\n", findLands(grid5))
}

// O(N), where N=m*n (time and space) (recursive approach)
func findLands(grid [][]int) int {

	landNum := 0
	replaceTarget := 2

	// iterate through all cells
	for i, row := range grid {
		for j, v := range row {
			if v != 1 {
				continue
			}

			landNum++
			c := cell{x: i, y: j}
			// replaces all bordered cells with target value
			findAllBorderLandCellsAndMarkRecursive(c, grid, len(grid), len(grid[0]), replaceTarget)
		}
	}

	// restore land cells value
	for i, row := range grid {
		for j, v := range row {
			if v == replaceTarget {
				grid[i][j] = 1
			}
		}
	}

	return landNum
}

type cell struct {
	x int
	y int
}

// replace land cell with target value
func findAllBorderLandCellsAndMarkRecursive(c cell, grid [][]int, xLen int, yLen int, t int) {

	// reset current value
	grid[c.x][c.y] = t

	// find border land cells and recall recursively
	if (c.x-1) >= 0 && grid[c.x-1][c.y] == 1 {
		findAllBorderLandCellsAndMarkRecursive(cell{x: c.x - 1, y: c.y}, grid, xLen, yLen, t)
	}
	if (c.y-1) >= 0 && grid[c.x][c.y-1] == 1 {
		findAllBorderLandCellsAndMarkRecursive(cell{x: c.x, y: c.y - 1}, grid, xLen, yLen, t)
	}
	if (c.x+1) < xLen && grid[c.x+1][c.y] == 1 {
		findAllBorderLandCellsAndMarkRecursive(cell{x: c.x + 1, y: c.y}, grid, xLen, yLen, t)
	}
	if (c.y+1) < yLen && grid[c.x][c.y+1] == 1 {
		findAllBorderLandCellsAndMarkRecursive(cell{x: c.x, y: c.y + 1}, grid, xLen, yLen, t)
	}
}
