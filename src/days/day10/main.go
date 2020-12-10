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

var cache map[int]int

func part2(numbers []int) int {
	deviceJolts := max(numbers) + 3

	set := make(map[int]bool, len(numbers)+2)
	for _, n := range numbers {
		set[n] = true
	}
	set[deviceJolts] = true

	cache = make(map[int]int, len(numbers)+2)
	return calcPart2(set, deviceJolts)
}

func max(numbers []int) int {
	max := numbers[0]
	for i := range numbers {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return max
}

func calcPart2(numbers map[int]bool, n int) int {
	if n == 0 {
		return 1
	}

	if !numbers[n] {
		return 0
	}

	value, ok := cache[n]
	if ok {
		return value
	}

	value = calcPart2(numbers, n-1) + calcPart2(numbers, n-2) + calcPart2(numbers, n-3)
	cache[n] = value
	return value
}

func main() {
	args := os.Args[1:]
	input, err := parseInput(args[0])
	if err != nil {
		panic(err)
	}

	inputP1 := make([]int, len(input))
	copy(inputP1, input)
	output := part1(inputP1)
	fmt.Println(output)

	output = part2(input)
	fmt.Println(output)
}
