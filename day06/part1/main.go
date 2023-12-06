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
	timeArr := getArrayFromFile(lines, "Time")
	distArr := getArrayFromFile(lines, "Distance")
	fmt.Println(timeArr, distArr)

	res := 1
	for i := 0; i < len(timeArr); i++ {
		cnt := getCountToBeat(timeArr[i], distArr[i])
		fmt.Printf("Time=%v | Dist=%v | Count=%v\n", timeArr[i], distArr[i], cnt)
		res *= cnt
	}
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

func getArrayFromFile(lines []string, str string) (nums []int) {
	for _, line := range lines {
		if strings.HasPrefix(line, str) {
			i := strings.Index(line, ":")
			numsStr := strings.Split(line[i+1:], " ")
			nums = removeEmptyStrAndTurnStrToInt(numsStr)
			break
		}
	}
	return nums
}

func removeEmptyStrAndTurnStrToInt(strings []string) (nums []int) {
	for _, s := range strings {
		if s != "" && s != " " && s != "   " {
			num, _ := strconv.Atoi(s)
			nums = append(nums, num)
		}
	}
	return nums
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
