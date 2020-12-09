package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2(t *testing.T) {
	numbers := []int{100, 1, 2, 3, 5, 7, 5, 8, 101}
	i, j := getSumIndices(numbers, 101)
	assert.Equal(t, 0, i)
	assert.Equal(t, 1, j)

	numbers = []int{100, 1, 2, 3, 5, 7, 5, 8, 103}
	i, j = getSumIndices(numbers, 103)
	assert.Equal(t, 0, i)
	assert.Equal(t, 2, j)

	numbers = []int{100, 1, 2, 3, 5, 7, 5, 8, 25}
	i, j = getSumIndices(numbers, 25)
	assert.Equal(t, 4, i)
	assert.Equal(t, 7, j)
}
