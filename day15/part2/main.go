package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr := loadFile("../input.txt")
	fmt.Println(arr)

	res := 0
	for _, str := range arr {
		hash := getHash(str)
		// fmt.Println(str, hash)
		res += hash
	}
	fmt.Println(res)
	boxMap := makeMap(arr)
	fmt.Printf("boxMap=%v\n", boxMap)
	power := getFocusingPower(boxMap)
	fmt.Printf("The answer is %v\n", power)
}

func loadFile(fileName string) (lines []string) {
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Split(line, ",")
		lines = append(lines, lineArr...)
	}
	return lines
}

func getHash(str string) (hash int) {
	for i := 0; i < len(str); i++ {
		hash += int(str[i])
		hash *= 17
		hash %= 256
	}
	return hash
}

func makeMap(strArr []string) (boxMap map[int][][2]string) {
	boxMap = make(map[int][][2]string)
	for _, str := range strArr {
		i := strings.Index(str, "=")
		if i != -1 {
			label, focal := str[:i], str[i+1:]
			k := getHash(label)
			oldLensArr, ok := boxMap[k]
			if !ok {
				boxMap[k] = [][2]string{{label, focal}}
			} else {
				exist := false
				// If there is already a lens in the box with the same label, replace the old lens with the new lens
				for i, oldLens := range oldLensArr {
					if oldLens[0] == label {
						boxMap[k][i][1] = focal
						exist = true
					}
				}
				// If there is not already a lens in the box with the same label,
				// add the lens to the box immediately behind any lenses already in the box
				if !exist {
					oldLensArr = append(oldLensArr, [2]string{label, focal})
					boxMap[k] = oldLensArr
				}
			}
		} else {
			i := strings.Index(str, "-")
			label, _ := str[:i], str[i+1:]
			k := getHash(label)
			oldLensArr, ok := boxMap[k]
			// go to the relevant box and remove the lens with the given label if it is present in the box
			if ok {
				newArr := [][2]string{}
				for i, oldLens := range oldLensArr {
					if oldLens[0] == label {
						preArr, postArr := oldLensArr[:i], oldLensArr[i+1:]
						newArr = append(newArr, preArr...)
						newArr = append(newArr, postArr...)
						oldLensArr = newArr
						boxMap[k] = oldLensArr
					}
				}
			}
		}
		// fmt.Println(boxMap)

	}
	return boxMap
}

func getFocusingPower(boxMap map[int][][2]string) (power int) {
	for boxNum, lensArr := range boxMap {
		if len(lensArr) == 0 {
			continue
		}
		for i, lens := range lensArr {
			focal, _ := strconv.Atoi(lens[1])
			power += (boxNum + 1) * (i + 1) * focal
		}

	}
	return power
}
