package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccepts(t *testing.T) {
	g := NewGrammar()
	g.Terminals["a"] = true
	g.Terminals["b"] = true

	g.Symbols["0"] = true
	g.Symbols["1"] = true
	g.Symbols["2"] = true
	g.Symbols["3"] = true
	g.Symbols["4"] = true
	g.Symbols["5"] = true
	g.Symbols["6"] = true

	g.Start = "0"

	g.Rules["0"] = [][]string{{"4", "6"}}
	g.Rules["1"] = [][]string{{"2", "3"}, {"3", "2"}}
	g.Rules["2"] = [][]string{{"4", "4"}, {"5", "5"}}
	g.Rules["3"] = [][]string{{"4", "5"}, {"5", "4"}}
	g.Rules["4"] = [][]string{{"a"}}
	g.Rules["5"] = [][]string{{"b"}}
	g.Rules["6"] = [][]string{{"1", "5"}}

	b := g.Accepts("ababbb")
	assert.True(t, b)

	b = g.Accepts("bababa")
	assert.False(t, b)

	b = g.Accepts("abbbab")
	assert.True(t, b)

	b = g.Accepts("aaabbb")
	assert.False(t, b)

	b = g.Accepts("aaaabbb")
	assert.False(t, b)

	g.Rules["0"] = [][]string{{"6", "4"}}
	b = g.Accepts("aaabba")
	assert.True(t, b)
}
