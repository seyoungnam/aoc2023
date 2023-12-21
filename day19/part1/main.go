package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type workflow struct {
	name  string
	rules []rule
	last  string
}

type rule struct {
	xmas string
	lt   bool
	num  int
	dest string
}

type part struct {
	x int
	m int
	a int
	s int
}

func main() {
	workflows, parts := loadFile("../input.txt")
	fmt.Println(workflows)
	fmt.Println(parts)

	wfMap := parseWorkflows(workflows)
	for k, v := range wfMap {
		fmt.Println(k, v)
	}

	partArr := parseParts(parts)
	for _, p := range partArr {
		fmt.Println(p)
	}

	res := 0
	for _, p := range partArr {
		subRes := getScoreForPart(&wfMap, p)
		fmt.Printf("subRes is %v\n", subRes)
		res += subRes
	}
	fmt.Printf("The answer is %v\n", res)

}

func loadFile(fileName string) (workflows, parts []string) {
	workflows = []string{}
	parts = []string{}
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)

	isRule := true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			isRule = false
			continue
		}
		if isRule {
			workflows = append(workflows, line)
		} else {
			parts = append(parts, line)
		}
	}
	return workflows, parts
}

func parseWorkflows(wfArr []string) (wfMap map[string]workflow) {
	wfMap = map[string]workflow{}

	for _, wf := range wfArr {
		i := strings.Index(wf, "{")
		j := strings.Index(wf, "}")
		name := wf[:i]
		rulesArr := strings.Split(wf[i+1:j], ",")
		last := rulesArr[len(rulesArr)-1]
		rulesArr = rulesArr[:len(rulesArr)-1]
		rules := []rule{}
		for _, ruleString := range rulesArr {
			ruleArr := strings.Split(ruleString, ":")
			part := string(ruleArr[0][0])
			lt := true
			if string(ruleArr[0][1]) == "<" {
				lt = false
			}
			num, _ := strconv.Atoi(ruleArr[0][2:])
			dest := ruleArr[1]
			newRule := rule{
				xmas: part,
				lt:   lt,
				num:  num,
				dest: dest,
			}
			rules = append(rules, newRule)
		}
		workflowStruct := workflow{
			name:  name,
			rules: rules,
			last:  last,
		}
		wfMap[name] = workflowStruct
	}
	return wfMap
}

func parseParts(parts []string) (partStructs []part) {
	for _, line := range parts {
		xmasArr := strings.Split(line[1:len(line)-1], ",")
		pt := part{}
		for _, xmas := range xmasArr {
			key := string(xmas[0])
			num, _ := strconv.Atoi(xmas[2:])

			if key == "x" {
				pt.x = num
			} else if key == "m" {
				pt.m = num
			} else if key == "a" {
				pt.a = num
			} else {
				pt.s = num
			}
		}
		partStructs = append(partStructs, pt)
	}
	return partStructs
}

func getScoreForPart(wfMap *map[string]workflow, part part) int {
	key := "in"
	for {
		fmt.Println(part, key)
		if key == "A" {
			return part.x + part.m + part.a + part.s
		} else if key == "R" {
			return 0
		} else {
			next := "placeholder"
			wf, _ := (*wfMap)[key]
			for _, r := range wf.rules {
				tgt := 0
				if r.xmas == "x" {
					tgt = part.x
				} else if r.xmas == "m" {
					tgt = part.m
				} else if r.xmas == "a" {
					tgt = part.a
				} else if r.xmas == "s" {
					tgt = part.s
				}
				if (r.lt && tgt > r.num) || (!r.lt && tgt < r.num) {
					next = r.dest
					break
				}
			}
			if next != "placeholder" {
				key = next
			} else {
				key = wf.last
			}
		}
	}
}
