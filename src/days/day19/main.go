package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type Grammar struct {
	Terminals map[string]bool
	Symbols   map[string]bool
	Rules     map[string][][]string
	Start     string
}

func NewGrammar() *Grammar {
	g := new(Grammar)

	g.Terminals = make(map[string]bool)
	g.Symbols = make(map[string]bool)
	g.Rules = make(map[string][][]string)
	g.Start = "0"

	return g
}

func (g *Grammar) matches(symbol, s string, mem map[string]map[string]string) bool {
	if mem[symbol][s] == "True" {
		return true
	}
	if mem[symbol][s] == "False" {
		return false
	}

	if len(s) == 1 {
		for _, rule := range g.Rules[symbol] {
			// Substitution is a terminal that matches s
			if len(rule) == 1 && g.Terminals[rule[0]] && s == rule[0] {
				return true
			}
		}
	}

	mem[symbol][s] = "False"

	for _, rule := range g.Rules[symbol] {
		if len(rule) == 1 {
			if g.matches(rule[0], s, mem) {
				mem[symbol][s] = "True"
				return true
			}
		} else if len(rule) == 2 {
			for p := 1; p <= len(s)-1; p++ {
				a, b := s[:p], s[p:]

				if g.matches(rule[0], a, mem) && g.matches(rule[1], b, mem) {
					mem[symbol][s] = "True"
					return true
				}
			}
		}
	}

	return false
}

func (g *Grammar) Accepts(s string) bool {
	mem := make(map[string]map[string]string)
	for s := range g.Symbols {
		mem[s] = make(map[string]string)
	}
	for t := range g.Terminals {
		mem[t] = make(map[string]string)
	}
	return g.matches(g.Start, s, mem)
}

func parseInput(path string) (*Grammar, []string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	rd := bufio.NewReader(f)
	ruleRegexp := regexp.MustCompile("^([0-9]+): (.*)")
	terminalRegexp := regexp.MustCompile("\"(.+)\"")

	g := NewGrammar()
	var messages []string
	for {
		var line string
		line, err = rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, nil, err
			}
		}
		line = strings.TrimSuffix(line, "\n")
		if line == "" {
			continue
		}

		match := ruleRegexp.FindStringSubmatch(line)
		if match == nil {
			// line is a message
			messages = append(messages, line)
			continue
		}

		symbol := match[1]
		rewrites := match[2]

		g.Symbols[symbol] = true
		for _, rewrite := range strings.Split(rewrites, " | ") {
			rewriteTokens := make([]string, 0)
			for _, rewriteSymbol := range strings.Split(rewrite, " ") {
				// Check if the symbol is a terminal
				// A terminal is surrounded by double-quotes ("a")
				m := terminalRegexp.FindStringSubmatch(rewriteSymbol)
				if m != nil {
					terminal := m[1]
					g.Terminals[terminal] = true
					rewriteTokens = append(rewriteTokens, terminal)
					continue
				}
				rewriteTokens = append(rewriteTokens, rewriteSymbol)
			}
			g.Rules[symbol] = append(g.Rules[symbol], rewriteTokens)
		}
	}

	return g, messages, nil
}

func part1(g *Grammar, messages []string) int {
	res := 0
	for _, m := range messages {
		if g.Accepts(m) {
			res++
		}
	}
	return res
}

func part2(g *Grammar, messages []string) int {
	res := 0

	g.Rules["8"] = [][]string{{"42"}, {"42", "8"}}

	// 11: 42 31 | 42 11 31
	// In CNF:
	// 11: 42 31 | 999 31
	// 999: 42 11
	g.Symbols["999"] = true
	g.Rules["11"] = [][]string{{"42", "31"}, {"999", "31"}}
	g.Rules["999"] = [][]string{{"42", "11"}}

	for _, m := range messages {
		if g.Accepts(m) {
			res++
		}
	}
	return res
}

func main() {
	args := os.Args[1:]
	g, messages, err := parseInput(args[0])

	if err != nil {
		panic(err)
	}

	output := part1(g, messages)
	fmt.Println(output)

	output = part2(g, messages)
	fmt.Println(output)
}
