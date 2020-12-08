package main

import (
	"fmt"
	"io"
	"os"
)

type Operation string

const (
	ACC Operation = "acc"
	JMP Operation = "jmp"
	NOP Operation = "nop"
)

type Instruction struct {
	Operation Operation
	Argument  int
}

type Program struct {
	Code   []Instruction
	ip     int
	runned []bool

	Accumulator int
}

func NewProgram(code []Instruction) *Program {
	p := new(Program)

	p.Code = code
	p.runned = make([]bool, len(code))

	return p
}

// Run the instruction pointed by ip and update the pointer
func (p *Program) Step() {
	inst := p.Code[p.ip]
	newIP := p.ip + 1
	switch inst.Operation {
	case ACC:
		p.Accumulator += inst.Argument
	case JMP:
		newIP += inst.Argument - 1
	case NOP:
		// No op
	}

	p.runned[p.ip] = true
	p.ip = newIP
}

func (p *Program) RunAndHaltOnLoop() {
	for !p.runned[p.ip] {
		p.Step()
	}
}

func parseInput(path string) ([]Instruction, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var instructions []Instruction
	for {
		var op string
		var arg int
		_, err := fmt.Fscanf(f, "%s %d\n", &op, &arg)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		instructions = append(instructions, Instruction{Operation: Operation(op), Argument: arg})
	}

	return instructions, nil
}

func part1(instructions []Instruction) int {
	program := NewProgram(instructions)
	program.RunAndHaltOnLoop()
	return program.Accumulator
}

func main() {
	args := os.Args[1:]
	code, err := parseInput(args[0])
	if err != nil {
		panic(err)
	}

	output := part1(code)
	fmt.Println(output)
}
