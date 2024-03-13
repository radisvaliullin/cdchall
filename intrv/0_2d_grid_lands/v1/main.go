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

// Complexity
// O(N^2) where N=n*m, it is N^2 because for each N we need check in Lands list (worse case ~N size)
// space: O(N^2) in worse case we have N size grid and N size lands.
func findLands(grid [][]int) int {
	// list of lands, extend by interations
	lands := [][]cell{}

	// iterate through all cells
	for i, row := range grid {
		for j, v := range row {
			if v == 0 {
				continue
			}
			c := cell{x: i, y: j}
			bCells := getLeftAndUpLandBorderCells(i, j, grid)
			// check if bordered cells already in list of lands
			// if not add new land to list
			// if yes add to existing land
			if len(bCells) == 0 {
				l := []cell{c}
				lands = append(lands, l)
			} else {
				if len(bCells) == 1 {
					// we iterate from left to right, from up to down, we add cell to land in same order
					// so we always should get land idx, not need check OK
					lIdx, _ := getBorderCellLandIdx(bCells[0], lands)
					lands[lIdx] = append(lands[lIdx], c)
					continue
				}
				// or we have two border cells
				lIdx0, _ := getBorderCellLandIdx(bCells[0], lands)
				lIdx1, _ := getBorderCellLandIdx(bCells[1], lands)
				if lIdx0 == lIdx1 {
					lands[lIdx0] = append(lands[lIdx0], c)
					continue
				}
				// if two land idx not equal we should merge it because it is same land via current cell
				lands[lIdx0] = append(lands[lIdx0], lands[lIdx1]...)
				lands = append(lands[:lIdx1], lands[lIdx1+1:]...)
			}
		}
	}
	numLand := len(lands)
	return numLand
}

type cell struct {
	x int
	y int
}

// max possible value 2 cells, one left and one right
func getLeftAndUpLandBorderCells(x, y int, grid [][]int) []cell {
	cells := []cell{}
	if (x-1) >= 0 && grid[x-1][y] == 1 {
		cells = append(cells, cell{x: x - 1, y: y})
	}
	if (y-1) >= 0 && grid[x][y-1] == 1 {
		cells = append(cells, cell{x: x, y: y - 1})
	}
	return cells
}

// return land index in lands for cell
func getBorderCellLandIdx(c cell, lands [][]cell) (int, bool) {
	for i, l := range lands {
		for _, lc := range l {
			if lc.x == c.x && lc.y == c.y {
				return i, true
			}
		}
	}
	return 0, false
}
