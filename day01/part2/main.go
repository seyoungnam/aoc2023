package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	letterMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	res := 0

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		res += getCalibrationValue(turnStringToDigitSlice(line, letterMap))
	}

	fmt.Println("The sum is : ", res)
}

func turnStringToDigitSlice(line string, letterMap map[string]int) []int {
	res := make([]int, len(line))
	// scan digit and store in slice
	for i, v := range line {
		if int(v) >= 48 && int(v) <= 57 {
			res[i] = int(v) - 48
		}
	}
	// scan letters
	for letter := range letterMap {
		// get indexes where letter is located in each line
		idxSlice := getSubstringIndex(line, letter, 0)
		// insert digits to res
		for _, i := range idxSlice {
			res[i] = letterMap[letter]
		}
	}
	return res
}

func getSubstringIndex(line string, substr string, startIdx int) []int {
	i := strings.Index(line, substr)
	// fmt.Printf("%d | %d | %s | %s", startIdx, i, line, substr)
	if len(line) < len(substr) || i == -1 {
		return nil
	}
	return append([]int{startIdx + i}, getSubstringIndex(line[i+len(substr):], substr, startIdx+i+len(substr))...)
}

func getCalibrationValue(digits []int) (val int) {
	l, r := 0, len(digits)-1
	first, last := -1, -1

	for l <= r {
		// fmt.Println(l, r, first, last)
		if first > 0 && last > 0 {
			break
		} else if l == r {
			first, last = digits[l], digits[r]
			break
		} else if digits[l] == 0 {
			l++
		} else if digits[r] == 0 {
			r--
		} else if first < 0 && digits[l] > 0 {
			first = digits[l]
		} else if last < 0 && digits[r] > 0 {
			last = digits[r]
		}
	}
	val = first*10 + last
	return val
}
