package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := loadFile("../input.txt")
	time := getArrayFromFile(lines, "Time")
	dist := getArrayFromFile(lines, "Distance")
	// fmt.Println(time, dist)

	res := 0
	cnt := getCountToBeat(time, dist)
	fmt.Printf("Time=%v | Dist=%v | Count=%v\n", time, dist, cnt)
	res += cnt
	fmt.Printf("The answer is : %v\n", res)

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

func getArrayFromFile(lines []string, str string) (num int) {
	for _, line := range lines {
		if strings.HasPrefix(line, str) {
			i := strings.Index(line, ":")
			num = getIntFromLine(line, i)
			break
		}
	}
	return num
}

func getIntFromLine(line string, begIdx int) (num int) {
	numArr := []int{}
	for i := begIdx + 1; i < len(line); i++ {
		if int(line[i]) >= 48 && int(line[i]) < 58 {
			numArr = append(numArr, int(line[i])-48)
		}
	}
	for _, val := range numArr {
		num *= 10
		num += val
	}
	return num
}

func getCountToBeat(time int, dist int) (cnt int) {
	speed := time / 2
	runTime := time - speed
	if speed*runTime <= dist {
		return cnt
	}
	cnt++
	// reduce speed
	cnt += getCountToBeatWithDirection(speed-1, runTime+1, dist, true)
	// increase speed
	cnt += getCountToBeatWithDirection(speed+1, runTime-1, dist, false)
	return cnt
}

func getCountToBeatWithDirection(speed int, runTime int, dist int, downSpeed bool) (cnt int) {
	for speed*runTime > dist {
		cnt++
		if downSpeed {
			speed--
			runTime++
		} else {
			speed++
			runTime--
		}
	}
	return cnt
}
