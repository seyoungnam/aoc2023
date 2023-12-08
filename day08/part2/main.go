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
	lines := loadFile("../input.txt")
	inst, nodeMap, vals := getInstNodeMapAndVals(lines)
	// fmt.Println(inst)
	// fmt.Println(nodeMap)

	steps := []int{}
	for _, val := range vals {
		step := getStepsToZZZ(inst, nodeMap, val)
		steps = append(steps, step)
	}

	fmt.Println(steps)
	lcm := LCM2(steps)
	fmt.Println(lcm)

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

func getInstNodeMapAndVals(lines []string) (inst []string, nodeMap map[string]node, vals []string) {
	nodeMap = map[string]node{}
	for i, line := range lines {
		if i == 0 {
			inst = append(inst, strings.Split(line, "")...)
		} else if i == 1 {
			continue
		} else {
			curr := constructNodeMap(line, nodeMap)
			if curr[2:] == "Z" {
				vals = append(vals, curr)
			}
		}
	}
	return inst, nodeMap, vals
}

func constructNodeMap(line string, nodeMap map[string]node) (curr string) {
	curr = line[:3]
	lval := line[7:10]
	rval := line[12:15]

	nodeMap[curr] = node{lval, rval}
	return curr
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

		if curr[2:] == "Z" {
			return step
		}

		if i == len(inst)-1 {
			i = -1
		}
	}
	return -1
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func LCM2(steps []int) int {
	result := steps[0] * steps[1] / GCD(steps[0], steps[1])

	for i := 2; i < len(steps); i++ {
		result = LCM(result, steps[i])
	}

	return result
}
