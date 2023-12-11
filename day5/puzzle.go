package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "5"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	segs := makeSegments(lines, false)
	return doCalc(segs)
}

func part2(lines []string) int {
	segs := makeSegments(lines, true)
	return doCalc(segs)
}

func doCalc(segs []segment) int {
	data := make(map[string]int)
	for _, seg := range segs {
		for _, p := range seg.points() {
			if v, ok := data[p]; ok {
				data[p] = v + 1
			} else {
				data[p] = 1
			}
		}
	}
	total := 0
	for _, i := range data {
		if i > 1 {
			total++
		}
	}
	return total
}

func makeSegments(lines []string, diag bool) []segment {
	segments := make([]segment, 0)
	for _, line := range lines {
		start, end, _ := strings.Cut(line, " -> ")
		x1, y1 := coords(start)
		x2, y2 := coords(end)
		if diag || x1 == x2 || y1 == y2 {
			segments = append(segments, segment{x1, y1, x2, y2})
		}
	}
	return segments
}

func coords(s string) (int, int) {
	idx := strings.Index(s, ",")
	x, _ := strconv.Atoi(s[:idx])
	y, _ := strconv.Atoi(s[idx+1:])
	return x, y
}

func (s segment) points() []string {
	points := make([]string, 0)
	dx, dy := s.x2-s.x1, s.y2-s.y1
	absX, absY := utils.Max(dx, -dx), utils.Max(dy, -dy)
	dm := utils.Max(absX, absY)
	stepX, stepY := 0, 0
	if absX > 0 {
		stepX = dx / absX
	}
	if absY > 0 {
		stepY = dy / absY
	}
	for k := 0; k <= dm; k++ {
		points = append(points, fmt.Sprintf("%d,%d", s.x1+k*stepX, s.y1+k*stepY))
	}
	return points
}

type (
	segment struct {
		x1, y1, x2, y2 int
	}
)
