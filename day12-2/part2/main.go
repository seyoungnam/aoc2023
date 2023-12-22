package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type key struct {
	field string
	nums  []int
}

func main() {
	field, count := loadFile("../input.txt")
	fmt.Println("field below:")
	for i, row := range field {
		fmt.Println(i, row)
	}

	fmt.Println("count below:")
	for i, row := range count {
		fmt.Println(i, row)
	}

	cache := map[*key]int{}
	res := 0
	for i := 0; i < len(field); i++ {
		subRes := getCount(field[i], count[i], cache)
		fmt.Printf("subRes : %v\n", subRes)
		res += subRes
	}
	fmt.Printf("The answer is %v\n", res)

}

func loadFile(fileName string) (field []string, count [][]int) {

	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, " ")
		field = append(field, row[0]+"?"+row[0]+"?"+row[0]+"?"+row[0]+"?"+row[0])
		cntArr := []int{}
		for _, strNum := range strings.Split(row[1], ",") {
			num, _ := strconv.Atoi(strNum)
			cntArr = append(cntArr, num)
		}
		cntArr5x := []int{}
		for i := 0; i < 5; i++ {
			cntArr5x = append(cntArr5x, cntArr...)
		}
		count = append(count, cntArr5x)
	}
	return field, count
}

func getCount(field string, nums []int, cache map[*key]int) int {
	if field == "" {
		if len(nums) == 0 {
			return 1
		}
		return 0
	}

	if len(nums) == 0 {
		if strings.Contains(field, "#") {
			return 0
		}
		return 1
	}

	key := &key{
		field: field,
		nums:  nums,
	}

	if val, ok := cache[key]; ok {
		return val
	}

	cnt := 0

	if strings.Contains(".?", string(field[0])) {
		cnt += getCount(field[1:], nums, cache)
	}

	if strings.Contains("#?", string(field[0])) {
		if nums[0] <= len(field) && !strings.Contains(field[:nums[0]], ".") {
			if nums[0] == len(field) {
				cnt += getCount(field[nums[0]:], nums[1:], cache)
			} else if string(field[nums[0]]) != "#" {
				cnt += getCount(field[nums[0]+1:], nums[1:], cache)
			}
		}
	}
	cache[key] = cnt
	return cnt
}
