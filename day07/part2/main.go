package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := loadFile("../input.txt")
	scoreBid, hand := getScoreBidAndHand(lines)
	fmt.Println(scoreBid)
	fmt.Println(hand)

	kind := "AKQT98765432J"
	table := &kind
	sortedScoreBid := sortArray(scoreBid, hand, table)
	total := getTotalScore(sortedScoreBid)
	fmt.Printf("The total winnings is : %v\n", total)

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

func getScoreBidAndHand(lines []string) (pairs [][2]int, hands []string) {
	for _, line := range lines {
		handBid := strings.Split(line, " ")
		score := getScore(handBid[0])
		hands = append(hands, handBid[0])
		bid := changeStrToInt(handBid[1])
		pairs = append(pairs, [2]int{score, bid})
	}
	return pairs, hands
}

func getScore(hand string) (score int) {
	table := [13]int{}
	for _, char := range hand {
		switch int(char) {
		case 65:
			table[0]++
		case 75:
			table[1]++
		case 81:
			table[2]++
		case 84:
			table[3]++
		case 57:
			table[4]++
		case 56:
			table[5]++
		case 55:
			table[6]++
		case 54:
			table[7]++
		case 53:
			table[8]++
		case 52:
			table[9]++
		case 51:
			table[10]++
		case 50:
			table[11]++
		case 74:
			table[12]++
		}
	}
	adjTable := getAdjustedTable(table)
	for _, v := range adjTable {
		score += v * v
	}
	return score
}

func getAdjustedTable(table [13]int) [13]int {
	largestNum := 0
	largestIdx := -1
	for i := 0; i < 12; i++ {
		if table[i] > largestNum {
			largestNum = table[i]
			largestIdx = i
		}
	}
	if largestIdx != -1 {
		table[largestIdx] += table[12]
		table[12] = 0
	}
	return table
}

func changeStrToInt(strNum string) (num int) {
	for _, str := range strNum {
		num *= 10
		num += int(str) - 48
	}
	return num
}

func sortArray(scoreBid [][2]int, hands []string, table *string) [][2]int {
	for i := 0; i <= len(scoreBid)-1; i++ {
		for j := 0; j < len(scoreBid)-1-i; j++ {
			if scoreBid[j][0] > scoreBid[j+1][0] || (scoreBid[j][0] == scoreBid[j+1][0] && leftIsLarger(hands[j], hands[j+1], table)) {
				scoreBid[j], scoreBid[j+1] = scoreBid[j+1], scoreBid[j]
				hands[j], hands[j+1] = hands[j+1], hands[j]
			}
		}
	}
	return scoreBid
}

func leftIsLarger(left string, right string, table *string) bool {
	for i := 0; i < len(left); i++ {
		li := strings.Index(*table, string(left[i]))
		ri := strings.Index(*table, string(right[i]))
		if li < ri {
			return true
		} else if li > ri {
			return false
		}
	}
	return false
}

func getTotalScore(scoreBid [][2]int) (tot int) {
	for i, v := range scoreBid {
		tot += (i + 1) * v[1]
	}
	return tot
}
