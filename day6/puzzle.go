package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "6"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	fish := utils.ListOfNumbers(lines[0], ",")
	return doCalc(fish, 80)
}

func part2(lines []string) int {
	fish := utils.ListOfNumbers(lines[0], ",")
	return doCalc(fish, 256)
}

func doCalc(fish []int, days int) int {
	dayMap := make([]int, 9)
	for _, f := range fish {
		dayMap[f]++
	}
	for i := 0; i < days; i++ {
		dayMap[0], dayMap[1], dayMap[2], dayMap[3], dayMap[4], dayMap[5], dayMap[6], dayMap[7], dayMap[8] = dayMap[1], dayMap[2], dayMap[3], dayMap[4], dayMap[5], dayMap[6], dayMap[0]+dayMap[7], dayMap[8], dayMap[0]
	}
	return dayMap[0] + dayMap[1] + dayMap[2] + dayMap[3] + dayMap[4] + dayMap[5] + dayMap[6] + dayMap[7] + dayMap[8]
}
