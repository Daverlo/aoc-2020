package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	Passport       string
	Country        string
}

func (p *Passport) AddFields(fields map[string]string) {
	for k, v := range fields {
		switch k {
		case "byr":
			p.BirthYear = v
		case "iyr":
			p.IssueYear = v
		case "eyr":
			p.ExpirationYear = v
		case "hgt":
			p.Height = v
		case "hcl":
			p.HairColor = v
		case "ecl":
			p.EyeColor = v
		case "pid":
			p.Passport = v
		case "cid":
			p.Country = v
		}
	}
}

func (p *Passport) IsValid() bool {
	if p.BirthYear == "" {
		return false
	}

	if p.IssueYear == "" {
		return false
	}

	if p.ExpirationYear == "" {
		return false
	}

	if p.Height == "" {
		return false
	}

	if p.HairColor == "" {
		return false
	}

	if p.EyeColor == "" {
		return false
	}

	if p.Passport == "" {
		return false
	}

	// Ignore this ¯\_(ツ)_/¯
	// if p.Country == "" {
	// 	return false
	// }

	return true
}

func parseInput(path string) ([]Passport, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	rd := bufio.NewReader(f)

	var passports []Passport
	for {
		var passport Passport
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

			// Passport delimiter
			if len(line) == 0 {
				break
			}

			fields := parseLine(line)
			passport.AddFields(fields)
		}

		passports = append(passports, passport)

		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
	}

	return passports, nil
}

func parseLine(line string) map[string]string {
	fields := make(map[string]string)

	tokens := strings.Split(line, " ")
	for _, token := range tokens {
		field := strings.Split(token, ":")
		fields[field[0]] = field[1]
	}

	return fields
}

func part1(passports []Passport) int {
	validPassports := 0

	for _, p := range passports {
		if p.IsValid() {
			validPassports++
		}
	}

	return validPassports
}

func main() {
	args := os.Args[1:]
	level, err := parseInput(args[0])

	if err != nil {
		panic(err)
	}

	output := part1(level)
	fmt.Println(output)

	// output = part2(level)
	// fmt.Println(output)
}
