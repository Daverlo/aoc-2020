package main

import (
	"fmt"
	"io"
	"os"
)

type passwordPolicy struct {
	Character rune
	Min       int
	Max       int
}

func (pp *passwordPolicy) isValid1(password string) bool {
	n := 0
	for _, c := range password {
		if c == pp.Character {
			n++
		}
	}
	return pp.Min <= n && n <= pp.Max
}

func (pp *passwordPolicy) isValid2(password string) bool {
	first := rune(password[pp.Min-1]) == pp.Character
	second := rune(password[pp.Max-1]) == pp.Character
	return first != second
}

func parseInput(path string) ([]passwordPolicy, []string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	policies := make([]passwordPolicy, 0)
	passwords := make([]string, 0)

	for {
		var min int
		var max int
		var character rune
		var password string
		_, err := fmt.Fscanf(f, "%d-%d %c: %s\n", &min, &max, &character, &password)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, nil, err
			}
		}
		policies = append(policies, passwordPolicy{Character: character, Min: min, Max: max})
		passwords = append(passwords, password)
	}

	return policies, passwords, nil
}

func part1(policies []passwordPolicy, passwords []string) int {
	res := 0
	for i, pp := range policies {
		if pp.isValid1(passwords[i]) {
			res++
		}
	}
	return res
}

func part2(policies []passwordPolicy, passwords []string) int {
	res := 0
	for i, pp := range policies {
		if pp.isValid2(passwords[i]) {
			res++
		}
	}
	return res
}

func main() {
	args := os.Args[1:]
	policies, passwords, err := parseInput(args[0])

	if err != nil {
		panic(err)
	}

	output := part1(policies, passwords)
	fmt.Println(output)

	output = part2(policies, passwords)
	fmt.Println(output)
}
