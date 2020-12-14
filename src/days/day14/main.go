package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Address uint64
	Value   uint64
	Mask    string
}

func parseInput(path string) ([]Instruction, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	rd := bufio.NewReader(f)

	var mask string
	res := make([]Instruction, 0)

	for {
		var line string
		line, err = rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		line = strings.TrimSuffix(line, "\n")

		if strings.HasPrefix(line, "mask") {
			mask = line[7:]
		} else {
			var m uint64
			var n uint64
			_, err = fmt.Sscanf(line, "mem[%d] = %d", &m, &n)
			if err != nil {
				return nil, err
			}
			res = append(res, Instruction{Address: m, Value: n, Mask: mask})
		}
	}

	return res, nil
}

func part1(instructions []Instruction) uint64 {
	memory := make(map[uint64]uint64)
	for _, inst := range instructions {
		memory[inst.Address] = applyMask(inst.Mask, inst.Value)
	}

	var res uint64
	for _, v := range memory {
		res += v
	}

	return res
}

func applyMask(mask string, v uint64) uint64 {
	bin := fmt.Sprintf("%036v", strconv.FormatUint(v, 2))
	for i := range mask {
		if mask[i] == '1' {
			bin = bin[:i] + "1" + bin[i+1:]

		}
		if mask[i] == '0' {
			bin = bin[:i] + "0" + bin[i+1:]
		}
	}

	v, _ = strconv.ParseUint(bin, 2, 64)
	return v
}

func part2(numbers []int) int {
	return 0
}

func main() {
	args := os.Args[1:]
	instructions, err := parseInput(args[0])
	if err != nil {
		// panic(err)
	}

	output := part1(instructions)
	fmt.Println(output)

	// output = part2(input)
	// fmt.Println(output)
}
