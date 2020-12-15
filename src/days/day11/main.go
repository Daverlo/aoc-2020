package main

import (
	"fmt"
	"io"
	"os"
)

type Position struct {
	I int
	J int
}

type WaitingArea struct {
	Area [][]rune
	Rows int
	Cols int

	visibleSeats map[Position][]Position
}

func (wa *WaitingArea) ComputeVisibleSeats1() {
	wa.visibleSeats = make(map[Position][]Position)

	for i := 0; i < wa.Rows; i++ {
		for j := 0; j < wa.Cols; j++ {
			currentPositon := Position{I: i, J: j}

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

					if wa.Area[ii][jj] == 'L' {
						wa.visibleSeats[currentPositon] = append(wa.visibleSeats[currentPositon], Position{I: ii, J: jj})
					}
				}
			}
		}
	}
}

func (wa *WaitingArea) ComputeVisibleSeats2() {
	wa.visibleSeats = make(map[Position][]Position)

	for i := 0; i < wa.Rows; i++ {
		for j := 0; j < wa.Cols; j++ {
			currentPositon := Position{I: i, J: j}
			for iStep := -1; iStep <= 1; iStep++ {
				for jStep := -1; jStep <= 1; jStep++ {
					if iStep == 0 && jStep == 0 {
						continue
					}
					ii := i
					jj := j
					for {
						ii += iStep
						jj += jStep

						if ii < 0 || ii >= wa.Rows {
							break
						}
						if jj < 0 || jj >= wa.Cols {
							break
						}

						if wa.Area[ii][jj] == 'L' {
							wa.visibleSeats[currentPositon] = append(wa.visibleSeats[currentPositon], Position{I: ii, J: jj})
							break
						}
					}
				}
			}
		}
	}
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
	pos := Position{I: i, J: j}
	visible, _ := wa.visibleSeats[pos]

	for _, p := range visible {
		if wa.Area[p.I][p.J] == '#' {
			occupied++
		}
	}

	return occupied
}

func (wa *WaitingArea) Step(tolerancy int) bool {
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
				if wa.OccupiedAround(i, j) >= tolerancy {
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
	wa.ComputeVisibleSeats1()
	for {
		// fmt.Println(wa)
		changed := wa.Step(4)
		if !changed {
			break
		}
	}
	return wa.OccupiedSeats()
}

func part2(wa *WaitingArea) int {
	wa.ComputeVisibleSeats2()
	for {
		// fmt.Println(wa)
		changed := wa.Step(5)
		if !changed {
			break
		}
	}
	return wa.OccupiedSeats()
}

func main() {
	args := os.Args[1:]
	input1, err := parseInput(args[0])
	if err != nil {
		panic(err)
	}

	output := part1(input1)
	fmt.Println(output)

	input2, err := parseInput(args[0])
	if err != nil {
		panic(err)
	}
	output = part2(input2)
	fmt.Println(output)
}
