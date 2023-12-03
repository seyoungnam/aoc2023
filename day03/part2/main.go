package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputGrid := linesInFile("../input.txt")
	coordinates := getStarCoordinates(inputGrid)
	fmt.Printf("star coordinates: %d\n", coordinates)
	numSetArr := getNumbersAroundStar(coordinates, inputGrid)
	fmt.Printf("nums array: %d\n", numSetArr)
	res := multiplyAndSum(numSetArr)
	fmt.Printf("The answer is : %d", res)

}

func linesInFile(fileName string) [][]string {
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)
	result := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, strings.Split(line, ""))
	}
	return result
}

// get coordinates where star is located [[1,3], [4,3], [8,5]]
func getStarCoordinates(grid [][]string) (stars [][]int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == "*" {
				star := []int{y, x}
				stars = append(stars, star)
			}
		}
	}
	return stars
}

func getNumbersAroundStar(stars [][]int, grid [][]string) (numSetArr [][]int) {
	dydx := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, coordinate := range stars {
		numSet := []int{}
		py, px := coordinate[0], coordinate[1]
		for _, delta := range dydx {
			y, x := py+delta[0], px+delta[1]
			if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
				break
			}
			if strings.Contains("0123456789", grid[y][x]) {
				num := turnStringsToNum(y, x, grid)
				if len(numSet) == 0 || num != numSet[len(numSet)-1] {
					numSet = append(numSet, num)
				}
			}
		}
		numSetArr = append(numSetArr, numSet)
	}
	return numSetArr
}

func turnStringsToNum(y int, x int, grid [][]string) (num int) {
	for x >= 0 && strings.Contains("0123456789", grid[y][x]) {
		x--
	}
	x++
	for x < len(grid[y]) && strings.Contains("0123456789", grid[y][x]) {
		digit, err := strconv.Atoi(grid[y][x])
		if err != nil {
			panic(err)
		}
		num *= 10
		num += digit
		x++
	}
	return num

}

func multiplyAndSum(numSetArr [][]int) (res int) {
	for _, numSet := range numSetArr {
		if len(numSet) < 2 {
			continue
		}
		multiplied := 1
		for _, num := range numSet {
			multiplied *= num
		}
		res += multiplied
	}
	return res
}
