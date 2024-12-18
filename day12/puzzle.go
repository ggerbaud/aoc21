package main

import (
	"advent/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const day = "12"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	caves := make(map[string]*cave)
	for _, line := range lines {
		conn := strings.Split(line, "-")
		from := getOrMake(caves, conn[0])
		to := getOrMake(caves, conn[1])
		from.links = append(from.links, to)
		to.links = append(to.links, from)
	}
	paths := allPaths(caves)
	return len(paths)
}

func part2(lines []string) int {
	caves := make(map[string]*cave)
	for _, line := range lines {
		conn := strings.Split(line, "-")
		from := getOrMake(caves, conn[0])
		to := getOrMake(caves, conn[1])
		from.links = append(from.links, to)
		to.links = append(to.links, from)
	}
	paths := allPaths2(caves)
	return len(paths)
}

func allPaths(caves map[string]*cave) []path {
	start, end := caves["start"], caves["end"]
	if start == nil || end == nil {
		return nil
	}
	paths := make([]path, 0)
	queue := []path{{start}}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		last := p[len(p)-1]
		if last == end {
			paths = append(paths, p)
		} else {
			for _, c := range last.links {
				if c.big || !slices.Contains(p, c) {
					newPath := make(path, len(p))
					copy(newPath, p)
					newPath = append(newPath, c)
					queue = append(queue, newPath)
				}
			}
		}
	}
	return paths
}

func allPaths2(caves map[string]*cave) []path {
	start, end := caves["start"], caves["end"]
	if start == nil || end == nil {
		return nil
	}
	paths := make([]path, 0)
	queue := []path2{{false, path{start}}}
	for len(queue) > 0 {
		p2 := queue[0]
		queue = queue[1:]
		last := p2.p[len(p2.p)-1]
		if last == end {
			paths = append(paths, p2.p)
		} else {
			for _, c := range last.links {
				contains := slices.Contains(p2.p, c)
				if c != start && (c.big || !p2.visited || !contains) {
					newPath := make(path, len(p2.p))
					copy(newPath, p2.p)
					newPath = append(newPath, c)
					visited := p2.visited || (!c.big && contains)
					queue = append(queue, path2{visited: visited, p: newPath})
				}
			}
		}
	}
	return paths
}

func getOrMake(caves map[string]*cave, name string) *cave {
	c, ok := caves[name]
	if !ok {
		c = &cave{name: name, big: strings.ToUpper(name) == name}
		caves[name] = c
	}
	return c
}

type path []*cave
type path2 struct {
	visited bool
	p       path
}
type cave struct {
	name  string
	big   bool
	links []*cave
}
