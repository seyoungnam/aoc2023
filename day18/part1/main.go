package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dirc = map[string][2]int{
	"R": {0, 1},
	"L": {0, -1},
	"U": {-1, 0},
	"D": {1, 0},
}

func main() {
	plan, num := loadFile("../input.txt")
	fmt.Println(plan)
	fmt.Println(num)

	r, l, u, d := getLengthsAndStartingPoint(plan, num)
	fmt.Println(r, l, u, d)
	grid := createEmptyGrid(r, l, u, d)
	// fmt.Println("Empty grid: ")
	// for _, row := range grid {
	// 	fmt.Println(row)
	// }

	markGrid(plan, num, &grid, l, u)
	fmt.Println("Marked grid: ")
	for _, row := range grid {
		fmt.Println(row)
	}

	lGrid := markNumInGridFromLeft(&grid)
	// fmt.Println("Left Num Grid: ")
	// for _, row := range lGrid {
	// 	fmt.Println(row)
	// }

	markPoundInGrid(&grid, lGrid)
	fmt.Println("Pound Inserted Grid: ")
	for _, row := range grid {
		fmt.Println(row)
	}

	res := countPound(&grid)
	fmt.Printf("The answer is %v\n", res)

}

func loadFile(fileName string) (plan []string, num []int) {
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		plan = append(plan, arr[0])
		n, _ := strconv.Atoi(arr[1])
		num = append(num, n)
	}
	return plan, num
}

func getLengthsAndStartingPoint(plan []string, num []int) (r, l, u, d int) {
	r, l, u, d = 0, 0, 0, 0
	x, y := 0, 0
	for i := 0; i < len(plan); i++ {
		if plan[i] == "R" {
			x += num[i]
			if x > r {
				r = x
			}
		} else if plan[i] == "L" {
			x -= num[i]
			if x < l {
				l = x
			}
		} else if plan[i] == "D" {
			y += num[i]
			if y > d {
				d = y
			}
		} else { // U
			y -= num[i]
			if y < u {
				u = y
			}
		}
	}
	return r, l, u, d
}

func createEmptyGrid(r, l, u, d int) (grid [][]string) {
	grid = [][]string{}
	yLen := -u + d + 1
	xLen := -l + r + 1
	for y := 0; y < yLen; y++ {
		row := []string{}
		for x := 0; x < xLen; x++ {
			row = append(row, ".")
		}
		grid = append(grid, row)
	}
	return grid
}

func markGrid(plan []string, num []int, grid *[][]string, l, u int) {
	y, x := -u, -l
	for i := 0; i < len(plan); i++ {
		d, step := plan[i], num[i]
		for j := 0; j < step; j++ {
			dy, dx := dirc[d][0], dirc[d][1]
			y += dy
			x += dx
			// fmt.Println(y, x)
			(*grid)[y][x] = "#"
		}
	}
}

func markNumInGridFromLeft(grid *[][]string) (leftNumGrid [][]int) {
	leftNumGrid = [][]int{}
	for y := 0; y < len(*grid); y++ {
		row := []int{}
		curr := 0
		x := 0
		for x < len((*grid)[0]) {
			if (*grid)[y][x] == "#" {
				s, e := x, x
				for e < len((*grid)[0]) && (*grid)[y][e] == "#" {
					e++
				}
				// now (*grid)[y][e] == "."
				if e-s == 1 {
					curr++
					row = append(row, curr)
					x = e
					continue
				}
				// edge condition : do not curr++
				if y == 0 || y == len(*grid)-1 || ((*grid)[y-1][s] == "#" && (*grid)[y-1][e-1] == "#") || ((*grid)[y+1][s] == "#" && (*grid)[y+1][e-1] == "#") {
					for s < e {
						row = append(row, curr)
						s++
					}
				} else {
					curr++
					for s < e {
						row = append(row, curr)
						s++
					}
				}
				x = e
			} else {
				row = append(row, curr)
				x++
			}
		}
		leftNumGrid = append(leftNumGrid, row)
	}
	return leftNumGrid
}

func markPoundInGrid(grid *[][]string, lGrid [][]int) {
	for y := 0; y < len(*grid); y++ {
		for x := 0; x < len((*grid)[0]); x++ {
			if (*grid)[y][x] == "." && lGrid[y][x]%2 == 1 {
				(*grid)[y][x] = "#"
			}
		}
	}
}

func countPound(grid *[][]string) (res int) {
	for y := 0; y < len(*grid); y++ {
		for x := 0; x < len((*grid)[0]); x++ {
			if (*grid)[y][x] == "#" {
				res++
			}
		}
	}
	return res
}
