package main

import "testing"

func TestPart1_1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 10
	if result != expect {
		t.Fatalf("Part1_1 returns %d, we expect %d", result, expect)
	}
}

func TestPart1_2(t *testing.T) {
	lines := testData2()
	result := part1(lines)
	expect := 19
	if result != expect {
		t.Fatalf("Part1_2 returns %d, we expect %d", result, expect)
	}
}

func TestPart1_3(t *testing.T) {
	lines := testData3()
	result := part1(lines)
	expect := 226
	if result != expect {
		t.Fatalf("Part1_3 returns %d, we expect %d", result, expect)
	}
}

func TestPart2_1(t *testing.T) {
	lines := testData()
	result := part2(lines)
	expect := 36
	if result != expect {
		t.Fatalf("Part2_1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2_2(t *testing.T) {
	lines := testData2()
	result := part2(lines)
	expect := 103
	if result != expect {
		t.Fatalf("Part2_2 returns %d, we expect %d", result, expect)
	}
}

func TestPart2_3(t *testing.T) {
	lines := testData3()
	result := part2(lines)
	expect := 3509
	if result != expect {
		t.Fatalf("Part2_2 returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{
		"start-A",
		"start-b",
		"A-c",
		"A-b",
		"b-d",
		"A-end",
		"b-end",
	}
}

func testData2() []string {
	return []string{
		"dc-end",
		"HN-start",
		"start-kj",
		"dc-start",
		"dc-HN",
		"LN-dc",
		"HN-end",
		"kj-sa",
		"kj-HN",
		"kj-dc",
	}
}

func testData3() []string {
	return []string{
		"fs-end",
		"he-DX",
		"fs-he",
		"start-DX",
		"pj-DX",
		"end-zg",
		"zg-sl",
		"zg-pj",
		"pj-he",
		"RW-he",
		"fs-DX",
		"pj-RW",
		"zg-RW",
		"start-pj",
		"he-WI",
		"zg-he",
		"pj-fs",
		"start-RW",
	}
}
