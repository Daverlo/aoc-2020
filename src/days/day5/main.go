package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func parseInput(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	res := make([]string, 0)

	for {
		var line string
		_, err := fmt.Fscanf(f, "%s\n", &line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		res = append(res, line)
	}

	return res, nil
}

func part1(input []string) int {
	max := 0
	for _, line := range input {
		row, col := DecodeSeat(line)
		id := row*8 + col
		if id > max {
			max = id
		}
	}
	return max
}

func part2(numbers []int) int {
	sort.Ints(numbers)
	for i, n := range numbers {
		t := 2020 - n

		present, j, k := TargetSum(t, numbers[i+1:])
		// j and k are indices of the subarray used in the TargetSum
		j += i + 1
		k += i + 1

		if present {
			return numbers[i] * numbers[j] * numbers[k]
		}
	}
	present, i, j := TargetSum(2020, numbers)
	if present {
		return numbers[i] * numbers[j]
	}
	return 0
}

func DecodeSeat(line string) (int, int) {
	firstRow := 0
	lastRow := 127
	firstCol := 0
	lastCol := 7

	for _, c := range line {
		switch c {
		case 'F':
			lastRow = (firstRow + lastRow) / 2
		case 'B':
			newFirstRow := (firstRow + lastRow) / 2
			if (firstRow+lastRow)%2 == 1 {
				newFirstRow++
			}
			firstRow = newFirstRow
		case 'L':
			lastCol = (firstCol + lastCol) / 2
		case 'R':
			newFirstCol := (firstCol + lastCol) / 2
			if (firstCol+lastCol)%2 == 1 {
				newFirstCol++
			}
			firstCol = newFirstCol
		}
	}

	return firstRow, firstCol
}

func main() {
	args := os.Args[1:]
	input, err := parseInput(args[0])
	if err != nil {
		panic(err)
	}

	output := part1(input)
	fmt.Println(output)

	// output = part2(input)
	// fmt.Println(output)
}
