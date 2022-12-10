package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const input = "./8/input.txt"

func populateGrid(input []string) [][]int {
	var ret [][]int
	for _, line := range input {
		var gridLine []int
		for _, letter := range strings.Split(line, "") {
			if tree, err := strconv.Atoi(letter); err == nil {
				gridLine = append(gridLine, tree)
			}
		}
		ret = append(ret, gridLine)
	}
	return ret
}

// isVisible takes [row,col] co-ordinates of a point and the grid they reside within and indicates whether the tree is visible
func isVisible(currentTreeRow, currentTreeCol int, grid [][]int) bool {
	currentTreeHeight := grid[currentTreeRow][currentTreeCol]
	if currentTreeRow == 0 || currentTreeCol == 0 {
		return true
	}
	if currentTreeRow-1 == len(grid[0]) {
		return true
	}
	if currentTreeCol-1 == len(grid) {
		return true
	}

	// assumed visible unless checks say otherwise
	var top, bot, left, right = true, true, true, true
	for row := 0; row < len(grid); row++ {
		if grid[row][currentTreeCol] >= currentTreeHeight && row < currentTreeRow {
			top = false
		}
		if row == currentTreeRow {
			for col := 0; col < len(grid[row]); col++ {
				if grid[currentTreeRow][col] >= currentTreeHeight && col < currentTreeCol {
					left = false
				}
				if col == currentTreeCol {
					continue
				}
				if grid[currentTreeRow][col] >= currentTreeHeight && col > currentTreeCol {
					right = false
				}
			}
		}
		if grid[row][currentTreeCol] >= currentTreeHeight && row > currentTreeRow {
			bot = false
			if !bot && !top && !left && !right {
				return false
			}
		}
	}

	// if visible from any of these angles
	if top || bot || left || right {
		return true
	}
	return false
}

func scenicScore(currentTreeRow, currentTreeCol int, grid [][]int) int {
	// assumed visible unless checks say otherwise
	up, down, left, right := 1, 1, 1, 1
	currentTreeHeight := grid[currentTreeRow][currentTreeCol]

	if currentTreeRow == 0 {
		return 0
	}
	if currentTreeCol == 0 {
		return 0
	}
	if currentTreeRow-1 == len(grid[0]) {
		return 0
	}
	if currentTreeCol-1 == len(grid) {
		return 0
	}

	// Going UP from current tree!
	for row := currentTreeRow - 1; row >= 0 && grid[row][currentTreeCol] < currentTreeHeight; row-- {
		if row > 0 {
			up++
		}
	}

	// Going DOWN from current tree!
	for row := currentTreeRow + 1; row < len(grid) && grid[row][currentTreeCol] < currentTreeHeight; row++ {
		if row < len(grid)-1 {
			down++
		}
	}

	// Going LEFT from current tree!
	for col := currentTreeCol - 1; col >= 0 && grid[currentTreeRow][col] < currentTreeHeight; col-- {
		if col > 0 {
			left++
		}
	}

	// Going RIGHT from current tree!
	for col := currentTreeCol + 1; col < len(grid) && grid[currentTreeRow][col] < currentTreeHeight; col++ {
		if col < len(grid[currentTreeRow])-1 {
			right++
		}
	}
	return up * down * left * right
}

func main() {
	data, err := os.ReadFile(input)
	if err != nil {
		panic("Could not open file: " + err.Error())
	}
	inputLines := strings.Split(bytes.NewBuffer(data).String(), "\n")
	grid := populateGrid(inputLines)

	visibleTrees := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// iterate through points to check
			if isVisible(i, j, grid) {
				visibleTrees++
			}
		}
	}
	fmt.Printf("%d trees are visible from outside the grid\n", visibleTrees)

	score := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// iterate through points to check
			temp := scenicScore(i, j, grid)
			if temp > score {
				score = temp
			}
		}
	}
	fmt.Printf("%d is the highest scenic score for any tree\n", score)
}
