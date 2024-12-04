package main

import (
	"advent/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const day = "8"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	total := 0
	for _, line := range lines {
		lr := strings.Split(line, "|")
		fourDigits := strings.Split(strings.TrimSpace(lr[1]), " ")
		for _, fd := range fourDigits {
			if len(fd) == 2 || len(fd) == 3 || len(fd) == 4 || len(fd) == 7 {
				total++
			}
		}
	}
	return total
}

func part2(lines []string) int {
	total := 0
	for _, line := range lines {
		total += solve(line)
	}
	return total
}

func solve(line string) int {
	lr := strings.Split(line, "|")
	fourDigits := strings.Split(strings.TrimSpace(lr[1]), " ")
	r, ok := quickSolve(fourDigits)
	if ok {
		return r
	}
	wires := strings.Split(strings.TrimSpace(lr[0]), " ")
	return longSolve(wires, fourDigits)
}

func quickSolve(fd []string) (int, bool) {
	result := 0
	for i, d := range fd {
		x := 0
		if len(d) == 2 {
			x = 1
		} else if len(d) == 3 {
			x = 7
		} else if len(d) == 4 {
			x = 4
		} else if len(d) == 7 {
			x = 8
		} else {
			return 0, false
		}
		result += int(math.Pow(10, float64(3-i))) * x
	}
	return result, true
}

func longSolve(wires, fd []string) int {
	len2digit := make(map[int][]string)
	var up, down, mid, ur, ul, dr, dl int32
	var w0, w1, w2, w3, w4, w5, _, w7, w8, w9 string
	for _, w := range wires {
		utils.AddToMapOfSlice(len2digit, len(w), w)
		if len(w) == 2 {
			w1 = w
		} else if len(w) == 3 {
			w7 = w
		} else if len(w) == 4 {
			w4 = w
		} else if len(w) == 7 {
			w8 = w
		}
	}
	cOrf := make([]string, 0)
	for _, d := range w7 {
		if !strings.Contains(w1, string(d)) {
			up = d
		} else {
			cOrf = append(cOrf, string(d))
		}
	}
	for _, w := range len2digit[5] {
		if strings.Contains(w, cOrf[0]) && strings.Contains(w, cOrf[1]) {
			w3 = w
			break
		}
	}
	for _, w := range len2digit[6] {
		num := 0
		for _, x := range w3 {
			if !strings.Contains(w, string(x)) {
				break
			}
			num++
		}
		if num == 5 {
			w9 = w
		}
	}
	for _, d := range w8 {
		if !strings.Contains(w9, string(d)) {
			dl = d
			break
		}
	}
	for _, w := range len2digit[5] {
		if w == w3 {
			continue
		}
		if strings.Contains(w, string(dl)) {
			w2 = w
		} else {
			w5 = w
		}
	}
	for _, d := range w3 {
		if !strings.Contains(w2, string(d)) {
			dr = d
		}
		if !strings.Contains(w5, string(d)) {
			ur = d
		}
	}
	for _, w := range len2digit[6] {
		if w == w9 {
			continue
		}
		if strings.Contains(w, string(ur)) {
			w0 = w
		}
	}
	for _, d := range w4 {
		if !strings.Contains(w3, string(d)) {
			ul = d
		}
	}
	for _, d := range w3 {
		if !strings.Contains(w0, string(d)) {
			mid = d
		}
		if !strings.Contains(w5, string(d)) {
			ur = d
		}
	}
	for _, d := range w8 {
		if d != up && d != dl && d != ul && d != dr && d != mid && d != ur {
			down = d
		}
	}
	p2map := make(map[int32]int)
	p2map[up] = 1
	p2map[ul] = 2
	p2map[ur] = 4
	p2map[mid] = 8
	p2map[dl] = 16
	p2map[dr] = 32
	p2map[down] = 64
	result := 0
	pDix := 1
	for i := 3; i >= 0; i-- {
		d := fd[i]
		sum := 0
		for _, x := range d {
			sum += p2map[x]
		}
		result += constInt[sum] * pDix
		pDix *= 10
	}
	return result
}

const (
	zero  = 1 | 2 | 4 | 16 | 32 | 64
	one   = 4 | 32
	two   = 1 | 4 | 8 | 16 | 64
	three = 1 | 4 | 8 | 32 | 64
	four  = 2 | 4 | 8 | 32
	five  = 1 | 2 | 8 | 32 | 64
	six   = 1 | 2 | 8 | 16 | 32 | 64
	seven = 1 | 4 | 32
	eight = 1 | 2 | 4 | 8 | 16 | 32 | 64
	nine  = 1 | 2 | 4 | 8 | 32 | 64
)

var constInt = map[int]int{zero: 0, one: 1, two: 2, three: 3, four: 4, five: 5, six: 6, seven: 7, eight: 8, nine: 9}
