package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var nextHop = map[string][][2]int{
	"S": {{0, -1}, {1, 0}, {0, 1}, {-1, 0}},
	"|": {{1, 0}, {-1, 0}},
	"-": {{0, -1}, {0, 1}},
	"L": {{0, 1}, {-1, 0}},
	"J": {{0, -1}, {-1, 0}},
	"7": {{0, -1}, {1, 0}},
	"F": {{1, 0}, {0, 1}},
}

func main() {
	res := 0
	tiles := loadFile("../input.txt")
	for _, line := range tiles {
		fmt.Println(line)
	}
	d := getSCoordinate(tiles)
	fmt.Printf("S coordinate: %v\n", d)

	// get grid
	grid := makeGrid(tiles)

	for y := 0; y < len(tiles); y++ {
		for x := 0; x < len(tiles[0]); x++ {
			if tiles[y][x] == "S" {
				res = markGrid(tiles, grid, y, x)
			}
		}
	}

	fmt.Printf("The answer is : %v\n", res)

	// get enclosed area
	area := getEnclosedArea(tiles, grid)
	fmt.Printf("The area is : %v\n", area)

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

func getSCoordinate(grid [][]string) (d [2]int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == "S" {
				d[0], d[1] = y, x
				break
			}
		}
	}
	return d
}

func makeGrid(lines [][]string) (grid [][]int) {
	for y := 0; y < len(lines); y++ {
		row := []int{}
		for x := 0; x < len(lines[0]); x++ {
			row = append(row, math.MaxInt)
		}
		grid = append(grid, row)
	}
	return grid
}

func markGrid(tiles [][]string, grid [][]int, sy, sx int) (res int) {
	// mark starting point
	grid[sy][sx] = 0

	// mark visited
	visited := map[string]bool{}
	visited[h([2]int{sy, sx})] = true

	arr := [][2]int{}
	arr = append(arr, [2]int{sy, sx})

	for len(arr) > 0 {
		length := len(arr)
		// fmt.Println(length)
		for i := 0; i < length; i++ {
			p := arr[0]
			arr = arr[1:]
			py, px := p[0], p[1]
			// fmt.Printf("p=%v\n", p)
			ds, ok := nextHop[tiles[py][px]]
			if !ok {
				continue
			}
			// fmt.Println(ds)
			for _, d := range ds {
				cy, cx := py+d[0], px+d[1]
				c := [2]int{cy, cx}
				// fmt.Printf("%v\n", c)
				// out of range, pass
				if cy < 0 || cy >= len(tiles) || cx < 0 || cx >= len(tiles[0]) {
					continue
				}
				// already visited, pass
				if visited[h(c)] {
					continue
				}

				// check if curr tile is connected to past tile
				connected := false
				for _, d := range nextHop[tiles[cy][cx]] {
					if cy+d[0] == py && cx+d[1] == px {
						connected = true
					}
				}
				if !connected {
					continue
				}

				// all filters are passed, the next point is valid
				if grid[cy][cx] > grid[py][px]+1 {
					grid[cy][cx] = grid[py][px] + 1
					if res < grid[cy][cx] {
						res = grid[cy][cx]
					}
				}
				arr = append(arr, [2]int{cy, cx})
				visited[h(c)] = true
			}
		}
	}
	return res
}

func h(d [2]int) string { return fmt.Sprintf("%q", d) }

func getEnclosedArea(tiles [][]string, grid [][]int) (area int) {
	// row
	for y := 0; y < len(grid); y++ {
		crossNum := 0
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] != math.MaxInt {
				if strings.Contains("|LJS", tiles[y][x]) {
					crossNum++
				}
			} else {
				if crossNum%2 == 1 {
					area++
				}
			}
		}
	}

	return area
}

func copy(grid [][]int) (copied [][]int) {
	for y := 0; y < len(grid); y++ {
		row := []int{}
		for x := 0; x < len(grid[0]); x++ {
			row = append(row, grid[y][x])
		}
		copied = append(copied, row)
	}
	return copied
}
