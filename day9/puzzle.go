package main

import (
	"advent/utils"
	"fmt"
	"sort"
	"strconv"
)

const day = "9"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	data := make(map[pt]int)
	for i, line := range lines {
		for j := range line {
			data[pt{i, j}] = utils.ParseInt(line[j : j+1])
		}
	}
	risk := 0
	for p, v := range data {
		if isLow(data, p) {
			risk += v + 1
		}
	}
	return risk
}

func part2(lines []string) int {
	data := make(map[pt]int)
	for i, line := range lines {
		for j := range line {
			data[pt{i, j}] = utils.ParseInt(line[j : j+1])
		}
	}
	basins := make([]int, 0)
	for p, _ := range data {
		if isLow(data, p) {
			basins = append(basins, basinSize(data, p))
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	return basins[0] * basins[1] * basins[2]
}

func isLow(data map[pt]int, p pt) bool {
	v := data[p]
	return checkPoint(data, p.i+1, p.j, v) &&
		checkPoint(data, p.i-1, p.j, v) &&
		checkPoint(data, p.i, p.j+1, v) &&
		checkPoint(data, p.i, p.j-1, v)
}

func checkPoint(data map[pt]int, i, j, v int) bool {
	if x, ok := data[pt{i, j}]; ok {
		return x > v
	}
	return true
}

func basinSize(data map[pt]int, p pt) int {
	pts := make(map[pt]interface{})
	pts[p] = nil
	queue := make([]pt, 0)
	queue = append(queue, pt{p.i - 1, p.j})
	queue = append(queue, pt{p.i + 1, p.j})
	queue = append(queue, pt{p.i, p.j - 1})
	queue = append(queue, pt{p.i, p.j + 1})
	for len(queue) > 0 {
		pt0 := queue[0]
		queue = queue[1:]
		if _, ok := pts[pt0]; ok {
			continue
		}
		v, ok := data[pt0]
		if ok && v != 9 {
			queue = append(queue, pt{pt0.i - 1, pt0.j})
			queue = append(queue, pt{pt0.i + 1, pt0.j})
			queue = append(queue, pt{pt0.i, pt0.j - 1})
			queue = append(queue, pt{pt0.i, pt0.j + 1})
			pts[pt0] = nil
		}
	}
	return len(pts)
}

type pt struct {
	i, j int
}
