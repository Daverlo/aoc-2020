package main

import (
	"fmt"
	"io"
	"os"
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

type Indices struct {
	I int
	J int
}

func part1(numbers []int) int {
	preamble := numbers[0]
	numbers = numbers[1:]

	possibleValues := make(map[int]Indices, preamble*preamble/2)

	for i := 0; i < preamble; i++ {
		for j := i + 1; j < preamble; j++ {
			value := numbers[i] + numbers[j]
			// Only update the indices if this is a better solution
			// It is better if it has a higher low index
			if j > possibleValues[value].I {
				possibleValues[value] = Indices{I: i, J: j}
			}
		}
	}

	for i := preamble; i < len(numbers); i++ {
		indices, ok := possibleValues[numbers[i]]
		if !ok {
			return numbers[i]
		}
		if indices.I < i-preamble {
			return numbers[i]
		}

		for j := i - preamble + 1; j < i; j++ {
			value := numbers[j] + numbers[i]
			// Only update the indices if this is a better solution
			// It is better if it has a higher low index
			if j > possibleValues[value].I {
				// Swap i and j to keep the Indices.I as the lowest index
				possibleValues[value] = Indices{I: j, J: i}
			}
		}
	}
	return 0
}

func part2(numbers []int) int {
	v := part1(numbers)
	numbers = numbers[1:]

	i, j := getSumIndices(numbers, v)
	min, max := minAndMax(numbers[i:j])
	return min + max
}

func getSumIndices(numbers []int, v int) (int, int) {
	sums := make([]int, len(numbers))
	sums[0] = numbers[0]
	for i := 1; i < len(numbers); i++ {
		sums[i] = sums[i-1] + numbers[i]
	}

	i := 0
	j := 1
	for ; j < len(numbers); j++ {
		currentValue := sums[j] - sums[i] + numbers[i]
		if currentValue == v {
			break
		}

		if currentValue < v {
			continue
		}

		for i < j-1 {
			i++
			currentValue := sums[j] - sums[i] + numbers[i]
			if currentValue == v {
				return i, j
			}

			if currentValue < v {
				break
			}
		}
	}

	return -1, -1
}
func minAndMax(numbers []int) (int, int) {
	min := numbers[0]
	max := numbers[0]
	for i := range numbers {
		if numbers[i] < min {
			min = numbers[i]
		}
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return min, max
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
