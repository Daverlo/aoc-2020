package main

import (
	"fmt"
	"io"
	"os"
)

type WaitingArea struct {
	Area [][]rune
	Rows int
	Cols int
}

func (wa *WaitingArea) OccupiedSeats() int {
	var occupied int
	for i := 0; i < wa.Rows; i++ {
		for j := 0; j < wa.Cols; j++ {
			if wa.Area[i][j] == '#' {
				occupied++
			}
		}
	}
	return occupied
}

func (wa *WaitingArea) OccupiedAround(i int, j int) int {
	var occupied int
	for ii := i - 1; ii <= i+1; ii++ {
		for jj := j - 1; jj <= j+1; jj++ {
			if ii == i && jj == j {
				continue
			}
			if ii < 0 || ii >= wa.Rows {
				continue
			}
			if jj < 0 || jj >= wa.Cols {
				continue
			}

			if wa.Area[ii][jj] == '#' {
				occupied++
			}
		}
	}
	return occupied
}

func (wa *WaitingArea) Step() bool {
	newArea := make([][]rune, wa.Rows)
	for i := 0; i < wa.Rows; i++ {
		newArea[i] = make([]rune, wa.Cols)
	}

	changed := false
	for i := 0; i < wa.Rows; i++ {
		for j := 0; j < wa.Cols; j++ {
			// Update the cell
			if wa.Area[i][j] == 'L' {
				if wa.OccupiedAround(i, j) == 0 {
					newArea[i][j] = '#'
					changed = true
				}
			} else if wa.Area[i][j] == '#' {
				if wa.OccupiedAround(i, j) >= 4 {
					newArea[i][j] = 'L'
					changed = true
				} else {
					newArea[i][j] = wa.Area[i][j]
				}
			} else {
				newArea[i][j] = wa.Area[i][j]
			}
		}
	}

	wa.Area = newArea
	return changed
}

func parseInput(path string) (*WaitingArea, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	wa := new(WaitingArea)
	for {
		var line string
		_, err = fmt.Fscanf(f, "%s\n", &line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		row := make([]rune, len(line))
		for i, c := range line {
			row[i] = c
		}
		wa.Cols = len(row)

		wa.Area = append(wa.Area, row)
		wa.Rows++
	}

	return wa, nil
}

func part1(wa *WaitingArea) int {
	for {
		// fmt.Println(wa)
		changed := wa.Step()
		if !changed {
			break
		}
	}
	return wa.OccupiedSeats()
}

func part2(numbers *WaitingArea) int {
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
