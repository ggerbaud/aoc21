package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "4"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	g := makeGame(lines)
	total := part1(g)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	g = makeGame(lines)
	total = part2(g)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(g *game) int {
	for _, n := range g.numbers {
		for _, b := range g.boards {
			b.handle(n)
			if b.isBingo() {
				return b.points()
			}
		}
	}
	return 0
}

func part2(g *game) int {
	for _, n := range g.numbers {
		unwinBoards := make([]*board, 0)
		var lastWinBoard *board
		for _, b := range g.boards {
			b.handle(n)
			if !b.isBingo() {
				unwinBoards = append(unwinBoards, b)
			} else {
				lastWinBoard = b
			}
		}
		if len(unwinBoards) == 0 {
			return lastWinBoard.points()
		}
		g.boards = unwinBoards
	}
	return 0
}

func makeGame(lines []string) *game {
	numbers := utils.ListOfNumbers(lines[0], ",")
	g := &game{numbers: numbers, boards: make([]*board, 0)}
	nb := (len(lines) - 1) / 6
	idx := 2
	for i := 0; i < nb; i++ {
		idx2, b := makeBoard(lines, idx)
		g.boards = append(g.boards, b)
		idx = idx2
	}
	return g
}

func makeBoard(lines []string, idx int) (int, *board) {
	b := &board{}
	for i := 0; i < 5; i++ {
		line := utils.ListOfNumbers(lines[idx+i], " ")
		for j := 0; j < 5; j++ {
			b.grid[j][i] = line[j]
		}
	}
	return idx + 6, b
}

func (b *board) handle(x int) {
	b.last = x
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.grid[j][i] == x {
				b.marked[j][i] = true
				return
			}
		}
	}
}

func (b *board) isBingo() bool {
	for i := 0; i < 5; i++ {
		if b.marked[i][0] && b.marked[i][1] && b.marked[i][2] && b.marked[i][3] && b.marked[i][4] {
			return true
		}
		if b.marked[0][i] && b.marked[1][i] && b.marked[2][i] && b.marked[3][i] && b.marked[4][i] {
			return true
		}
	}
	return false
}

func (b *board) points() int {
	unmarked := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.marked[j][i] {
				unmarked += b.grid[j][i]
			}
		}
	}
	return unmarked * b.last
}

type (
	game struct {
		numbers []int
		boards  []*board
	}
	board struct {
		last   int
		grid   [5][5]int
		marked [5][5]bool
	}
)
