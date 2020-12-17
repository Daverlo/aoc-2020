package main

import (
	"fmt"
	"io"
	"os"
)

type Coordinates struct {
	X int
	Y int
	Z int
}

type Conway struct {
	Area map[Coordinates]bool

	minX int
	maxX int
	minY int
	maxY int
	minZ int
	maxZ int
}

func NewConway(floor [][]rune) *Conway {
	conway := new(Conway)

	conway.Area = make(map[Coordinates]bool, 0)
	for x := range floor {
		for y := range floor[x] {
			if floor[x][y] == '#' {
				conway.Area[Coordinates{X: x, Y: y, Z: 0}] = true
			}
		}
	}

	conway.minX = 0
	conway.maxX = len(floor) - 1
	conway.minY = 0
	conway.maxY = len(floor[0]) - 1
	conway.minZ = 0
	conway.maxZ = 0

	return conway
}

func (c *Conway) CountActive() int {
	return len(c.Area)
}

func (c *Conway) CountNeighbours(coord Coordinates) int {
	actives := 0

	for x := coord.X - 1; x <= coord.X+1; x++ {
		for y := coord.Y - 1; y <= coord.Y+1; y++ {
			for z := coord.Z - 1; z <= coord.Z+1; z++ {
				if x == coord.X && y == coord.Y && z == coord.Z {
					continue
				}
				if c.Area[Coordinates{X: x, Y: y, Z: z}] {
					actives++
				}
			}
		}
	}

	return actives
}

func (c *Conway) Activate(coord Coordinates) {
	c.Area[coord] = true

	if coord.X < c.minX {
		c.minX = coord.X
	}
	if coord.X > c.maxX {
		c.maxX = coord.X
	}

	if coord.Y < c.minY {
		c.minY = coord.Y
	}
	if coord.Y > c.maxY {
		c.maxY = coord.Y
	}

	if coord.Z < c.minZ {
		c.minZ = coord.Z
	}
	if coord.Z > c.maxZ {
		c.maxZ = coord.Z
	}
}

func (c *Conway) Step() {
	newConway := new(Conway)
	newConway.Area = make(map[Coordinates]bool, len(c.Area))

	// We need to look 1 step outside the frontier too
	for x := c.minX - 1; x <= c.maxX+1; x++ {
		for y := c.minY - 1; y <= c.maxY+1; y++ {
			for z := c.minZ - 1; z <= c.maxZ+1; z++ {
				coord := Coordinates{X: x, Y: y, Z: z}
				neighbours := c.CountNeighbours(coord)
				if c.Area[coord] {
					if neighbours == 2 || neighbours == 3 {
						newConway.Activate(coord)
					}
				} else {
					if neighbours == 3 {
						newConway.Activate(coord)
					}
				}
			}
		}
	}
	*c = *newConway
}

func parseInput(path string) (*Conway, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	floor := make([][]rune, 0)
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

		floor = append(floor, row)
	}

	return NewConway(floor), nil
}

func part1(c *Conway) int {
	for i := 0; i < 6; i++ {
		c.Step()
	}
	return c.CountActive()
}

type Coordinates4D struct {
	X int
	Y int
	Z int
	W int
}

type Conway4D struct {
	Area map[Coordinates4D]bool

	minX int
	maxX int
	minY int
	maxY int
	minZ int
	maxZ int
	minW int
	maxW int
}

func NewConway4D(floor [][]rune) *Conway4D {
	conway := new(Conway4D)

	conway.Area = make(map[Coordinates4D]bool, 0)
	for x := range floor {
		for y := range floor[x] {
			if floor[x][y] == '#' {
				conway.Area[Coordinates4D{X: x, Y: y, Z: 0, W: 0}] = true
			}
		}
	}

	conway.minX = 0
	conway.maxX = len(floor) - 1
	conway.minY = 0
	conway.maxY = len(floor[0]) - 1
	conway.minZ = 0
	conway.maxZ = 0
	conway.minW = 0
	conway.maxW = 0

	return conway
}

func (c *Conway4D) CountActive() int {
	return len(c.Area)
}

func (c *Conway4D) CountNeighbours(coord Coordinates4D) int {
	actives := 0

	for x := coord.X - 1; x <= coord.X+1; x++ {
		for y := coord.Y - 1; y <= coord.Y+1; y++ {
			for z := coord.Z - 1; z <= coord.Z+1; z++ {
				for w := coord.W - 1; w <= coord.W+1; w++ {
					if x == coord.X && y == coord.Y && z == coord.Z && w == coord.W {
						continue
					}
					if c.Area[Coordinates4D{X: x, Y: y, Z: z, W: w}] {
						actives++
					}
				}
			}
		}
	}

	return actives
}

func (c *Conway4D) Activate(coord Coordinates4D) {
	c.Area[coord] = true

	if coord.X < c.minX {
		c.minX = coord.X
	}
	if coord.X > c.maxX {
		c.maxX = coord.X
	}

	if coord.Y < c.minY {
		c.minY = coord.Y
	}
	if coord.Y > c.maxY {
		c.maxY = coord.Y
	}

	if coord.Z < c.minZ {
		c.minZ = coord.Z
	}
	if coord.Z > c.maxZ {
		c.maxZ = coord.Z
	}

	if coord.W < c.minW {
		c.minW = coord.W
	}
	if coord.W > c.maxW {
		c.maxW = coord.W
	}
}

func (c *Conway4D) Step() {
	newConway := new(Conway4D)
	newConway.Area = make(map[Coordinates4D]bool, len(c.Area))

	// We need to look 1 step outside the frontier too
	for x := c.minX - 1; x <= c.maxX+1; x++ {
		for y := c.minY - 1; y <= c.maxY+1; y++ {
			for z := c.minZ - 1; z <= c.maxZ+1; z++ {
				for w := c.minW - 1; w <= c.maxW+1; w++ {
					coord := Coordinates4D{X: x, Y: y, Z: z, W: w}
					neighbours := c.CountNeighbours(coord)
					if c.Area[coord] {
						if neighbours == 2 || neighbours == 3 {
							newConway.Activate(coord)
						}
					} else {
						if neighbours == 3 {
							newConway.Activate(coord)
						}
					}
				}

			}
		}
	}
	*c = *newConway
}

func parseInput4D(path string) (*Conway4D, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	floor := make([][]rune, 0)
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

		floor = append(floor, row)
	}

	return NewConway4D(floor), nil
}

func part2(c *Conway4D) int {
	for i := 0; i < 6; i++ {
		c.Step()
	}
	return c.CountActive()
}

func main() {
	args := os.Args[1:]
	input1, err := parseInput(args[0])
	if err != nil {
		panic(err)
	}

	output := part1(input1)
	fmt.Println(output)

	input2, err := parseInput4D(args[0])
	if err != nil {
		panic(err)
	}
	output = part2(input2)
	fmt.Println(output)
}
