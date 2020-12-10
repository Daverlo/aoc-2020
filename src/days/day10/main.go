package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func parseInput(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	res := make([]int, 0)

	for {
		var n int
		_, err := fmt.Fscanf(f, "%d\n", &n)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		res = append(res, n)
	}

	return res, nil
}

func part1(numbers []int) int {
	numbers = append(numbers, 0)
	sort.Ints(numbers)

	gap1 := 0
	gap3 := 1
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i+1]-numbers[i] == 1 {
			gap1++
		}
		if numbers[i+1]-numbers[i] == 3 {
			gap3++
		}
	}

	return gap1 * gap3
}

func part2(numbers []int) int {
	return 0
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
