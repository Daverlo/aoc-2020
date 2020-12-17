package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var orientationsCW = "NESW"
var orientationsCCW = "NWSE"

type Instruction struct {
	Action rune
	Value  int
}

func (i *Instruction) Execute(x, y int, dir rune) (int, int, rune) {
	switch i.Action {
	case 'N':
		y += i.Value
	case 'S':
		y -= i.Value
	case 'E':
		x += i.Value
	case 'W':
		x -= i.Value
	case 'L':
		index := strings.Index(orientationsCCW, string(dir))
		turns := i.Value / 90
		newIndex := (index + turns) % 4
		dir = rune(orientationsCCW[newIndex])
	case 'R':
		index := strings.Index(orientationsCW, string(dir))
		turns := i.Value / 90
		newIndex := (index + turns) % 4
		dir = rune(orientationsCW[newIndex])
	case 'F':
		newI := Instruction{Action: dir, Value: i.Value}
		x, y, dir = newI.Execute(x, y, dir)
	}
	return x, y, dir
}

func (i *Instruction) ExecuteWaypoint(x, y int, wpX, wpY int) (int, int, int, int) {
	switch i.Action {
	case 'N':
		wpY += i.Value
	case 'S':
		wpY -= i.Value
	case 'E':
		wpX += i.Value
	case 'W':
		wpX -= i.Value
	case 'L':
		turns := i.Value / 90
		for i := 0; i < turns; i++ {
			wpX, wpY = -wpY, wpX
		}
	case 'R':
		turns := i.Value / 90
		for i := 0; i < turns; i++ {
			wpX, wpY = wpY, -wpX
		}
	case 'F':
		x += wpX * i.Value
		y += wpY * i.Value
	}
	return x, y, wpX, wpY
}

func parseInput(path string) ([]Instruction, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	res := make([]Instruction, 0)

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
		a := rune(line[0])
		v, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}

		res = append(res, Instruction{Action: a, Value: v})
	}

	return res, nil
}

func part1(instructions []Instruction) int {
	x := 0
	y := 0
	dir := 'E'

	for _, i := range instructions {
		// fmt.Println(x, y, dir)
		x, y, dir = i.Execute(x, y, dir)
	}
	// fmt.Println(x, y, dir)

	if x < 0 {
		x = -x
	}

	if y < 0 {
		y = -y
	}

	return x + y
}

func part2(instructions []Instruction) int {
	x := 0
	y := 0
	wpX := 10
	wpY := 1

	for _, i := range instructions {
		// fmt.Println(x, y, dir)
		x, y, wpX, wpY = i.ExecuteWaypoint(x, y, wpX, wpY)
	}
	// fmt.Println(x, y, dir)

	if x < 0 {
		x = -x
	}

	if y < 0 {
		y = -y
	}

	return x + y
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
