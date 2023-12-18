package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Element struct {
	hl int
	r  int
	c  int
	dr int
	dc int
	n  int
}

type MinHeap struct {
	array []Element
}

func (h *MinHeap) Insert(ele Element) {
	h.array = append(h.array, ele)
	h.minHeapifyUp(len(h.array) - 1)
}

// Extract returns the smallest key, and removes it from the heap
func (h *MinHeap) Extract() Element {
	minElement := h.array[0]
	lastIdx := len(h.array) - 1
	// when the array is empty
	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return Element{}
	}
	// take out the last index and put it in the root
	h.array[0] = h.array[lastIdx]
	h.array = h.array[:lastIdx]

	h.minHeapifyDown(0)
	return minElement
}

// minHeapifyUp will heapify from bottom to top
func (h *MinHeap) minHeapifyUp(index int) {
	for h.array[parent(index)].hl > h.array[index].hl {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// minHeapifyDown will heapify top to bottom
func (h *MinHeap) minHeapifyDown(index int) {
	lastIdx := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	// loop while index has at least one child
	for l <= lastIdx {
		if l == lastIdx { // when left child is the only child
			childToCompare = l
		} else if h.array[l].hl < h.array[r].hl { // when left child is smaller
			childToCompare = l
		} else { // when right child is smaller
			childToCompare = r
		}
		// compare array value of current index to smaller child and swap if larger
		if h.array[index].hl > h.array[childToCompare].hl {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index
func left(i int) int {
	return 2*i + 1
}

// get the right child index
func right(i int) int {
	return 2*i + 2
}

// swap keys in the array
func (h *MinHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

var dir = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

var mark = [4]string{
	"<",
	"v",
	">",
	"^",
}

func main() {
	grid := loadFile("../input.txt")
	fmt.Println("original grid: ")
	for _, row := range grid {
		fmt.Println(row)
	}

	// m := &MinHeap{}

	// ele1 := Element{50, 1, 1, 1, 1, 1}
	// ele2 := Element{49, 2, 2, 2, 2, 2}
	// ele3 := Element{45, 3, 3, 3, 3, 3}
	// ele4 := Element{12, 4, 4, 4, 4, 4}
	// ele5 := Element{68, 5, 5, 5, 5, 5}
	// ele6 := Element{99, 6, 6, 6, 6, 6}

	// buildHeap := []Element{ele1, ele2, ele3, ele4, ele5, ele6}
	// for _, v := range buildHeap {
	// 	m.Insert(v)
	// 	fmt.Println(m)
	// }

	// for i := 0; i < 5; i++ {
	// 	m.Extract()
	// 	fmt.Println(m)
	// }

	seen := map[Element]bool{}

	minHeap := &MinHeap{}
	minHeap.Insert(Element{0, 0, 0, 0, 0, 0})

	for len(minHeap.array) != 0 {
		curr := minHeap.Extract()

		if curr.r == len(grid)-1 && curr.c == len(grid[0])-1 && curr.n >= 4 {
			fmt.Println(curr.hl)
			break
		}

		if _, ok := seen[curr]; ok {
			continue
		}
		seen[curr] = true

		if curr.n < 10 && !(curr.dr == 0 && curr.dc == 0) {
			nr := curr.r + curr.dr
			nc := curr.c + curr.dc
			if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) {
				minHeap.Insert(Element{curr.hl + grid[nr][nc], nr, nc, curr.dr, curr.dc, curr.n + 1})
			}
		}

		if curr.n >= 4 || (curr.dr == 0 && curr.dc == 0) {
			for _, nd := range dir {
				if !(nd[0] == curr.dr && nd[1] == curr.dc) && !(nd[0] == -curr.dr && nd[1] == -curr.dc) {
					nr := curr.r + nd[0]
					nc := curr.c + nd[1]
					if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) {
						minHeap.Insert(Element{curr.hl + grid[nr][nc], nr, nc, nd[0], nd[1], 1})
					}
				}
			}
		}
	}

}

func loadFile(fileName string) (grid [][]int) {
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Split(line, "")
		numArr := []int{}
		for _, str := range lineArr {
			num, _ := strconv.Atoi(str)
			numArr = append(numArr, num)
		}
		grid = append(grid, numArr)
	}
	return grid
}
