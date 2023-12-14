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
	table := getCountTable(grid)
	res := getTotalFromTable(table)
	fmt.Printf("The answer is %v\n", res)

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
