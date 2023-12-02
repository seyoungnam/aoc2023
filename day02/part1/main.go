package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Bag = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

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
		num, ok := getGameNumAndJudge(line)
		if ok {
			res += num
		}
	}

	fmt.Println("The sum is : ", res)
}

func getGameNumAndJudge(line string) (num int, judge bool) {
	i := strings.Index(line, ":")
	numStr := line[5:i]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return num, getJudge(line[i+1:])
}

func getJudge(gamesStr string) (judge bool) {
	judge = true
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
				if digit > Bag[setArr[2]] {
					judge = false
					break
				}
			} else {
				digit, err := strconv.Atoi(setArr[0])
				if err != nil {
					panic(err)
				}
				if digit > Bag[setArr[1]] {
					judge = false
					break
				}
			}
		}
	}
	return judge
}
