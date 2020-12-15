package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStep(t *testing.T) {
	wa := WaitingArea{
		Rows: 10,
		Cols: 10,
		Area: [][]rune{
			{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
			{'L', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
			{'L', '.', 'L', '.', 'L', '.', '.', 'L', '.', '.'},
			{'L', 'L', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
			{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
			{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
			{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
			{'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L'},
			{'L', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
			{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'}}}
	fmt.Println(wa)
	wa.Step(4)
	fmt.Println(wa)
}

func TestPart2(t *testing.T) {
	wa := WaitingArea{
		Rows: 10,
		Cols: 10,
		Area: [][]rune{
			{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
			{'L', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
			{'L', '.', 'L', '.', 'L', '.', '.', 'L', '.', '.'},
			{'L', 'L', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
			{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
			{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
			{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
			{'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L'},
			{'L', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
			{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'}}}
	res := part2(&wa)
	assert.Equal(t, 26, res)
}
