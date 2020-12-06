package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type GroupAnswers []string

func (ga GroupAnswers) countYes() int {
	count := 0
	yes := make(map[rune]bool)
	for i := 0; i < len(ga); i++ {
		for _, c := range ga[i] {
			if !yes[c] {
				yes[c] = true
				count++
			}
		}
	}

	return count
}

func parseInput(path string) ([]GroupAnswers, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	rd := bufio.NewReader(f)

	var groups []GroupAnswers
	for {
		var answers GroupAnswers
		var err error

		for {
			var line string
			line, err = rd.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return nil, err
				}
			}
			line = strings.TrimSuffix(line, "\n")

			// Group delimiter
			if len(line) == 0 {
				break
			}

			answers = append(answers, line)
		}

		groups = append(groups, answers)

		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
	}

	return groups, nil
}

func part1(groups []GroupAnswers) int {
	res := 0

	for _, answers := range groups {
		res += answers.countYes()
	}

	return res
}

func part2(groups []GroupAnswers) int {
	return 0
}

func main() {
	args := os.Args[1:]
	passports, err := parseInput(args[0])

	if err != nil {
		panic(err)
	}

	output := part1(passports)
	fmt.Println(output)

	// output = part2(passports)
	// fmt.Println(output)
}
