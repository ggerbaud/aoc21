package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 26397
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData()
	result := part2(lines)
	expect := 288957
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestPartIllegalChar(t *testing.T) {
	lines := testIllegalCharData()
	for line, r := range lines {
		result, ok := illegalChar(line)
		if !ok {
			t.Fatalf("We expect %s to be corrupted", line)
		}
		if result != r {
			t.Fatalf("Illegal_char returns %s, we expect %s for %s", string(result), string(r), line)
		}
	}
}

func testData() []string {
	return []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}
}

func testIllegalCharData() map[string]rune {
	return map[string]rune{
		"{([(<{}[<>[]}>{[]{[(<()>": '}',
		"[[<[([]))<([[{}[[()]]]":   ')',
		"[{[{({}]{}}([{[{{{}}([]":  ']',
		"[<(<(<(<{}))><([]([]()":   ')',
		"<{([([[(<>()){}]>(<<{{":   '>',
	}
}
