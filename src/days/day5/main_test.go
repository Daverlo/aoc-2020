package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeSeat(t *testing.T) {
	line := "FBFBBFFRLR"

	row, col := DecodeSeat(line)
	assert.Equal(t, 44, row)
	assert.Equal(t, 5, col)
}
