package main

import (
	"fmt"
	"testing"
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
	wa.Step()
	fmt.Println(wa)
}
