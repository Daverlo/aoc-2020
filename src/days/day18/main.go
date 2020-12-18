package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func parseInput(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	rd := bufio.NewReader(f)

	var expressions []string
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
		expressions = append(expressions, line)
	}

	return expressions, nil
}

func part1(expressions []string) int {
	res := 0

	for _, e := range expressions {
		v, err := eval(e)
		if err != nil {
			panic(err)
		}

		// fmt.Println(v)
		res += v
	}

	return res
}

func eval(e string) (int, error) {
	res := 0

	operand1 := -1
	operand2 := -1
	operation := -1
	i := 0
	for i < len(e) {
		if e[i] == ' ' {
			i++
			continue
		}
		if e[i] == '(' {
			closingsNeeded := 1
			j := i + 1
			for j < len(e) {
				if e[j] == ')' {
					closingsNeeded--
					if closingsNeeded == 0 {
						break
					}
				}
				if e[j] == '(' {
					closingsNeeded++
				}
				j++
			}

			if j == len(e) {
				return 0, errors.New("Malformed expression")
			}

			v, err := eval(e[i+1 : j])
			if err != nil {
				return 0, err
			}

			e = e[:i] + fmt.Sprint(v) + e[j+1:]
			continue
		}
		if e[i] == '+' {
			operation = 1
			i++
			continue
		}
		if e[i] == '*' {
			operation = 2
			i++
			continue
		}

		// Now e[i] must be a number
		j := i + 1
		for j < len(e) {
			if e[j] == ' ' {
				break
			}
			j++
		}

		n, err := strconv.Atoi(e[i:j])
		if err != nil {
			return 0, nil
		}

		if operation == -1 {
			operand1 = n
			i = j
			continue
		} else {
			operand2 = n
		}

		// Full operation with both operands is available now
		if operation == 1 {
			res = operand1 + operand2
		} else if operation == 2 {
			res = operand1 * operand2

		}
		operand1 = res
		operation = -1
		operand2 = -1
		i = j
	}

	return res, nil
}

func part2(expressions []string) int {
	res := 0

	return res
}

func main() {
	args := os.Args[1:]
	expressions, err := parseInput(args[0])

	if err != nil {
		panic(err)
	}

	output := part1(expressions)
	fmt.Println(output)

	output = part2(expressions)
	fmt.Println(output)
}
