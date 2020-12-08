package main

import (
	"fmt"
	"io"
	"os"

	"github.com/daverlo/aoc-2020/src/graphs"
)

type Operation string

const (
	ACC Operation = "acc"
	JMP Operation = "jmp"
	NOP Operation = "nop"
)

func (o *Operation) Repair() {
	if *o == JMP {
		*o = NOP
	} else {
		*o = JMP
	}
}

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

func (p *Program) Reset() {
	p.runned = make([]bool, len(p.Code))
	p.ip = 0
	p.Accumulator = 0
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

func (p *Program) Run() {
	for p.ip < len(p.Code) {
		p.Step()
	}
}

func (p *Program) RunAndHaltOnLoop() {
	for p.ip < len(p.Code) && !p.runned[p.ip] {
		p.Step()
	}
}

func (p *Program) buildGraph() *graphs.Graph {
	g := graphs.NewGraph(len(p.Code) + 1)
	for i, inst := range p.Code {
		switch inst.Operation {
		case ACC:
			g.AddEdge(i, i+1, 100)
		case JMP:
			g.AddEdge(i, i+inst.Argument, 100)
			if inst.Argument != 1 {
				g.AddEdge(i, i+1, 1)
			}
		case NOP:
			g.AddEdge(i, i+1, 100)
			if inst.Argument != 1 {
				g.AddEdge(i, i+inst.Argument, 1)
			}
		}
	}

	return g
}

func (p *Program) Repair() {
	g := p.buildGraph()

	_, cuts := graphs.MinCut(g, 0, g.Nodes-2)
	for _, edge := range cuts {
		corruptInst := edge.Src
		p.Code[corruptInst].Operation.Repair()

		p.RunAndHaltOnLoop()
		if p.ip >= len(p.Code) {
			// The corrupt instruction was found
			p.Reset()
			return
		}

		// Not the correct change. Revert
		p.Code[corruptInst].Operation.Repair()
		p.Reset()
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

func part2(instructions []Instruction) int {
	program := NewProgram(instructions)
	program.Repair()
	program.Run()
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

	output = part2(code)
	fmt.Println(output)
}
