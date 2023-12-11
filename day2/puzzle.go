package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "2"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	p := &position{}
	for _, line := range lines {
		p.handle(line)
	}
	return p.depth * p.horizontal
}

func part2(lines []string) int {
	p := &position{}
	for _, line := range lines {
		p.handle2(line)
	}
	return p.depth * p.horizontal
}

func (p *position) handle(instr string) {
	if strings.HasPrefix(instr, "forward ") {
		p.horizontal += utils.ParseInt(instr[8:])
	} else if strings.HasPrefix(instr, "up ") {
		p.depth -= utils.ParseInt(instr[3:])
	} else if strings.HasPrefix(instr, "down ") {
		p.depth += utils.ParseInt(instr[5:])
	}
}

func (p *position) handle2(instr string) {
	if strings.HasPrefix(instr, "forward ") {
		n := utils.ParseInt(instr[8:])
		p.horizontal += n
		p.depth += n * p.aim
	} else if strings.HasPrefix(instr, "up ") {
		p.aim -= utils.ParseInt(instr[3:])
	} else if strings.HasPrefix(instr, "down ") {
		p.aim += utils.ParseInt(instr[5:])
	}
}

type (
	position struct {
		horizontal, depth, aim int
	}
)
