package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	res := 0

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		res += getTwoDigitNumber(line)
	}

	fmt.Println("The sum is : ", res)

}

func getTwoDigitNumber(line string) (num int) {
	first := -1
	last := -1
	l := 0
	r := len(line) - 1

	for l <= r {
		if first >= 0 && last >= 0 {
			break
		} else if l == r {
			first = int(line[l])
			last = int(line[r])
			break
		} else if line[l] < 48 || line[l] > 57 {
			l++
		} else if line[r] < 48 || line[r] > 57 {
			r--
		} else if first < 0 && line[l] >= 48 && line[l] <= 57 {
			first = int(line[l])
		} else if last < 0 && line[r] >= 48 && line[r] <= 57 {
			last = int(line[r])
		}
	}

	num = (first-48)*10 + (last - 48)
	return num
}
