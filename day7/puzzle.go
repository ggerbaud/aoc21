package main

import (
	"advent/utils"
	"fmt"
	"math"
	"strconv"
)

const day = "7"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	pos := utils.ListOfNumbers(lines[0], ",")
	maxP := len(pos)
	minCost := cost(pos, 0, math.MaxInt)
	for i := 1; i < maxP; i++ {
		if i < len(pos) && pos[i] > maxP {
			maxP = pos[i]
		}
		costi := cost(pos, i, minCost)
		if costi < minCost {
			minCost = costi
		}
	}
	return minCost
}

func part2(lines []string) int {
	pos := utils.ListOfNumbers(lines[0], ",")
	maxP := len(pos)
	minCost := cost2(pos, 0, math.MaxInt)
	for i := 1; i < maxP; i++ {
		if i < len(pos) && pos[i] > maxP {
			maxP = pos[i]
		}
		costi := cost2(pos, i, minCost)
		if costi < minCost {
			minCost = costi
		}
	}
	return minCost
}

func cost(pos []int, toP, max int) int {
	total := 0
	for _, p := range pos {
		if total > max {
			return total
		}
		if p < toP {
			total += toP - p
		} else {
			total += p - toP
		}
	}
	return total
}

func cost2(pos []int, toP, max int) int {
	total := 0
	for _, p := range pos {
		if total > max {
			return total
		}
		n := 0
		if p < toP {
			n = toP - p
		} else {
			n = p - toP
		}
		total += n * (n + 1) / 2
	}
	return total
}
