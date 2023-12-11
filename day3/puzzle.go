package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "3"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	n := len(lines[0])
	gamma, epsilon := 0, 0
	for i := 0; i < n; i++ {
		zeros, ones := getZerosAndOnes(lines, i, "")
		if ones > zeros {
			gamma = gamma | (1 << (n - i - 1))
		} else {
			epsilon = epsilon | (1 << (n - i - 1))
		}
	}
	return gamma * epsilon
}

func part2(lines []string) int {
	oxStr := seekOxygen(lines, 0)
	ox, _ := strconv.ParseInt(oxStr, 2, 64)
	co2Str := seekCO2(lines, 0)
	co2, _ := strconv.ParseInt(co2Str, 2, 64)
	return int(ox * co2)
}

func seekOxygen(lines []string, pos int) string {
	zeros, ones := getZerosAndOnesFiltered(lines, pos)
	selected := ones
	if len(zeros) > len(ones) {
		selected = zeros
	}
	if len(selected) == 1 {
		return selected[0]
	}
	return seekOxygen(selected, pos+1)
}

func seekCO2(lines []string, pos int) string {
	zeros, ones := getZerosAndOnesFiltered(lines, pos)
	selected := zeros
	if len(zeros) > len(ones) {
		selected = ones
	}
	if len(selected) == 1 {
		return selected[0]
	}
	return seekCO2(selected, pos+1)
}

func getZerosAndOnes(lines []string, pos int, filter string) (int, int) {
	ones, zeros := 0, 0
	for _, line := range lines {
		if strings.HasPrefix(line, filter) {
			if line[pos] == '1' {
				ones++
			} else {
				zeros++
			}
		}
	}
	return zeros, ones
}

func getZerosAndOnesFiltered(lines []string, pos int) ([]string, []string) {
	ones, zeros := make([]string, 0), make([]string, 0)
	for _, line := range lines {
		if line[pos] == '1' {
			ones = append(ones, line)
		} else {
			zeros = append(zeros, line)
		}
	}
	return zeros, ones
}
