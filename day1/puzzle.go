package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "1"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	total := 0
	prev := 0
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		if num > prev {
			total++
		}
		prev = num
	}
	return total - 1
}

func part2(lines []string) int {
	total := 0
	a, b, c := utils.ParseInt(lines[0]), utils.ParseInt(lines[1]), utils.ParseInt(lines[2])
	prev := a + b + c
	for i := 3; i < len(lines); i++ {
		n := utils.ParseInt(lines[i])
		num := n + b + c
		if num > prev {
			total++
		}
		prev = num
		a, b, c = b, c, n
	}
	return total
}
