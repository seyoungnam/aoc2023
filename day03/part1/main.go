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
	markingGrid := markingInGrid(inputGrid)
	res := 0
	for y, gridRow := range markingGrid {
		for x, targetFound := range gridRow {
			// fmt.Printf("y=%d | f=%d | tf=%t\n", y, x, targetFound)
			if targetFound {
				fmt.Printf("y=%d | x=%d\n", y, x)
				fmt.Printf("inputGrid[y]=%q\n", inputGrid[y])
				num := getNumber(y, x, inputGrid[y], markingGrid)
				fmt.Println(num)
				res += num
			}
		}
	}
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

func markingInGrid(inputGrid [][]string) [][]bool {
	markingGrid := make([][]bool, len(inputGrid))
	prev := make([]bool, len(inputGrid[0]))
	curr := make([]bool, len(inputGrid[0]))
	for y := 0; y < len(inputGrid); y++ {
		next := make([]bool, len(inputGrid[0]))
		for x := 0; x < len(inputGrid[y]); x++ {
			if !strings.Contains("0123456789.", inputGrid[y][x]) {
				// fmt.Printf("y=%d | x=%d | element=%q | bool=%t\n", y, x, inputGrid[y][x], !strings.Contains("0123456789.", inputGrid[y][x]))
				// left col
				if x != 0 {
					prev[x-1] = true
					curr[x-1] = true
					next[x-1] = true
				}
				// current col
				prev[x] = true
				next[x] = true
				// right col
				if x != len(inputGrid[y])-1 {
					prev[x+1] = true
					curr[x+1] = true
					next[x+1] = true
				}
			}
		}
		if y != 0 {
			markingGrid[y-1] = prev
		}
		markingGrid[y] = curr
		if y != len(inputGrid)-1 {
			markingGrid[y+1] = next
		}
		prev = curr
		curr = next
	}
	return markingGrid
}

func getNumber(y int, x int, inputRow []string, markingGrid [][]bool) int {
	if !strings.Contains("0123456789", inputRow[x]) {
		return 0
	}
	for x >= 0 && strings.Contains("0123456789", inputRow[x]) {
		x--
	}
	x++
	num := 0
	for x < len(inputRow) && strings.Contains("0123456789", inputRow[x]) {
		markingGrid[y][x] = false
		digit, err := strconv.Atoi(inputRow[x])
		if err != nil {
			panic(err)
		}
		num *= 10
		num += digit
		x++
	}
	return num

}
