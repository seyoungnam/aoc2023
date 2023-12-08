package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	left  string
	right string
}

func main() {
	lines := loadFile("../input2.txt")
	inst, nodeMap := getInstAndNodeMap(lines)
	// fmt.Println(inst)
	// fmt.Println(nodeMap)

	step := getStepsToZZZ(inst, nodeMap, "AAA")
	fmt.Println(step)

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

func getInstAndNodeMap(lines []string) (inst []string, nodeMap map[string]node) {
	nodeMap = map[string]node{}
	for i, line := range lines {
		if i == 0 {
			inst = append(inst, strings.Split(line, "")...)
		} else if i == 1 {
			continue
		} else {
			constructNodeMap(line, nodeMap)
		}
	}
	return inst, nodeMap
}

func constructNodeMap(line string, nodeMap map[string]node) {
	curr := line[:3]
	lval := line[7:10]
	rval := line[12:15]

	nodeMap[curr] = node{lval, rval}
}

func getStepsToZZZ(inst []string, nodeMap map[string]node, begVal string) (step int) {
	curr := begVal
	for i := 0; i < len(inst); i++ {
		if inst[i] == "L" {
			curr = nodeMap[curr].left
		} else {
			curr = nodeMap[curr].right
		}
		step++

		if curr == "ZZZ" {
			return step
		}

		if i == len(inst)-1 {
			i = -1
		}
	}
	return -1
}
