package main

import (
	"fmt"
	"strings"
)

func parse(input string) [][]string {
	lines := strings.Split(input, "\n")
	grid := make([][]string, len(lines))
	for i, l := range lines {
		row := strings.Split(strings.TrimSpace(l), "")
		grid[i] = row
	}
	return grid
}

func getNeighbors(grid [][]string, i, j int) [][]int {
	moves := [][]int{{0, 1}, {1, 0}}
	neighbors := [][]int{}
	for _, v := range moves {
		di, dj := v[0], v[1]
		if di+i < len(grid) && di+i >= 0 && dj+j < len(grid[0]) && dj+j >= 0 && grid[i][j] != "*" {
			neighbors = append(neighbors, []int{di + i, dj + j})
		}
	}
	return neighbors
}

func id(i, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}

func gridPathsRecursive(grid [][]string, i, j int) int {

	if len(grid)-1 == i && len(grid[0])-1 == j {
		return 1
	}

	sum := 0
	for _, n := range getNeighbors(grid, i, j) {
		sum += gridPathsRecursive(grid, n[0], n[1])
	}

	return sum

}

// func gridPathsMemoized(number int, m *map[int]int) int {

// }

// func gridPathsBottomUp(number int) int {

// }

func main() {
	grid := parse(`....
	.*..
	...*
	*...`)
	fmt.Println(gridPathsRecursive(grid, 0, 0))
	// fmt.Println(gridPathsMemoized(4434, &map[int]int{}))
	// fmt.Println(gridPathsBottomUp(4434))

}
