package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type point struct {
	Y int
	X int
}

func main() {
	image := loadFile("../input.txt")
	expended := expandGalaxies(image)
	fmt.Println("expanded image below:")
	for _, row := range expended {
		fmt.Println(row)
	}

	pts := collectPoints(expended)
	fmt.Printf("points : %v\n", pts)

	res := sumSortestPaths(pts)
	fmt.Printf("The sum of shortest paths : %v\n", res)

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

func expandGalaxies(grid [][]string) (expended [][]string) {
	X := map[int]bool{}
	Y := map[int]bool{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == "#" {
				Y[y] = true
				X[x] = true
			}
		}
	}

	for y := 0; y < len(grid); y++ {
		row := []string{}
		for x := 0; x < len(grid[0]); x++ {
			if X[x] {
				row = append(row, grid[y][x])
			} else {
				row = append(row, grid[y][x], grid[y][x])
			}
		}
		if Y[y] {
			expended = append(expended, row)
		} else {
			expended = append(expended, row, row)
		}
	}
	return expended
}

func collectPoints(grid [][]string) (pts [][2]int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == "#" {
				pts = append(pts, [2]int{y, x})
			}
		}
	}
	return pts
}

func sumSortestPaths(pts [][2]int) (res int) {
	for i := 0; i < len(pts); i++ {
		for j := i + 1; j < len(pts); j++ {
			res += abs(pts[i][0]-pts[j][0]) + abs(pts[i][1]-pts[j][1])
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
