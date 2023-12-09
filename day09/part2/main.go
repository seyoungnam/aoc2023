package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := loadFile("../input.txt")
	res := 0
	for _, line := range lines {
		arr := makeArr(line)
		// fmt.Println(arr)
		ev := getExtrapolValBackwards(arr)
		fmt.Println(ev)
		res += ev
	}
	fmt.Printf("The sum is : %v\n", res)

}

func loadFile(fileName string) (lines []string) {
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func makeArr(line string) (arr []int) {
	strArr := strings.Split(line, " ")
	for _, num := range strArr {
		numInt, _ := strconv.Atoi(num)
		arr = append(arr, numInt)
	}
	return arr
}

func getExtrapolVal(nums []int) (val int) {
	val += nums[len(nums)-1]
	curr := nums

	for {
		cnt := 0
		next := []int{}
		for i := 0; i < len(curr)-1; i++ {
			d := curr[i+1] - curr[i]
			next = append(next, d)
			if d == 0 {
				cnt++
			}
		}
		if cnt == len(next) {
			break
		}
		val += next[len(next)-1]
		fmt.Printf("next : %v\n", next)
		curr = next
	}

	return val
}

func getExtrapolValBackwards(curr []int) (val int) {
	cnt := 0
	next := []int{}
	for i := 0; i < len(curr)-1; i++ {
		if curr[i] == 0 {
			cnt++
		}
		d := curr[i+1] - curr[i]
		next = append(next, d)
	}

	if curr[len(curr)-1] == 0 {
		cnt++
	}

	if cnt == len(curr) {
		return 0
	}

	return curr[0] - getExtrapolValBackwards(next)
}
