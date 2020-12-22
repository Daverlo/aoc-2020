package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayGame(t *testing.T) {
	d1 := []int{9, 2, 6, 3, 1}
	d2 := []int{5, 8, 4, 7, 10}
	w, v := PlayGame(d1, d2)
	assert.Equal(t, 2, w)
	assert.Equal(t, 291, v)

	d1 = []int{43, 19}
	d2 = []int{2, 29, 14}
	w, v = PlayGame(d1, d2)
	assert.Equal(t, w, 1)
	assert.Equal(t, v, 105)
}
