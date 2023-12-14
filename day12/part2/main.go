package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fields, counts := loadFile("../example.txt")
	fmt.Printf("field : %v\n", fields)
	fmt.Printf("count : %v\n", counts)

	res := 0
	for i, field := range fields {
		subRes := 0
		for _, can := range deleteZero(createTargetIntArray(field)) {
			if checkTwoArrSame(can, counts[i]) {
				subRes++
			}
		}
		fmt.Printf("i=%v | subRes=%v\n", i, subRes)
		res += subRes
	}
	fmt.Printf("The answer is %v\n", res)

	intArr := createTargetIntArray(".??..??...?##.")
	fmt.Println(intArr)

}

func loadFile(fileName string) (field []string, count [][]int) {

	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, " ")
		// field = append(field, row[0]+"?"+row[0]+"?"+row[0]+"?"+row[0]+"?"+row[0])
		field = append(field, row[0])

		nums := []int{}
		for _, s := range strings.Split(row[1], ",") {
			num, _ := strconv.Atoi(s)
			nums = append(nums, num)
		}
		count = append(count, nums)

	}
	return field, count
}

func checkTwoArrSame(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func createTargetIntArray(field string) [][]int {
	if len(field) == 1 {
		if field == "." {
			ans := [][]int{}
			ans = append(ans, []int{0})
			return ans
		}
		if field == "#" {
			ans := [][]int{}
			ans = append(ans, []int{1})
			return ans
		}
		if field == "?" {
			ans := [][]int{}
			ans = append(ans, []int{0}, []int{1})
			return ans
		}
	}
	prevAns := createTargetIntArray(field[1:])
	currAns := [][]int{}
	if string(field[0]) == "." {
		for _, numArr := range prevAns {
			numArr = append([]int{0}, numArr...)
			currAns = append(currAns, numArr)
		}
	} else if string(field[0]) == "#" {
		for _, numArr := range prevAns {
			numArr[0]++
			currAns = append(currAns, numArr)
		}
	} else { // string(field[0]) == "?"
		for _, numArr := range prevAns {
			curArr := numArr
			curArr = append([]int{0}, curArr...)
			currAns = append(currAns, curArr)
			numArr[0]++
			currAns = append(currAns, numArr)
		}
	}
	// // clean up zeros
	// nonZeroCurrAns := [][]int{}
	// for _, arr := range currAns {
	// 	nonZeroArr := []int{}
	// 	for _, num := range arr {
	// 		if num != 0 {
	// 			nonZeroArr = append(nonZeroArr, num)
	// 		}
	// 	}
	// 	nonZeroCurrAns = append(nonZeroCurrAns, nonZeroArr)
	// }
	// return nonZeroCurrAns
	// fmt.Printf("currAns=%v\n", currAns)
	return currAns
}

func deleteZero(numArr [][]int) (newArr [][]int) {
	for _, row := range numArr {
		newRow := []int{}
		for _, ele := range row {
			if ele != 0 {
				newRow = append(newRow, ele)
			}
		}
		newArr = append(newArr, newRow)
	}
	return newArr
}
