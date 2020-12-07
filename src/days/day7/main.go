package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Color string
type Contents map[Color]int

func parseInput(path string) (map[Color]Contents, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	rd := bufio.NewReader(f)

	rules := make(map[Color]Contents, 0)

	for {
		// light red bags contain 1 bright white bag, 2 muted yellow bags.
		var line string
		line, err = rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		line = strings.TrimSuffix(line, "\n")

		r := regexp.MustCompile("^([a-z]+ [a-z]+) bags contain (.*)")
		matches := r.FindAllStringSubmatch(line, -1)

		color := matches[0][1]
		contents, err := parseContents(matches[0][2])
		if err != nil {
			return nil, err
		}
		rules[Color(color)] = contents
	}

	return rules, nil
}

func parseContents(rawContents string) (Contents, error) {
	contents := make(map[Color]int)

	r := regexp.MustCompile("([0-9]) ([a-z]+ [a-z]+) bags?")
	matches := r.FindAllStringSubmatch(rawContents, -1)

	for _, match := range matches {
		count, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, err
		}
		contents[Color(match[2])] = count
	}

	return contents, nil
}

func part1(rules map[Color]Contents) int {
	transposed := transposeRules(rules)
	return traverse(transposed, Color("shiny gold")) - 1
}

func transposeRules(rules map[Color]Contents) map[Color]Contents {
	transposed := make(map[Color]Contents, len(rules))

	for color, contents := range rules {
		for containedColor, quantity := range contents {
			if transposed[containedColor] == nil {
				transposed[containedColor] = Contents{}
			}
			transposed[containedColor][color] = quantity
		}
	}

	return transposed
}

// Count all the nodes that can be traversed from the starting point
func traverse(rules map[Color]Contents, start Color) int {
	visited := make(map[Color]bool)

	q := list.New()
	q.PushBack(start)
	for q.Len() > 0 {
		e := q.Front()
		q.Remove(e)
		c := e.Value.(Color)

		visited[c] = true
		for target, _ := range rules[c] {
			if _, ok := visited[target]; !ok {
				q.PushBack(target)
			}
		}
	}

	return len(visited)
}

func part2(rules map[Color]Contents) int {
	return countBags(rules, Color("shiny gold"))
}

func countBags(rules map[Color]Contents, color Color) int {
	res := 0
	for target, quantity := range rules[color] {
		res += quantity * (1 + countBags(rules, target))
	}

	return res
}

func main() {
	args := os.Args[1:]
	input, err := parseInput(args[0])
	if err != nil {
		panic(err)
	}

	output := part1(input)
	fmt.Println(output)

	output = part2(input)
	fmt.Println(output)
}
