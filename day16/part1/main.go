package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var dir = [4][2]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

var mark = [4]string{
	"<",
	"v",
	">",
	"^",
}

func main() {
	grid := loadFile("../input.txt")
	fmt.Println("original grid: ")
	for _, row := range grid {
		fmt.Println(row)
	}

	memo := makeEmptyGrid(grid)
	fmt.Println("original memo: ")
	for _, row := range memo {
		fmt.Println(row)
	}

	scanGrid(0, 0, 2, &grid, &memo)

	fmt.Println("marked grid: ")
	for _, row := range grid {
		fmt.Println(row)
	}

	fmt.Println("original memo: ")
	for _, row := range memo {
		fmt.Println(row)
	}

	res := getPoundNumber(memo)
	fmt.Printf("The answer is %v\n", res)

}

func loadFile(fileName string) (grid [][]string) {
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Split(line, "")
		grid = append(grid, lineArr)
	}
	return grid
}

func makeEmptyGrid(grid [][]string) [][]string {
	emptyGrid := [][]string{}
	for y := 0; y < len(grid); y++ {
		row := []string{}
		for x := 0; x < len(grid[0]); x++ {
			row = append(row, ".")
		}
		emptyGrid = append(emptyGrid, row)
	}
	return emptyGrid
}

func scanGrid(y, x, d int, grid, memo *[][]string) {
	// fmt.Println(y, x, y+dir[d][0], x+dir[d][1])
	if y < 0 || y >= len(*grid) || x < 0 || x >= len((*grid)[0]) {
		return
	}
	if (*grid)[y][x] == "." {
		(*memo)[y][x] = "#"
		(*grid)[y][x] = mark[d]
		scanGrid(y+dir[d][0], x+dir[d][1], d, grid, memo)
		return
	}
	if (*grid)[y][x] == "|" {
		(*memo)[y][x] = "#"
		if d == 1 || d == 3 {
			scanGrid(y+dir[d][0], x+dir[d][1], d, grid, memo)
		} else {
			scanGrid(y+dir[1][0], x+dir[1][1], 1, grid, memo)
			scanGrid(y+dir[3][0], x+dir[3][1], 3, grid, memo)
		}
		return
	}
	if (*grid)[y][x] == "-" {
		(*memo)[y][x] = "#"
		if d == 0 || d == 2 {
			scanGrid(y+dir[d][0], x+dir[d][1], d, grid, memo)
		} else {
			scanGrid(y+dir[0][0], x+dir[0][1], 0, grid, memo)
			scanGrid(y+dir[2][0], x+dir[2][1], 2, grid, memo)
		}
		return
	}
	if (*grid)[y][x] == "/" {
		(*memo)[y][x] = "#"
		if d == 0 {
			scanGrid(y+dir[1][0], x+dir[1][1], 1, grid, memo)
		} else if d == 1 {
			scanGrid(y+dir[0][0], x+dir[0][1], 0, grid, memo)
		} else if d == 2 {
			scanGrid(y+dir[3][0], x+dir[3][1], 3, grid, memo)
		} else {
			scanGrid(y+dir[2][0], x+dir[2][1], 2, grid, memo)
		}
		return
	}
	if (*grid)[y][x] == "\\" {
		(*memo)[y][x] = "#"
		if d == 0 {
			scanGrid(y+dir[3][0], x+dir[3][1], 3, grid, memo)
		} else if d == 1 {
			scanGrid(y+dir[2][0], x+dir[2][1], 2, grid, memo)
		} else if d == 2 {
			scanGrid(y+dir[1][0], x+dir[1][1], 1, grid, memo)
		} else {
			scanGrid(y+dir[0][0], x+dir[0][1], 0, grid, memo)
		}
		return
	}
	if (*grid)[y][x] != mark[d] {
		(*memo)[y][x] = "#"
		scanGrid(y+dir[d][0], x+dir[d][1], d, grid, memo)
	}
}

func getPoundNumber(memo [][]string) (tot int) {
	for y := 0; y < len(memo); y++ {
		for x := 0; x < len(memo[0]); x++ {
			if memo[y][x] == "#" {
				tot++
			}
		}
	}
	return tot
}
