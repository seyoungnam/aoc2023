package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fields, counts := loadFile("../input.txt")
	fmt.Printf("field : %v\n", fields)
	fmt.Printf("count : %v\n", counts)

	res := 0
	for i, field := range fields {
		// subRes := len(createTargetIntArray(field, counts[i]))
		subRes := 0
		count := copyCount(counts[i])
		for _, can := range deleteZero(createTargetIntArray(field, count)) {
			if checkTwoArrSame(can, count) {
				subRes++
			}
		}
		fmt.Printf("i=%v | subRes=%v\n", i, subRes)
		res += subRes
	}
	fmt.Printf("The answer is %v\n", res)

	// intArr := createTargetIntArray("??.??????#???.?#???.??????#???.?#???.??????#???.?#???.??????#???.?#???.??????#???.?#", []int{1, 1, 3, 1, 2, 1, 1, 3, 1, 2, 1, 1, 3, 1, 2, 1, 1, 3, 1, 2, 1, 1, 3, 1, 2})
	// fmt.Println(len(intArr))

	// intArr := createTargetIntArray(".??..??...?##.?.??..??...?##.?.??..??...?##.?.??..??...?##.?.??..??...?##.", []int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3})
	// fmt.Println(len(intArr))

}

func loadFile(fileName string) (field []string, count [][]int) {

	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, " ")
		field = append(field, row[0]+"?"+row[0]+"?"+row[0]+"?"+row[0]+"?"+row[0])
		// field = append(field, row[0])

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

func createTargetIntArray(field string, count []int) [][]int {
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
	prevAns := createTargetIntArray(field[1:], count)
	currAns := [][]int{}
	if string(field[0]) == "." {
		// fmt.Printf("field=%v | prevAns=%v\n", field, prevAns)
		for _, numArr := range prevAns {
			if numArr[0] != 0 {
				numArr = append([]int{0}, numArr...)
			}
			if len(numArr) == 1 && numArr[0] == 0 {
				currAns = append(currAns, numArr)
				continue
			}
			// fmt.Printf("numArr=%v | count[p]=%v\n", numArr, count[len(count)-(len(numArr)-1)])
			if len(numArr) <= len(count)+1 && numArr[1] == count[len(count)-(len(numArr)-1)] {
				currAns = append(currAns, numArr)
			}
		}
	} else if string(field[0]) == "#" {
		for _, numArr := range prevAns {
			numArr[0]++
			currAns = append(currAns, numArr)
		}
	} else { // string(field[0]) == "?"
		for _, numArr := range prevAns {
			// case "."
			curArr := copy(numArr)
			if curArr[0] != 0 {
				curArr = append([]int{0}, curArr...)
			}
			if len(curArr) == 1 && curArr[0] == 0 {
				currAns = append(currAns, curArr)
				// case "#"
				numArr[0]++
				currAns = append(currAns, numArr)
				continue
			}
			// fmt.Printf("curArr=%v | p=%v | count[p]=%v\n", curArr, len(count)-(len(curArr)-1), count[len(count)-(len(curArr)-1)])
			if len(curArr) <= len(count)+1 && curArr[1] == count[len(count)-(len(curArr)-1)] {
				currAns = append(currAns, curArr)
			}
			// case "#"
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

func copy(arr []int) []int {
	arr2 := []int{}
	for i := 0; i < len(arr); i++ {
		arr2 = append(arr2, arr[i])
	}
	return arr2
}

func copyCount(count []int) (res []int) {
	res = append(res, count...)
	res = append(res, count...)
	res = append(res, count...)
	res = append(res, count...)
	res = append(res, count...)
	return res
}
