package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(path string) (int, []int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, nil, err
	}

	var time int
	_, _ = fmt.Fscanf(f, "%d\n", &time)

	var line string
	_, _ = fmt.Fscanf(f, "%s\n", &line)

	routes := make([]int, 0)
	for _, r := range strings.Split(line, ",") {
		if r == "x" {
			continue
		}
		v, _ := strconv.Atoi(r)
		routes = append(routes, v)
	}

	return time, routes, nil
}

func part1(arrival int, routes []int) int {
	time := arrival
	for {
		// fmt.Println(time)
		for _, r := range routes {
			if time%r == 0 {
				return (time - arrival) * r
			}
		}
		time++
	}
}

func part2(numbers []int) int {
	return 0
}

func main() {
	args := os.Args[1:]
	time, routes, err := parseInput(args[0])
	if err != nil {
		panic(err)
	}

	output := part1(time, routes)
	fmt.Println(output)

	// output = part2(input)
	// fmt.Println(output)
}
