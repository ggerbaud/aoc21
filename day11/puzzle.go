package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "11"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	octopuses := make([][]int, len(lines))
	for j, line := range lines {
		octopuses[j] = make([]int, len(line))
		for i, d := range line {
			octopuses[j][i] = int(d - '0')
		}
	}
	flashes := 0
	for k := 0; k < 100; k++ {
		// incr
		for y, octopus := range octopuses {
			for x := range octopus {
				octopuses[y][x]++
			}
		}
		hasFlash := true
		for hasFlash {
			hasFlash = false
			for j, octopus := range octopuses {
				for i, o := range octopus {
					if o > 9 {
						octopuses[j][i] = -1
						incrAround(octopuses, i, j)
						hasFlash = true
						flashes++
					}
				}
			}
		}
		for j, octopus := range octopuses {
			for i, o := range octopus {
				if o == -1 {
					octopuses[j][i] = 0
				}
			}
		}
	}
	return flashes
}

func part2(lines []string) int {
	octopuses := make([][]int, len(lines))
	for j, line := range lines {
		octopuses[j] = make([]int, len(line))
		for i, d := range line {
			octopuses[j][i] = int(d - '0')
		}
	}
	step := 0
	flashes := 0
	for flashes < 100 {
		flashes = 0
		// incr
		for y, octopus := range octopuses {
			for x := range octopus {
				octopuses[y][x]++
			}
		}
		hasFlash := true
		for hasFlash {
			hasFlash = false
			for j, octopus := range octopuses {
				for i, o := range octopus {
					if o > 9 {
						octopuses[j][i] = -1
						incrAround(octopuses, i, j)
						hasFlash = true
						flashes++
					}
				}
			}
		}
		for j, octopus := range octopuses {
			for i, o := range octopus {
				if o == -1 {
					octopuses[j][i] = 0
				}
			}
		}
		step++
	}
	return step
}

func incrAround(data [][]int, x, y int) {
	incr(data, x-1, y-1)
	incr(data, x-1, y)
	incr(data, x-1, y+1)
	incr(data, x, y+1)
	incr(data, x, y-1)
	incr(data, x+1, y-1)
	incr(data, x+1, y)
	incr(data, x+1, y+1)
}
func incr(data [][]int, x, y int) {
	if y >= 0 && y < len(data) && x >= 0 && x < len(data[y]) && data[y][x] >= 0 {
		data[y][x]++
	}
}
