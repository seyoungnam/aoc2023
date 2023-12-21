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

var xmas = map[string][2]int{
	"x": {1, 4000},
	"m": {1, 4000},
	"a": {1, 4000},
	"s": {1, 4000},
}

func main() {
	workflows, parts := loadFile("../input.txt")
	fmt.Println(workflows)
	fmt.Println(parts)

	wfMap := parseWorkflows(workflows)
	for k, v := range wfMap {
		fmt.Println(k, v)
	}

	res := getCount(&wfMap, xmas, "in")
	fmt.Println(res)

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

func getCount(wfMap *map[string]workflow, xmas map[string][2]int, name string) int {
	// fmt.Printf("name: %v | xmas: %v\n", name, xmas)
	if name == "R" {
		return 0
	}
	if name == "A" {
		fmt.Printf("Accepted xmas: %v\n", xmas)
		res := 1
		for _, val := range xmas {
			res *= (val[1] - val[0] + 1)
		}
		return res
	}

	total := 0

	wf, _ := (*wfMap)[name]
	runAllRules := true
	var newXmas map[string][2]int
	for _, r := range wf.rules {
		c := r.xmas
		lo, hi := xmas[c][0], xmas[c][1]
		t, f := [2]int{}, [2]int{}
		if r.lt {
			t = [2]int{max(lo, r.num+1), hi}
			f = [2]int{lo, min(r.num, hi)}
		} else {
			t = [2]int{lo, min(r.num-1, hi)}
			f = [2]int{max(lo, r.num), hi}
		}
		if t[0] <= t[1] {
			newXmas = copyMap(xmas)
			newXmas[c] = t
			total += getCount(wfMap, newXmas, r.dest)
		}
		if f[0] <= f[1] {
			xmas[c] = f
		} else {
			runAllRules = false
			break
		}
	}

	if runAllRules {
		total += getCount(wfMap, xmas, wf.last)
	}
	return total
}

func copyMap(a map[string][2]int) map[string][2]int {
	b := make(map[string][2]int)
	for k, v := range a {
		b[k] = v
	}
	return b
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
