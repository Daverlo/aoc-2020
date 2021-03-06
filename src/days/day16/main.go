package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Name string
	L1   int
	H1   int
	L2   int
	H2   int
}

func (r *Rule) IsValid(n int) bool {
	return (n >= r.L1 && n <= r.H1) || (n >= r.L2 && n <= r.H2)
}

func parseInput(path string) ([]Rule, [][]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	rules := make([]Rule, 0)
	tickets := make([][]int, 0)

	readingRules := true
	for {
		if readingRules {
			var name string
			var l1, h1, l2, h2 int
			_, err := fmt.Fscanf(f, "%s %d-%d or %d-%d\n", &name, &l1, &h1, &l2, &h2)
			if err != nil {
				if err.Error() == "unexpected newline" {
					readingRules = false
					var meh, meh2 string
					_, _ = fmt.Fscanf(f, "%s %s\n", &meh, &meh2)
					continue
				} else {
					return nil, nil, err
				}
			}

			rules = append(rules, Rule{Name: name[0 : len(name)-1], L1: l1, H1: h1, L2: l2, H2: h2})
		} else {
			var line string
			_, err = fmt.Fscanf(f, "%s\n", &line)
			if err != nil {
				if err == io.EOF {
					break
				} else if err.Error() == "unexpected newline" {
					var meh, meh2 string
					_, _ = fmt.Fscanf(f, "%s %s\n", &meh, &meh2)
					continue
				} else {
					return nil, nil, err
				}
			}

			rawNumbers := strings.Split(line, ",")
			numbers := make([]int, len(rawNumbers))
			for i, n := range rawNumbers {
				numbers[i], err = strconv.Atoi(n)
				if err != nil {
					return nil, nil, err
				}
			}

			tickets = append(tickets, numbers)
		}
	}

	return rules, tickets, nil
}

func part1(rules []Rule, tickets [][]int) int {
	errorRate := 0

	for i := 1; i < len(tickets); i++ {
		t := tickets[i]
		for _, n := range t {
			isValid := false
			for _, r := range rules {
				if r.IsValid(n) {
					isValid = true
					break
				}
			}
			if !isValid {
				errorRate += n
			}
		}
	}

	return errorRate
}

func part2(rules []Rule, tickets [][]int) int {
	validTickets := make([][]int, 0)
	validTickets = append(validTickets, tickets[0])

	for i := 1; i < len(tickets); i++ {
		t := tickets[i]
		isTicketValid := true
		for _, n := range t {
			isValid := false
			for _, r := range rules {
				if r.IsValid(n) {
					isValid = true
					break
				}
			}
			if !isValid {
				isTicketValid = false
			}
		}
		if isTicketValid {
			validTickets = append(validTickets, t)
		}
	}

	sortedRules := sortFields(rules, validTickets)

	res := 1
	for i, r := range sortedRules {
		if strings.HasPrefix(r.Name, "departure") {
			res *= validTickets[0][i]
		}
	}
	return res
}

func sortFields(rules []Rule, tickets [][]int) []Rule {
	ruleToPossibleIndices := make(map[Rule][]int)
	for _, r := range rules {
		for i := 0; i < len(rules); i++ {
			isValid := true
			for _, t := range tickets {
				if !r.IsValid(t[i]) {
					isValid = false
					break
				}
			}
			if isValid {
				ruleToPossibleIndices[r] = append(ruleToPossibleIndices[r], i)
			}
		}
	}

	sortedRules := make([]Rule, len(rules))
	for len(ruleToPossibleIndices) > 0 {
		for r, indices := range ruleToPossibleIndices {
			if len(indices) == 1 {
				sortedRules[indices[0]] = r
				for r2, indices2 := range ruleToPossibleIndices {
					ruleToPossibleIndices[r2] = deleteElement(indices2, indices[0])
				}
				delete(ruleToPossibleIndices, r)
				break
			}
		}
	}

	return sortedRules
}

func deleteElement(a []int, e int) []int {
	for i, n := range a {
		if n == e {
			a[len(a)-1], a[i] = a[i], a[len(a)-1]
			return a[:len(a)-1]
		}
	}

	return a
}

func main() {
	args := os.Args[1:]
	rules, tickets, err := parseInput(args[0])
	if err != nil {
		panic(err)
	}

	output := part1(rules, tickets)
	fmt.Println(output)

	output = part2(rules, tickets)
	fmt.Println(output)
}
