package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}

	present, i := BinarySearch(1, a)
	assert.True(t, present)
	assert.Equal(t, 0, i)

	present, i = BinarySearch(2, a)
	assert.True(t, present)
	assert.Equal(t, 1, i)

	present, i = BinarySearch(3, a)
	assert.True(t, present)
	assert.Equal(t, 2, i)

	present, i = BinarySearch(5, a)
	assert.True(t, present)
	assert.Equal(t, 4, i)

	present, i = BinarySearch(6, a)
	assert.False(t, present)
	assert.Equal(t, -1, i)
}

func TestTargetSum(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}

	n := 3
	present, i, j := TargetSum(n, a)
	assert.True(t, present)
	assert.Equal(t, 0, i)
	assert.Equal(t, 1, j)

	n = 7
	present, i, j = TargetSum(n, a)
	assert.True(t, present)
	assert.Equal(t, 1, i)
	assert.Equal(t, 4, j)

	n = 6
	present, i, j = TargetSum(n, a)
	assert.True(t, present)
	assert.Equal(t, 0, i)
	assert.Equal(t, 4, j)

	n = 10
	present, i, j = TargetSum(n, a)
	assert.False(t, present)
	assert.Equal(t, -1, i)
	assert.Equal(t, -1, j)
}
