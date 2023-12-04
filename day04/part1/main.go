package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	winNumSet, ownNumSet := getWinNumsAndOwnNumsFromFile("../input.txt")
	// fmt.Printf("winNumSet : %q | ownNumSet : %q\n", winNumSet, ownNumSet)
	res := 0
	for i := 0; i < len(winNumSet); i++ {
		res += getPoints(winNumSet[i], ownNumSet[i])
	}
	fmt.Printf("The total point is : %v", res)

}

func getWinNumsAndOwnNumsFromFile(fileName string) (winNumSet [][]string, ownNumSet [][]string) {
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		colon := strings.Index(line, ":")
		bar := strings.Index(line, "|")
		winNums, ownNums := getWinNumsAndOwnNumsFromLine(line, colon, bar)

		winNumSet = append(winNumSet, winNums)
		ownNumSet = append(ownNumSet, ownNums)
	}
	return winNumSet, ownNumSet
}

func getWinNumsAndOwnNumsFromLine(line string, colonIdx int, barIdx int) (winNums []string, ownNums []string) {
	winNumStr, ownNumStr := line[colonIdx+1:barIdx], line[barIdx+1:]
	winNumArr, ownNumArr := strings.Split(winNumStr, " "), strings.Split(ownNumStr, " ")
	winNums, ownNums = getNonEmptyArr(winNumArr), getNonEmptyArr(ownNumArr)
	return winNums, ownNums
}

func getNonEmptyArr(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" && s != " " && s != "   " {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func getPoints(winNums []string, ownNums []string) (pt int) {
	winNumsMap := map[string]int{}

	for _, winNum := range winNums {
		winNumsMap[winNum] = 2
	}

	for _, ownNum := range ownNums {
		if winNumsMap[ownNum] != 0 {
			if pt == 0 {
				pt = 1
			} else {
				pt *= 2
			}
		}
	}
	fmt.Printf("winNums : %q | ownNums : %q | point : %v\n", winNums, ownNums, pt)
	return pt
}
