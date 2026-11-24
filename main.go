package main

import "fmt"

// NumIslands counts the 4-connected groups of '1' cells in a grid. Each
// discovered island is sunk (marked '0') by DFS so it is counted once.
// The grid is modified; pass a copy to preserve it.
func NumIslands(grid [][]byte) int {
	count := 0
	var sink func(r, c int)
	sink = func(r, c int) {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || grid[r][c] != '1' {
			return
		}
		grid[r][c] = '0'
		sink(r+1, c)
		sink(r-1, c)
		sink(r, c+1)
		sink(r, c-1)
	}
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '1' {
				count++
				sink(r, c)
			}
		}
	}
	return count
}

func main() {
	grid := [][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	}
	fmt.Println(NumIslands(grid))
}
