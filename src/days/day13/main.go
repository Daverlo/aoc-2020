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

type Route struct {
	Id     int
	Offset int
}

func parseInput2(path string) (int, []Route, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, nil, err
	}

	var time int
	_, _ = fmt.Fscanf(f, "%d\n", &time)

	var line string
	_, _ = fmt.Fscanf(f, "%s\n", &line)

	routes := make([]Route, 0)
	offset := -1
	for _, r := range strings.Split(line, ",") {
		offset++
		if r == "x" {
			continue
		}
		v, _ := strconv.Atoi(r)
		routes = append(routes, Route{Id: v, Offset: offset})
	}

	return time, routes, nil
}

func part2(arrival int, routes []Route) int {
	t := 1
	step := 1
	for _, r := range routes {
		t, step = f(t, step, r.Id, r.Offset)
	}
	return t
}

func f(t int, step int, id int, offset int) (int, int) {
	// fmt.Println(t, step, id, offset)
	for {
		if (t+offset)%id == 0 {
			return t, LCM(step, id)
		}

		t += step
	}
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int) int {
	result := a * b / GCD(a, b)
	return result
}

func main() {
	args := os.Args[1:]
	time, routes, err := parseInput(args[0])
	if err != nil {
		panic(err)
	}

	output := part1(time, routes)
	fmt.Println(output)

	time, routes2, err := parseInput2(args[0])
	if err != nil {
		panic(err)
	}
	output = part2(time, routes2)
	fmt.Println(output)
}
