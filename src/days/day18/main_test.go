package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	e := "2 * 3 + (4 * 5)"
	v, err := eval(e)
	assert.NoError(t, err)
	assert.Equal(t, 26, v)
}
