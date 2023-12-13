package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	grid := loadFile("../input.txt")
	grids := sortGrid(grid)
	for _, grid := range grids {
		fmt.Println(grid)
	}

	res := 0
	for _, grid := range grids {
		res += scanCol(grid)
		res += scanRow(grid)
	}
	fmt.Println(res)

}

func loadFile(fileName string) (lines [][]string) {
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Split(line, "")
		lines = append(lines, lineArr)
	}
	return lines
}

func scanCol(grid [][]string) int {
	// ms := []int{}
	// if len(grid[0])%2 != 0 {
	// 	k := len(grid[0]) / 2
	// 	ms = append(ms, k, k+1)
	// } else {
	// 	ms = append(ms, len(grid[0])/2)
	// }
	// fmt.Printf("scanCol ms=%v\n", ms)

	// for _, m := range ms {
	// 	if isColMirrored(grid, m) {
	// 		return m + 1
	// 	}
	// }

	for m := 0; m < len(grid[0]); m++ {
		if isColMirrored(grid, m) {
			fmt.Printf("colMirrored=%v\n", m+1)
			return m + 1
		}
	}
	return 0
}

func scanRow(grid [][]string) int {
	// ms := []int{}
	// k := len(grid) / 2
	// if len(grid)%2 != 0 {
	// 	ms = append(ms, k, k+1)
	// } else {
	// 	ms = append(ms, k)
	// }
	// fmt.Printf("scanRow ms=%v\n", ms)

	// for _, m := range ms {
	// 	if isRowMirrored(grid, m) {
	// 		return (m + 1) * 100
	// 	}
	// }

	for m := 0; m < len(grid); m++ {
		if isRowMirrored(grid, m) {
			fmt.Printf("rowMirrored=%v\n", m+1)
			return (m + 1) * 100
		}
	}
	return 0
}

func sortGrid(grid [][]string) (grids [][][]string) {
	newGrid := [][]string{}
	for _, line := range grid {
		if len(line) == 0 {
			grids = append(grids, newGrid)
			newGrid = [][]string{}
			continue
		}
		newGrid = append(newGrid, line)

	}

	if len(newGrid) != 0 {
		grids = append(grids, newGrid)
	}

	return grids
}

func isColMirrored(grid [][]string, m int) bool {
	i, j := m, m+1
	if i < 0 || j >= len(grid[0]) {
		return false
	}
	smudge := 1
	for i >= 0 && j < len(grid[0]) {
		skip := false
		for y := 0; y < len(grid); y++ {
			if grid[y][i] != grid[y][j] {
				smudge--
				if smudge < 0 {
					skip = true
					break
				}
			}
		}
		if skip {
			return false
		}
		i--
		j++
	}

	if smudge == 1 {
		return false
	}
	return true
}

func isRowMirrored(grid [][]string, m int) bool {
	i, j := m, m+1
	if i < 0 || j >= len(grid) {
		return false
	}
	smudge := 1
	for i >= 0 && j < len(grid) {
		skip := false
		for x := 0; x < len(grid[0]); x++ {
			if grid[i][x] != grid[j][x] {
				smudge--
				if smudge < 0 {
					skip = true
					break
				}
			}
		}
		if skip {
			return false
		}
		i--
		j++
	}

	if smudge == 1 {
		return false
	}
	return true
}
