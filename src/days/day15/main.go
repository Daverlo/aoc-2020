package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var line string
	_, _ = fmt.Fscanf(f, "%s\n", &line)

	numbers := make([]int, 0)
	for _, r := range strings.Split(line, ",") {
		v, _ := strconv.Atoi(r)
		numbers = append(numbers, v)
	}

	return numbers, nil
}

func part1(numbers []int) int {
	return play(2020, numbers)
}

func play(target int, numbers []int) int {
	spoken := make(map[int][]int)
	turn := 1
	var lastTurnNumber int

	for _, n := range numbers {
		// fmt.Println(turn, lastTurnNumber, n)
		if spoken[n] == nil {
			spoken[n] = make([]int, 0)
		}
		spoken[n] = append(spoken[n], turn)

		turn++
		lastTurnNumber = n
	}

	for turn <= target {
		var n int
		if len(spoken[lastTurnNumber]) == 1 {
			n = 0
		} else {
			prevIndex := len(spoken[lastTurnNumber]) - 2
			n = (turn - 1) - spoken[lastTurnNumber][prevIndex]
		}
		// fmt.Println(turn, lastTurnNumber, n)

		if spoken[n] == nil {
			spoken[n] = make([]int, 0)
		}
		spoken[n] = append(spoken[n], turn)

		turn++
		lastTurnNumber = n
	}

	return lastTurnNumber
}

func part2(numbers []int) int {
	return 0
}

func main() {
	args := os.Args[1:]
	numbers, err := parseInput(args[0])
	if err != nil {
		panic(err)
	}

	output := part1(numbers)
	fmt.Println(output)

	// output = part2(numbers)
	// fmt.Println(output)
}
