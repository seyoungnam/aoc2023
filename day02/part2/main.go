package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	res := 0

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		res += getPower(line)
	}

	fmt.Println("The sum is : ", res)
}

func getPower(line string) int {
	i := strings.Index(line, ":")
	return calculatePower(line[i+1:])
}

func calculatePower(gamesStr string) (num int) {
	bag := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	// gamesArr = [" 2 blue, 3 red", " 3 green, 3 blue, 6 red", " 4 blue, 6 red", " 2 green, 2 blue, 9 red", " 2 red, 4 blue"]
	gamesArr := strings.Split(gamesStr, ";") // gameStr = " 3 green, 3 blue, 6 red"
	for _, gameStr := range gamesArr {
		sets := strings.Split(gameStr, ",") // sets = [" 3 green", "3 blue", "6 red"]
		for _, setStr := range sets {
			setArr := strings.Split(setStr, " ") // setArr = [" ", "3", "green"] or ["3", "blue"]
			if len(setArr) != 2 {
				digit, err := strconv.Atoi(setArr[1])
				if err != nil {
					panic(err)
				}
				if digit > bag[setArr[2]] {
					bag[setArr[2]] = digit
				}
			} else {
				digit, err := strconv.Atoi(setArr[0])
				if err != nil {
					panic(err)
				}
				if digit > bag[setArr[1]] {
					bag[setArr[1]] = digit
				}
			}
		}
	}

	num = 1
	for _, val := range bag {
		if val != 0 {
			num *= val
		}
	}
	return num
}
