package main

import (
	"fmt"
	"io"
	"os"
)

type Level struct {
	Level     [][]bool
	sizeX     int
	sizeY     int
	positionX int
	positionY int
}

func (l *Level) move(displacementX int, displacementY int) (int, int) {
	l.positionX = (l.positionX + displacementX) % l.sizeX
	l.positionY = l.positionY + displacementY

	return l.positionX, l.positionY
}

func (l *Level) isPositionEmpty() bool {
	return !l.Level[l.positionY][l.positionX]
}

func parseInput(path string) (*Level, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var level [][]bool

	for {
		var rawRow string
		_, err := fmt.Fscanf(f, "%s\n", &rawRow)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		row := make([]bool, 0, len(rawRow))
		for _, c := range rawRow {
			row = append(row, c == '#')
		}
		level = append(level, row)
	}

	return &Level{Level: level, sizeX: len(level[0]), sizeY: len(level)}, nil
}

func countTreesInPath(level *Level, displacementX int, displacementY int) int {
	trees := 0
	for level.positionY < level.sizeY {
		if !level.isPositionEmpty() {
			trees++
		}

		level.move(displacementX, displacementY)
	}
	return trees
}

func part1(level *Level) int {
	return countTreesInPath(level, 3, 1)
}

type Vec2 struct {
	X int
	Y int
}

func part2(level *Level) int {
	slopes := []Vec2{{X: 1, Y: 1}, {X: 3, Y: 1}, {X: 5, Y: 1}, {X: 7, Y: 1}, {X: 1, Y: 2}}

	res := 1
	for _, vec := range slopes {
		// Reset the position
		level.positionX = 0
		level.positionY = 0

		res *= countTreesInPath(level, vec.X, vec.Y)
	}
	return res
}

func main() {
	args := os.Args[1:]
	level, err := parseInput(args[0])

	if err != nil {
		panic(err)
	}

	output := part1(level)
	fmt.Println(output)

	output = part2(level)
	fmt.Println(output)
}
