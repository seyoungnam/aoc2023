package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	field, count := loadFile("../input.txt")
	fmt.Println("field below:")
	for i, row := range field {
		fmt.Println(i, row)
	}
	// fmt.Printf("count = %q\n", count)

	sharp := changeNumToSharp(count)
	fmt.Println("sharp below:")
	for i, row := range sharp {
		fmt.Println(i, row)
	}

	res := 0
	for i := 0; i < len(field); i++ {
		subRes := getCount(field[i], sharp[i])
		// fmt.Println(i, subRes)
		res += subRes
	}
	fmt.Printf("The answer is %v\n", res)

	// eleCases := getCasesForEle("???")
	// fmt.Println(eleCases)

	// row := []string{"???", "###"}
	// cases := getCasesForRow(row)
	// fmt.Println(cases)
	// for _, c := range cases {
	// 	fmt.Println(breakRowByDot(c))
	// }

}

func loadFile(fileName string) (field [][]string, count [][]string) {

	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, " ")
		rowNew := []string{}
		for _, ele := range strings.Split(row[0], ".") {
			if ele != "" {
				rowNew = append(rowNew, ele)
			}
		}
		field = append(field, rowNew)
		count = append(count, strings.Split(row[1], ","))
	}
	return field, count
}

func splitByDot(ele string) (eles []string) {
	for _, v := range strings.Split(ele, ".") {
		if v != "" {
			eles = append(eles, v)
		}
	}
	return eles
}

func changeNumToSharp(countArr [][]string) (sharpArr [][]string) {
	for i := 0; i < len(countArr); i++ {
		row := []string{}
		for j := 0; j < len(countArr[i]); j++ {
			repeat, _ := strconv.Atoi(countArr[i][j])
			row = append(row, strings.Repeat("#", repeat))
		}
		sharpArr = append(sharpArr, row)
	}
	return sharpArr
}

func getCasesForEle(ele string) []string {
	if len(ele) == 0 {
		return nil
	}
	if len(ele) == 1 {
		cases := []string{"#"}
		if ele == "?" {
			cases = append(cases, ".")
		}
		return cases
	}
	prevCases := getCasesForEle(ele[1:])
	currCases := []string{}
	for _, c := range prevCases {
		currCases = append(currCases, "#"+c)
		if string(ele[0]) == "?" {
			currCases = append(currCases, "."+c)
		}
	}
	return currCases
}

func getCasesForRow(row []string) [][]string {
	if len(row) == 0 {
		return nil
	}
	if len(row) == 1 {
		cases := [][]string{}
		for _, c := range getCasesForEle(row[0]) {
			ele := []string{c}
			cases = append(cases, ele)
		}
		return cases
	}
	prevCases := getCasesForRow(row[1:])
	currCases := [][]string{}
	for _, c := range getCasesForEle(row[0]) {
		for _, cArr := range prevCases {
			currCases = append(currCases, append([]string{c}, cArr...))
		}
	}
	return currCases
}

func breakRowByDot(row []string) []string {
	var newRow []string
	for _, ele := range row {
		newRow = append(newRow, splitByDot(ele)...)
	}
	return newRow
}

func getCount(fieldRow, targetRow []string) (count int) {
	cases := getCasesForRow(fieldRow)
	for _, c := range cases {
		row := breakRowByDot(c)
		if checkTwoArrSame(row, targetRow) {
			count++
		}
	}
	return count
}

func checkTwoArrSame(arr1, arr2 []string) bool {
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
