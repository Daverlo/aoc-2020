package main

import "testing"

func TestParseInput(t *testing.T) {
	path := "/Users/david/workspace/aoc-2020/test/day16/in"
	_, _, _ = parseInput(path)
}

func TestPart2(t *testing.T) {
	path := "/Users/david/workspace/aoc-2020/test/day16/in"
	rules, tickets, _ := parseInput(path)
	_ = part2(rules, tickets)
}
