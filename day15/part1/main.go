package main

import (
	"bufio"
	"fmt"
	"os"
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
