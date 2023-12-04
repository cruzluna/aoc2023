package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	grid := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	visited := make([][]bool, len(grid))
	for row := range visited {
		visited[row] = make([]bool, len(grid[0]))
	}

	total := 0
	// transverse the grid
	for i := range grid {
		for j := range grid[i] {
			// period
			if grid[i][j] == 46 {
				continue
			}
			// visited
			if visited[i][j] {
				continue
			}

			if grid[i][j] >= 33 && grid[i][j] <= 47 && grid[i][j] != 46 {
				// dfs from the symbol
				if str := dfs(grid, visited, i, j, ""); str != "" {
					fmt.Println("DONE: ", str)
					cnt, _ := strconv.Atoi(str)
					total = total + cnt
				}
			}
			// once we find a sybol, add that to the total
		}
	}
	// fmt.Println("VISITED\n", visited)
	fmt.Println("TOTAL: ", total)
}

func dfs(grid [][]rune, visited [][]bool, x, y int, buildString string) string {
	// boundary check
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return ""
	}

	if visited[x][y] {
		return ""
	}

	// period, now we finish
	if grid[x][y] == 46 {
		return ""
	}

	// found a symbol
	// if grid[x][y] >= 33 && grid[x][y] <= 47 && grid[x][y] != 46 {
	// 	// fmt.Println("MADE IT IN HERE: ", buildString)
	// 	return buildString
	// }

	visited[x][y] = true

	fmt.Printf("grid[x][y]: %v , bs: %s , uni %v\n", string(grid[x][y]), buildString, grid[x][y])
	buildString = buildString + string(grid[x][y])
	// fmt.Println("\nNew Str: ", buildString)

	for _, dir := range getDirections() {
		dx, dy := dir[0], dir[1]
		res := dfs(grid, visited, x+dx, y+dy, buildString)
		if res != "" {
			return res
		}
	}

	return buildString
}

func getDirections() [][]int {
	return [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		// diagonals
		{-1, 1},
		{1, -1},
		{-1, -1},
		{1, 1},
	}
}
