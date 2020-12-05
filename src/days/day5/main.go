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

func part2(input []string) int {
	// Compute all the seat ids
	var seatIds []int
	for _, line := range input {
		row, col := DecodeSeat(line)
		id := row*8 + col
		seatIds = append(seatIds, id)
	}

	sort.Ints(seatIds)
	for i := 1; i < len(seatIds); i++ {
		// All the seatID should be contiguous
		// If there is a gap that is the seat
		if seatIds[i] == seatIds[i-1]+2 {
			return seatIds[i] - 1
		}
	}

	return -1
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

	output = part2(input)
	fmt.Println(output)
}
