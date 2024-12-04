package main

import (
	"advent/utils"
	"fmt"
	"sort"
	"strconv"
)

const day = "10"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	cpt := make(map[rune]int)
	for _, line := range lines {
		r, ok := illegalChar(line)
		if ok {
			n := cpt[r]
			cpt[r] = n + 1
		}
	}
	return cpt[')']*3 + cpt[']']*57 + cpt['}']*1197 + cpt['>']*25137
}

func part2(lines []string) int {
	scores := make([]int, 0)
	for _, line := range lines {
		ok, miss := incompleteChar(line)
		if ok {
			n := 0
			for len(miss) > 0 {
				n = 5 * n
				switch miss[len(miss)-1] {
				case ')':
					n += 1
				case ']':
					n += 2
				case '}':
					n += 3
				case '>':
					n += 4
				}
				miss = miss[:len(miss)-1]
			}
			scores = append(scores, n)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func illegalChar(line string) (rune, bool) {
	pile := make([]rune, 0)
	for _, c := range line {
		switch c {
		case '(':
			pile = append(pile, ')')
		case '[':
			pile = append(pile, ']')
		case '{':
			pile = append(pile, '}')
		case '<':
			pile = append(pile, '>')
		case ')', '>', '}', ']':
			if len(pile) == 0 || pile[len(pile)-1] != c {
				return c, true
			}
			pile = pile[:len(pile)-1]
		}
	}
	return 0, false
}

func incompleteChar(line string) (bool, []rune) {
	pile := make([]rune, 0)
	for _, c := range line {
		switch c {
		case '(':
			pile = append(pile, ')')
		case '[':
			pile = append(pile, ']')
		case '{':
			pile = append(pile, '}')
		case '<':
			pile = append(pile, '>')
		case ')', '>', '}', ']':
			if len(pile) == 0 || pile[len(pile)-1] != c {
				return false, nil
			}
			pile = pile[:len(pile)-1]
		}
	}
	return len(pile) > 0, pile
}
