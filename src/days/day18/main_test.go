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

func TestParseExpr(t *testing.T) {
	e := "2 * 3 + ( 4 * 5 )"

	p := make(map[string]int)
	p["*"] = 1
	p["+"] = 1

	a := make(map[string]int)
	a["*"] = -1
	a["+"] = -1

	expected := []string{"2", "3", "*", "4", "5", "*", "+"}
	v, err := parseExpr(e, p, a)
	assert.NoError(t, err)
	assert.Equal(t, expected, v)
}

func TestParseExpr2(t *testing.T) {
	e := "2 * 3 + ( 4 * 5 )"

	p := make(map[string]int)
	p["*"] = 1
	p["+"] = 2

	a := make(map[string]int)
	a["*"] = -1
	a["+"] = -1

	expected := []string{"2", "3", "4", "5", "*", "+", "*"}
	v, err := parseExpr(e, p, a)
	assert.NoError(t, err)
	assert.Equal(t, expected, v)
}

func TestEvalRPN(t *testing.T) {
	rpn := []string{"2", "3", "*", "4", "5", "*", "+"}
	v := evalRPN(rpn)
	assert.Equal(t, 26, v)

	rpn = []string{"2", "3", "4", "5", "*", "+", "*"}
	v = evalRPN(rpn)
	assert.Equal(t, 46, v)
}
