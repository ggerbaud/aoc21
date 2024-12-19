package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "13"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	d, folds := makeData(lines)
	//d.print()
	//println()
	d = d.fold(folds[0])
	//d.print()
	return d.count()
}

func part2(lines []string) int {
	d, folds := makeData(lines)
	for _, f := range folds {
		d = d.fold(f)
	}
	d.print()
	return 0
}

func makeData(lines []string) (data, []fold) {
	d := data{}
	brk := 0
	for i, line := range lines {
		if line == "" {
			brk = i
			break
		}
		xy := utils.ListOfNumbers(line, ",")
		d = d.set(xy[0], xy[1], true)
	}
	folds := make([]fold, 0)
	for _, line := range lines[brk+1:] {
		parts := strings.Split(line, " ")
		fd := strings.Split(parts[2], "=")
		folds = append(folds, fold{utils.ParseInt(fd[1]), fd[0] == "x"})
	}
	return d, folds
}

type data [][]bool
type fold struct {
	axis int
	isX  bool
}

func (d data) count() int {
	c := 0
	for _, v := range d {
		for _, b := range v {
			if b {
				c++
			}
		}
	}
	return c
}

func (d data) print() {
	for _, v := range d {
		for _, b := range v {
			if b {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (d data) set(x, y int, b bool) data {
	out := d
	for len(out) <= y {
		out = append(out, make([]bool, 0))
	}
	for len(out[y]) <= x {
		out[y] = append(out[y], false)
	}
	out[y][x] = b
	return out
}

func (d data) get(x, y int) bool {
	if y < 0 || y >= len(d) || x < 0 || x >= len(d[y]) {
		return false
	}
	return d[y][x]
}

func (d data) fold(f fold) data {
	if f.isX {
		return d.foldV(f.axis)
	}
	return d.foldH(f.axis)
}

func (d data) foldH(axis int) data {
	out := make(data, axis)
	for j := range out {
		out[j] = make([]bool, len(d[j]))
		copy(out[j], d[j])
	}
	for j := axis + 1; j < len(d); j++ {
		dst := axis - (j - axis)
		if dst >= 0 {
			for i, b := range d[j] {
				out = out.set(i, dst, out.get(i, dst) || b)
			}
		}
	}
	return out
}

func (d data) foldV(axis int) data {
	out := make(data, len(d))
	for j := range out {
		out[j] = make([]bool, axis)
		copy(out[j], d[j])
	}
	for j := 0; j < len(d); j++ {
		for i := axis + 1; i < len(d[j]); i++ {
			dst := axis - (i - axis)
			if dst >= 0 {
				out = out.set(dst, j, out.get(dst, j) || d[j][i])
			}
		}
	}
	return out
}
