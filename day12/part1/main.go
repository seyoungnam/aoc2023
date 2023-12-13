package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	field, count := loadFile("../example.txt")
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

	// res := 0
	// for i := 0; i < len(field); i++ {
	// 	subRes := getArrangeCount(field[i], sharp[i])
	// 	fmt.Println(i, subRes)
	// 	res += subRes
	// }
	// fmt.Println(res)

	row := []string{"???", "###"}
	cases := getCasesForRow(row)
	fmt.Println(cases)

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
			row = append(row, strings.Repeat("#", strings.Index("0123456789", countArr[i][j])))
		}
		sharpArr = append(sharpArr, row)
	}
	return sharpArr
}

func getCasesForEle(ele string) (cases [][]string) {
	eles := []string{}
	for _, v := range ele {
		if string(v) == "#" {
			for j := 0; j < len(eles); j++ {
				eles[j] += "#"
			}
		} else { // if v == ?
			curLen := len(eles)
			for j := 0; j < curLen; j++ {
				eles = append(eles, eles[j]+".")
				eles[j] += "#"
			}
			if len(eles) == 0 {
				eles = append(eles, ".", "#")
			}
		}
	}
	for _, ele := range eles {
		cases = append(cases, splitByDot(ele))
	}
	return cases
}

func getCasesForRow(row []string) (cases [][]string) {
	for i, ele := range row {
		if i == 0 {
			cases = getCasesForEle(ele)
		} else {
			for _, newCase := range getCasesForEle(ele) {
				currCases := cases
				newCases := [][]string{}
				for _, tg := range currCases {
					nc := append(tg, newCase...)
					newCases = append(newCases, nc)
				}
				cases = newCases
			}
		}
	}
	return cases
}

// func getArrangeCount(fieldRow, countRow []string) (count int) {
// 	totalCases := [][]string{}
// 	for _, ele := range fieldRow {

// 		if len(totalCases) == 0 {
// 			totalCases = getCases(ele)
// 		} else {
// 			currTotalCases := totalCases
// 			for _, cases := range getCases(ele) {
// 				newRow := []string{}
// 				for i := 0; i < len(currTotalCases); i++ {
// 					newRow = append(currTotalCases[i], cases...)
// 				}
// 				total
// 			}
// 			totalCases = currTotalCases
// 			fmt.Printf("cases: %v\n", totalCases)
// 		}
// 	}
// 	for _, c := range totalCases {
// 		if checkTwoArrSame(c, countRow) {
// 			count++
// 		}

// 	}
// 	fmt.Printf("cases: %v\n", totalCases)

// 	return count

// }

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
