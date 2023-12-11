package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 5934
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData()
	result := part2(lines)
	expect := 26984457539
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestDoCalc(t *testing.T) {
	fish := []int{3, 4, 3, 1, 2}
	result := doCalc(fish, 18)
	expect := 26
	if result != expect {
		t.Fatalf("doCalc returns %d, we expect %d", result, expect)
	}
	result = doCalc(fish, 80)
	expect = 5934
	if result != expect {
		t.Fatalf("doCalc returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{
		"3,4,3,1,2",
	}
}
