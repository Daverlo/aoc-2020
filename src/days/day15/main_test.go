package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	numbers := []int{0, 3, 6}
	n := part1(numbers)
	assert.Equal(t, 436, n)

	numbers = []int{1, 3, 2}
	n = part1(numbers)
	assert.Equal(t, 1, n)

	numbers = []int{2, 1, 3}
	n = part1(numbers)
	assert.Equal(t, 10, n)

	numbers = []int{1, 2, 3}
	n = part1(numbers)
	assert.Equal(t, 27, n)

	numbers = []int{2, 3, 1}
	n = part1(numbers)
	assert.Equal(t, 78, n)

	numbers = []int{3, 2, 1}
	n = part1(numbers)
	assert.Equal(t, 438, n)

	numbers = []int{3, 1, 2}
	n = part1(numbers)
	assert.Equal(t, 1836, n)
}

func TestPart2_1(t *testing.T) {
	numbers := []int{0, 3, 6}
	n := part2(numbers)
	assert.Equal(t, 175594, n)

	numbers = []int{1, 3, 2}
	n = part2(numbers)
	assert.Equal(t, 2578, n)

	numbers = []int{2, 1, 3}
	n = part2(numbers)
	assert.Equal(t, 3544142, n)
}

func TestPart2_2(t *testing.T) {
	numbers := []int{1, 2, 3}
	n := part2(numbers)
	assert.Equal(t, 261214, n)

	numbers = []int{2, 3, 1}
	n = part2(numbers)
	assert.Equal(t, 6895259, n)

	numbers = []int{3, 2, 1}
	n = part2(numbers)
	assert.Equal(t, 18, n)
}
