package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := loadFile("../input.txt")
	seeds := getSeedsFromFile(lines)
	// fmt.Println(seeds)
	begIdx := getMapBegIdx(lines)
	// fmt.Println(begIdx)
	mapSets := parsingMapNums(lines, begIdx)
	fmt.Println(mapSets)

	checkPoint := &[7]int{}
	for i := 0; i < 7; i++ {
		checkPoint[i] = math.MaxInt
	}

	sortedSeeds := sortingSeeds(seeds)
	fmt.Printf("sortedSeeds = %v\n", sortedSeeds)
	seedRange := getSeedRange(sortedSeeds)
	fmt.Printf("seedRange   = %v\n", seedRange)

	for _, s := range seedRange {
		beg, end := s[0], s[1]
		for i := beg; i < end; i++ {
			seedToLocation(mapSets, i, checkPoint)
		}
	}

	fmt.Println(checkPoint[len(checkPoint)-1])

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

func getSeedsFromFile(lines []string) (seeds []int) {
	for _, line := range lines {
		if strings.HasPrefix(line, "seeds:") {
			i := strings.Index(line, ":")
			numsStr := strings.Split(line[i+1:], " ")
			seeds = removeEmptyStrAndTurnStrToInt(numsStr)
			break
		}
	}
	return seeds
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

func getMapBegIdx(lines []string) (begIdx [8]int) {
	titleIdx := map[string]int{
		"seed-to-soil map:":            0,
		"soil-to-fertilizer map:":      1,
		"fertilizer-to-water map:":     2,
		"water-to-light map:":          3,
		"light-to-temperature map:":    4,
		"temperature-to-humidity map:": 5,
		"humidity-to-location map:":    6,
	}

	for i, line := range lines {
		idx, ok := titleIdx[line]
		if ok {
			begIdx[idx] = i + 1
		}
	}
	begIdx[7] = len(lines) + 2
	return begIdx
}

func parsingMapNums(lines []string, begIdx [8]int) (mapSets [7][][]int) {
	for i := 0; i < 7; i++ {
		beg, end := begIdx[i], begIdx[i+1]-2
		mapSet := [][]int{}
		for j := beg; j < end; j++ {
			mapSet = append(mapSet, parsingLine(lines[j]))
		}
		mapSets[i] = mapSet
	}
	return mapSets
}

func parsingLine(line string) (numsInt []int) {
	numsStr := strings.Split(line, " ")
	numsInt = removeEmptyStrAndTurnStrToInt(numsStr)
	return numsInt
}

func seedToLocation(mapSets [7][][]int, seed int, checkPoint *[7]int) {
	for i := 0; i < 7; i++ {
		mapSet := mapSets[i]

		for _, vals := range mapSet {
			// fmt.Printf("i=%v | seed=%v | vals=%v\n", i, seed, vals)
			if seed >= vals[1] && seed < vals[1]+vals[2] {
				seed += +vals[0] - vals[1]
				break
			}
		}
		// implement checkpoint
		if seed == checkPoint[i] {
			return
		} else if seed < checkPoint[i] {
			checkPoint[i] = seed
		}
		// fmt.Println(checkPoint)
	}
}

func minFind(arr []int) int {
	min := arr[0]
	for _, num1 := range arr {
		if num1 < min {
			min = num1
		}
	}
	return min
}

func sortingSeeds(seeds []int) (sortedSeeds [][2]int) {
	unsortedSeeds := [][2]int{}
	for i := 0; i < len(seeds); i++ {
		if i%2 != 0 {
			continue
		}
		begToEnd := [2]int{seeds[i], seeds[i] + seeds[i+1]}
		unsortedSeeds = append(unsortedSeeds, begToEnd)
	}
	return sortArray(unsortedSeeds)
}

func sortArray(arr [][2]int) [][2]int {
	for i := 0; i <= len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j][0] > arr[j+1][0] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func getSeedRange(arr [][2]int) (r [][2]int) {
	r = append(r, arr[0])
	for i, v := range arr {
		if i == 0 {
			continue
		}
		if v[0] > r[len(r)-1][1] {
			r = append(r, v)
		} else {
			r[len(r)-1][1] = v[1]
		}
	}
	return r
}
