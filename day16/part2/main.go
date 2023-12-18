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

	res := 0
	// top row
	for x := 0; x < len(grid[0]); x++ {
		gridC := copyGrid(grid)
		memoC := copyGrid(memo)
		newRes := scanGrid(0, x, 1, &gridC, &memoC)
		fmt.Println(0, x, newRes)
		if newRes > res {
			res = newRes
		}
	}

	// bottom row
	for x := 0; x < len(grid[0]); x++ {
		gridC := copyGrid(grid)
		memoC := copyGrid(memo)
		newRes := scanGrid(len(grid)-1, x, 3, &gridC, &memoC)
		fmt.Println(len(grid)-1, x, newRes)
		if newRes > res {
			res = newRes
		}
	}

	// left column
	for y := 0; y < len(grid[0]); y++ {
		gridC := copyGrid(grid)
		memoC := copyGrid(memo)
		newRes := scanGrid(y, 0, 2, &gridC, &memoC)
		fmt.Println(y, 0, newRes)
		if newRes > res {
			res = newRes
		}
	}

	// right column
	for y := 0; y < len(grid[0]); y++ {
		gridC := copyGrid(grid)
		memoC := copyGrid(memo)
		newRes := scanGrid(y, len(grid[0])-1, 0, &gridC, &memoC)
		fmt.Println(y, len(grid[0])-1, newRes)
		if newRes > res {
			res = newRes
		}
	}

	fmt.Println(res)
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

func scanGrid(y, x, d int, grid *[][]string, memo *[][]string) (step int) {
	// fmt.Println(y, x, y+dir[d][0], x+dir[d][1])
	if y < 0 || y >= len(*grid) || x < 0 || x >= len((*grid)[0]) {
		return 0
	}
	if (*grid)[y][x] == "." {
		if (*memo)[y][x] == "." {
			(*memo)[y][x] = "#"
			step++
		}
		(*grid)[y][x] = mark[d]
		step += scanGrid(y+dir[d][0], x+dir[d][1], d, grid, memo)
		return step
	}
	if (*grid)[y][x] == "|" {
		if (*memo)[y][x] == "." {
			(*memo)[y][x] = "#"
			step++
		}
		if d == 1 || d == 3 {
			step += scanGrid(y+dir[d][0], x+dir[d][1], d, grid, memo)
		} else {
			step += scanGrid(y+dir[1][0], x+dir[1][1], 1, grid, memo)
			step += scanGrid(y+dir[3][0], x+dir[3][1], 3, grid, memo)
		}
		return step
	}
	if (*grid)[y][x] == "-" {
		if (*memo)[y][x] == "." {
			(*memo)[y][x] = "#"
			step++
		}
		if d == 0 || d == 2 {
			step += scanGrid(y+dir[d][0], x+dir[d][1], d, grid, memo)
		} else {
			step += scanGrid(y+dir[0][0], x+dir[0][1], 0, grid, memo)
			step += scanGrid(y+dir[2][0], x+dir[2][1], 2, grid, memo)
		}
		return step
	}
	if (*grid)[y][x] == "/" {
		if (*memo)[y][x] == "." {
			(*memo)[y][x] = "#"
			step++
		}
		if d == 0 {
			step += scanGrid(y+dir[1][0], x+dir[1][1], 1, grid, memo)
		} else if d == 1 {
			step += scanGrid(y+dir[0][0], x+dir[0][1], 0, grid, memo)
		} else if d == 2 {
			step += scanGrid(y+dir[3][0], x+dir[3][1], 3, grid, memo)
		} else {
			step += scanGrid(y+dir[2][0], x+dir[2][1], 2, grid, memo)
		}
		return step
	}
	if (*grid)[y][x] == "\\" {
		if (*memo)[y][x] == "." {
			(*memo)[y][x] = "#"
			step++
		}
		if d == 0 {
			step += scanGrid(y+dir[3][0], x+dir[3][1], 3, grid, memo)
		} else if d == 1 {
			step += scanGrid(y+dir[2][0], x+dir[2][1], 2, grid, memo)
		} else if d == 2 {
			step += scanGrid(y+dir[1][0], x+dir[1][1], 1, grid, memo)
		} else {
			step += scanGrid(y+dir[0][0], x+dir[0][1], 0, grid, memo)
		}
		return step
	}
	if (*grid)[y][x] != mark[d] {
		step += scanGrid(y+dir[d][0], x+dir[d][1], d, grid, memo)
	}
	return step
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

func copyGrid(src [][]string) [][]string {
	duplicate := make([][]string, len(src))
	for i := range src {
		duplicate[i] = make([]string, len(src[i]))
		copy(duplicate[i], src[i])
	}
	return duplicate
}
