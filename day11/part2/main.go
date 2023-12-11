package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	image := loadFile("../input.txt")
	Y, X := sharpLocations(image)

	pts := collectPoints(image)
	fmt.Printf("points : %v\n", pts)

	res := sumSortestPaths(pts, Y, X, 1000000)
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

func sharpLocations(grid [][]string) (Y, X map[int]bool) {
	Y = map[int]bool{}
	X = map[int]bool{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == "#" {
				Y[y] = true
				X[x] = true
			}
		}
	}
	return Y, X
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

func sumSortestPaths(pts [][2]int, Y, X map[int]bool, times int) (res int) {
	for i := 0; i < len(pts); i++ {
		for j := i + 1; j < len(pts); j++ {
			ly, ry := sort(pts[i][0], pts[j][0])
			ny := countNonSharpNum(ly, ry, Y)
			res += ry - ly + ny*(times-1)

			lx, rx := sort(pts[i][1], pts[j][1])
			nx := countNonSharpNum(lx, rx, X)
			res += rx - lx + nx*(times-1)
		}
	}
	return res
}

func sort(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func countNonSharpNum(l, r int, S map[int]bool) (cnt int) {
	for i := l + 1; i < r; i++ {
		if !S[i] {
			cnt++
		}
	}
	return cnt
}
