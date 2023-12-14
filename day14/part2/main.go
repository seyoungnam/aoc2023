package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	grid := loadFile("../input.txt")
	for _, row := range grid {
		fmt.Println(row)
	}

	i, scores := createLoopArray(grid)
	fmt.Println(i)
	fmt.Println(scores)
	step := (1000000000 - i) % (len(scores) - i)
	fmt.Println(scores[i+step-1])

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

func getCountTable(grid [][]string) []int {
	table := make([]int, len(grid))
	for x := 0; x < len(grid[0]); x++ {
		blockIdx := len(grid)
		i := 0
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == "O" {
				if y < blockIdx {
					table[i]++
					i++
				} else {
					i = blockIdx + 1
					table[i]++
					blockIdx = i
					i++
				}
			} else if grid[y][x] == "#" {
				blockIdx = y
			}
		}
		fmt.Println(table)
	}

	return table
}

func getTotalFromTable(table []int) (res int) {
	for i, num := range table {
		res += (len(table) - i) * num
	}
	return res
}

func tiltNorth(grid [][]string) [][]string {
	for x := 0; x < len(grid[0]); x++ {
		i := 0
		b := len(grid)
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == "O" {
				if y < b {
					if i == y {
						i++
						continue
					}
					grid[i][x] = "O"
					grid[y][x] = "."
					i++
				} else {
					i = b + 1
					if i == y {
						b = i
						i++
						continue
					}
					grid[i][x] = "O"
					grid[y][x] = "."
					b = i
					i++
				}
			} else if grid[y][x] == "#" {
				b = y
			}
		}
	}
	return grid
}

func tiltWest(grid [][]string) [][]string {
	for y := 0; y < len(grid); y++ {
		i := 0
		b := len(grid[0])
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == "O" {
				if x < b {
					if i == x {
						i++
						continue
					}
					grid[y][i] = "O"
					grid[y][x] = "."
					i++
				} else {
					i = b + 1
					if i == x {
						b = i
						i++
						continue
					}
					grid[y][i] = "O"
					grid[y][x] = "."
					b = i
					i++
				}
			} else if grid[y][x] == "#" {
				b = x
			}
		}
	}
	return grid
}

func tiltSouth(grid [][]string) [][]string {
	for x := 0; x < len(grid[0]); x++ {
		i := len(grid) - 1
		b := -1
		for y := len(grid) - 1; y > -1; y-- {
			if grid[y][x] == "O" {
				if y > b {
					if i == y {
						i--
						continue
					}
					grid[i][x] = "O"
					grid[y][x] = "."
					i--
				} else {
					i = b - 1
					if i == y {
						b = i
						i--
						continue
					}
					grid[i][x] = "O"
					grid[y][x] = "."
					b = i
					i--
				}
			} else if grid[y][x] == "#" {
				b = y
			}
		}
	}
	return grid
}

func tiltEast(grid [][]string) [][]string {
	for y := 0; y < len(grid); y++ {
		i := len(grid[0]) - 1
		b := -1
		for x := len(grid[0]) - 1; x > -1; x-- {
			if grid[y][x] == "O" {
				if x > b {
					if i == x {
						i--
						continue
					}
					grid[y][i] = "O"
					grid[y][x] = "."
					i--
				} else {
					i = b - 1
					if i == x {
						b = i
						i--
						continue
					}
					grid[y][i] = "O"
					grid[y][x] = "."
					b = i
					i--
				}
			} else if grid[y][x] == "#" {
				b = x
			}
		}
	}
	return grid
}

func makeCycle(grid [][]string) [][]string {
	return tiltEast(tiltSouth(tiltWest(tiltNorth(grid))))
}

func getLoad(grid [][]string) int {
	table := make([]int, len(grid))
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == "O" {
				table[y]++
			}
		}
	}
	return getTotalFromTable(table)
}

func createLoopArray(inputGrid [][]string) (int, []int) {
	grids := [][][]string{}
	loads := []int{}

	grid := copy(inputGrid)
	i := 0
	for i < 1000 {
		// fmt.Println("input Grid:")
		// for _, row := range grid {
		// 	fmt.Println(row)
		// }
		newGrid := makeCycle(grid)
		// fmt.Println("after Grid:")
		// for _, row := range newGrid {
		// 	fmt.Println(row)
		// }
		load := getLoad(newGrid)
		for i, g := range grids {
			if isGridSame(g, newGrid) {
				fmt.Println(len(loads))
				return i, loads
			}
		}
		grids = append(grids, newGrid)
		// fmt.Println(grids)
		loads = append(loads, load)
		// fmt.Println("Grids:")
		// for _, gd := range grids {
		// 	for _, row := range gd {
		// 		fmt.Println(row)
		// 	}
		// }
		i++
		grid = copy(newGrid)
	}

	return i, loads
}

func isGridSame(g1 [][]string, g2 [][]string) bool {
	// fmt.Printf("g1=%v\n", g1)
	// fmt.Printf("g2=%v\n", g2)
	for y := 0; y < len(g1); y++ {
		for x := 0; x < len(g1[0]); x++ {
			if g1[y][x] != g2[y][x] {
				return false
			}
		}
	}
	return true
}

func copy(g [][]string) [][]string {
	g2 := [][]string{}
	for y := 0; y < len(g); y++ {
		row := []string{}
		for x := 0; x < len(g[0]); x++ {
			row = append(row, g[y][x])
		}
		g2 = append(g2, row)
	}
	return g2
}
